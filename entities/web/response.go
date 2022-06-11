package web

// SuccessResponse provide response struct formatting
// representing success of an endpoint hit
type SuccessResponse struct {
	Status string `json:"status"`
	Code int	`json:"code"`
	Data string `json:"data"`
}

// ValidationErrorResponse represent validation error
// on every failed validation endpoint hit
type ValidationErrorResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Error string `json:"error"`
	Errors []ValidationErrorItem `json:"errors"`
}

// ErrorResponse for giving web error response
type ErrorResponse struct {
	Status string `josn:"status"`
	Code int `json:"code"`
	Error string `json:"error"`
}