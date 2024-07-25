package models

type Profile struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"uniqueIndex"`
	Address string `gorm:"type:varchar(255)"`
	Phone   string `gorm:"type:varchar(20)"`
	Bio     string `gorm:"type:text"`
}
