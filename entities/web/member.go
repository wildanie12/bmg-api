package web

// UserRequest formatting entity
type UserRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Name string `form:"name" validate:"required"`
	Email string `form:"email" validate:"required"`
}

// UserResponse formatting entity
type UserResponse struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	ReferralCode string `json:"referral_code"`
}