package web

// UserResponse formatting entity
type UserResponse struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	ReferralCode string `json:"referral_code"`
}