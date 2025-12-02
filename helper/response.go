package helper

func ResponseSuccess(data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": message,
		"data":    data,
	}
}

func ResponseError(err error, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  false,
		"message": message,
		"error":   err.Error(),
	}
}
