package domain

// User entity reponsible for database migration
// and domain level manipulation data modeling
type User struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"size:255"`
	Name string `gorm:"size:255"`
	Email string `gorm:"size:255"`
	ReferralCode string `gorm:"size:255"`
}