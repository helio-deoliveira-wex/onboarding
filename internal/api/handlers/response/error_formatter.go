package response

import (
	"github.com/pkg/errors"
	"onboarding/internal/service/app_error"
)

func Format(e app_error.AppError) ErrorResponse {
	c := errors.Cause(e.(error))

	if detailedErr, ok := c.(app_error.DetailedError); ok && detailedErr.HasDetails() {
		return ErrorResponse{
			Message: c.Error(),
			Details: detailedErr.GetDetails()}
	} else {
		return ErrorResponse{Message: c.Error()}
	}
}
