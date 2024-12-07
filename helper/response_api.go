package helper

type meta struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type response struct {
	Meta meta `json:"meta"`
	Data any  `json:"data"`
}

// ResponseAPI generates a response object for the API.
//
// Parameters:
//
//	success: a boolean indicating whether the API call was successful or not.
//	code: an integer representing the status code of the response.
//	message: a string containing any additional message associated with the response.
//	data: any additional data that needs to be included in the response.
//
// Returns:
//
//	response: the generated response object.
func ResponseAPI(success bool, code int, message string, data any) response {
	return response{Meta: meta{
		Success: success,
		Code:    code,
		Message: message,
	}, Data: data}
}
