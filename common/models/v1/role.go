package models

type Role struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	UserRole []UserRole `gorm:"foreignKey:RoleID"`
	//Permissions []Permission `gorm:"many2many:role_permissions;"`
}
