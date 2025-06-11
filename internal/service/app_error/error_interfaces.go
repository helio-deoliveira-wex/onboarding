package app_error

type AppError interface {
	ErrorType() AppErrorType
}

type DetailedError interface {
	HasDetails() bool
	GetDetails() []string
}
