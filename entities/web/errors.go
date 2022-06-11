package web

// Error provide implementation of default error type interface
// To pack more information such as code and more if needed
type Error struct {
	Message string
	Code int
}

// Error returns error message created 
func (err Error) Error() string {
	return err.Message
}



// ValidationErrorItem contains every validation error happen
// which is part of ValidationError
type ValidationErrorItem struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
// ValidationError represent single item of each va
// validation that fails on every form submit endpoint hit
type ValidationError struct {
	Message string
	Code int
	Errors []ValidationErrorItem
}
func (err ValidationError) Error() string {
	return err.Message
}