package domain

type Member struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"size:255"`
	Name string `gorm:"size:255"`
	Email string `gorm:"size:255"`
	ReferralCode string `gorm:"size:255"`
}