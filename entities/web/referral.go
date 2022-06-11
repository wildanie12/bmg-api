package web

// ReferralRequest formatting and validation rules
type ReferralRequest struct {
	Referral string `form:"referral" validate:"required"`
}

// ReferralResponse formatting
type ReferralResponse struct {
	User UserResponse `json:"user"`
}