package helper

// HTTPResponse : Wrapper for HTTP Response
type HTTPResponse struct {
	Error       bool        `json:"error"`
	Data        interface{} `json:"data,omitempty"`
	Message     string      `json:"message"`
	Description string      `json:"description,omitempty"`
}

// NewSuccessResponse : Contstructor for generating new success HTTP Response
func NewSuccessResponse(field string, data interface{}) HTTPResponse {
	var dataWrapper = make(map[string]interface{})
	if field != "" {
		dataWrapper[field] = data
	}

	return HTTPResponse{
		Error:   false,
		Data:    dataWrapper,
		Message: "success",
	}
}

// NewErrorResponse : Contstructor for generating new error HTTP Response
func NewErrorResponse(message string, description string) HTTPResponse {
	WriteToLogFile(message, description)
	var dataWrapper = make(map[string]interface{})
	return HTTPResponse{
		Error:       false,
		Data:        dataWrapper,
		Message:     message,
		Description: description,
	}
}
