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



// ValidationError represent single item of each va
// validation that fails on every form submit endpoint hit
type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}