package models

type RolePermission struct {
	RoleID       int `gorm:"primaryKey"`
	PermissionID int `gorm:"primaryKey"`

	// Optionally, you can define relationships if needed
	Role       Role       `gorm:"foreignKey:RoleID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}
