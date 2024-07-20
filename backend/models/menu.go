// backend/models/menu.go
package models

import "golang-boilerplate/config"

type Menu struct {
	ID           int
	Name         string
	Link         string
	Parent       int
	Icon         string
	DisplayOrder int
	Children     []Menu // Submenu items (if any)
}

// GetMenus retrieves all menus from the database
func GetMenus() ([]Menu, error) {

	db := config.GetDB()
	rows, err := db.Query("SELECT id, name, link, parent, icon, display_order, created_at FROM menus ORDER BY display_order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []Menu
	for rows.Next() {
		var menu Menu
		if err := rows.Scan(&menu.ID, &menu.Name, &menu.Link, &menu.Parent, &menu.Icon, &menu.DisplayOrder); err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return menus, nil
}
