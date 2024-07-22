package models

type Product struct {
	ID       int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Price    float64 `gorm:"not null" json:"price"`
	Category string  `gorm:"type:varchar(255);not null" json:"category"`
	// Add more fields as needed
}
