package web

// AuthRequest formatting
type AuthRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// AuthResponse formatting, associated with UserResponse
type AuthResponse struct {
	Token string `json:"token"`
	User UserResponse `json:"user"`
}