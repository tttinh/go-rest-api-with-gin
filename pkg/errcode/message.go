package errcode

var errorMessages = map[Code]string{
	ErrInvalidInput: "Your input data is not valid!",
}

// GetMessage get error information based on Code
func GetMessage(code Code) string {
	msg, ok := errorMessages[code]
	if ok {
		return msg
	}

	return errorMessages["Undefined error!"]
}
