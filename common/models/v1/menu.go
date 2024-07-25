package models

import (
	"fmt"
	"golang-boilerplate/config"
	"strings"
)

type Menu struct {
	ID           int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string       `gorm:"type:varchar(255);not null" json:"name"`
	Link         string       `gorm:"type:varchar(255)" json:"link"`
	Parent       int          `gorm:"default:0" json:"parent"`
	Icon         string       `gorm:"type:varchar(255)" json:"icon"`
	DisplayOrder int          `gorm:"default:0" json:"display_order"`
	//Permissions  []Permission `gorm:"many2many:menu_permissions;"` // tambahan testing
}

// GetMenuItems retrieves menu items ordered by parent and display order using GORM
func GetMenuItems() ([]Menu, error) {
	db := config.GetDB()

	var menuItems []Menu
	if err := db.Order("parent, display_order").Find(&menuItems).Error; err != nil {
		return nil, err
	}

	return menuItems, nil
}

// HasChildren checks if there are any menu items with the specified parent ID
func HasChildren(menuItems []Menu, parentID int) bool {
	for _, menuItem := range menuItems {
		if menuItem.Parent == parentID {
			return true
		}
	}
	return false
}

// IsActive checks if the current URL matches the item URL
func IsActive(currentURL, itemURL string) bool {
	return strings.Trim(currentURL, "/") == strings.Trim(itemURL, "/")
}

// Dict creates a dictionary from key-value pairs
func Dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: missing key or value")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
