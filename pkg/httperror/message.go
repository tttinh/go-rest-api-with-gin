package httperror

var errorMessages = map[Code]string{
	ErrInvalidInput: "Your input data is not valid!",
}

// getErrorMessage get error information based on Code
func getErrorMessage(code Code) string {
	msg, ok := errorMessages[code]
	if ok {
		return msg
	}

	return errorMessages["Undefined error!"]
}
