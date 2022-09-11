package errors

type AppError struct {
	// err, "Record not found", 404
	Error      error
	Message    string
	StatusCode int16
}

func InternalError(err error) *AppError {
	return &AppError{err, "Internal server error", 500}
}

func BadRequest(msg string, err error) *AppError {
	return &AppError{err, msg, 400}
}

func NotFound(msg string, err error) *AppError {
	message := msg + " not found"
	return &AppError{err, message, 404}
}
