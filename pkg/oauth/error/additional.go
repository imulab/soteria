package error

func ContextCancelled() error {
	return ServerError("request timed out or was cancelled.")
}
