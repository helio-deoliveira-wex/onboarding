package database

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutUser(t *testing.T) {
	t.Run("Positive - Update User", func(t *testing.T) {
		user := &User{
			Id:        "user-1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Age:       30,
		}
		now := time.Now()
		err := PutUser(user)
		require.NoError(t, err)
		require.NotNil(t, user)
		assert.Greater(t, user.UpdatedAt, now)
	})
	t.Run("Positive - Concurrently update a user", func(t *testing.T) {
		// Create Wait Group
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			// Add one wait per iteration
			wg.Add(1)
			go func(wGroup *sync.WaitGroup) {
				user := &User{
					Id:        "user-1",
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john.doe@example.com",
					Age:       30,
				}
				// Defer to avoid infinite executions
				defer wGroup.Done()
				now := time.Now()
				err := PutUser(user)
				require.NoError(t, err)
				require.NotNil(t, user)
				assert.Greater(t, user.UpdatedAt, now)
				fmt.Print("1")
			}(&wg)
			fmt.Print("2")
		}
		wg.Wait()
		fmt.Print("3")
	})
}
