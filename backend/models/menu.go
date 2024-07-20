// backend/models/menu.go
package models

import (
	"fmt"
	"golang-boilerplate/config"
	"strings"
)

// type Menu struct {
// 	ID           int    `json:"id"`
// 	Name         string `json:"name"`
// 	Link         string `json:"link"`
// 	Parent       int    `json:"parent"`
// 	Icon         string `json:"icon"`
// 	DisplayOrder int    `json:"display_order"`
// 	SubMenus     []Menu `json:"sub_menus,omitempty"`
// }

type Menu struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Parent       int    `json:"parent"`
	Icon         string `json:"icon"`
	DisplayOrder int    `json:"display_order"`
}

// func GetAllMenus() ([]Menu, error) {
// 	db := config.GetDB()

// 	rows, err := db.Query("SELECT id, name, link, parent, icon, display_order FROM menus ORDER BY display_order")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	menuMap := make(map[int]*Menu)
// 	var menus []Menu

// 	for rows.Next() {
// 		var menu Menu
// 		if err := rows.Scan(&menu.ID, &menu.Name, &menu.Link, &menu.Parent, &menu.Icon, &menu.DisplayOrder); err != nil {
// 			return nil, err
// 		}
// 		menuMap[menu.ID] = &menu
// 		if menu.Parent == 0 {
// 			menus = append(menus, menu)
// 		} else {
// 			parentMenu := menuMap[menu.Parent]
// 			parentMenu.SubMenus = append(parentMenu.SubMenus, menu)
// 		}
// 	}

// 	return menus, nil
// }

// type MenuItem struct {
// 	ID           int
// 	Name         string
// 	Link         string
// 	Parent       int
// 	Icon         string
// 	DisplayOrder int
// }

func GetMenuItems() ([]Menu, error) {
	db := config.GetDB()
	rows, err := db.Query("SELECT id, name, link, parent, icon, display_order FROM menus ORDER BY parent, display_order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuItems []Menu
	for rows.Next() {
		var menuItem Menu
		if err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Link, &menuItem.Parent, &menuItem.Icon, &menuItem.DisplayOrder); err != nil {
			return nil, err
		}
		menuItems = append(menuItems, menuItem)
	}
	return menuItems, nil
}

func HasChildren(menuItems []Menu, parentID int) bool {
	for _, menuItem := range menuItems {
		if menuItem.Parent == parentID {
			return true
		}
	}
	return false
}

func IsActive(currentURL, itemURL string) bool {
	return strings.Trim(currentURL, "/") == strings.Trim(itemURL, "/")
}

// Custom function to create a dictionary
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
