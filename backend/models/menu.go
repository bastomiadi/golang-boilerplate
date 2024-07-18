// backend/models/menu.go
package models

type Menu struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Parent int    `json:"parent"`
}
