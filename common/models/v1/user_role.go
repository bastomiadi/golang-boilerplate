package models

type UserRole struct {
	UserID int `gorm:"primaryKey"`
	RoleID int `gorm:"primaryKey"`

	// Optionally, you can define relationships if needed
	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}
