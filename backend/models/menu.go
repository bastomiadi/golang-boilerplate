// backend/models/menu.go
package models

import "database/sql"

type Menu struct {
	ID           int
	Name         string
	Link         string
	Parent       int
	Icon         string
	DisplayOrder int
	Children     []Menu // Submenu items (if any)
}

// func GetMenus(db *sql.DB) ([]Menu, error) {
// 	rows, err := db.Query("SELECT id, name, url, parent_id, icon, `order` FROM menu ORDER BY `order`")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	menus := []Menu{}
// 	for rows.Next() {
// 		var menu Menu
// 		err := rows.Scan(&menu.ID, &menu.Name, &menu.Link, &menu.Parent, &menu.Icon, &menu.DisplayOrder)
// 		if err != nil {
// 			return nil, err
// 		}
// 		menus = append(menus, menu)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return menus, nil
// }

func FetchMenus(db *sql.DB) ([]Menu, error) {
	query := `
        SELECT id, name, link, parent, icon, display_order
        FROM menus
        ORDER BY display_order
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map to store parent items temporarily
	menuMap := make(map[int]*Menu)
	var Menus []Menu

	// Read rows and populate menu items
	for rows.Next() {
		var Menu Menu
		err := rows.Scan(&Menu.ID, &Menu.Name, &Menu.Link, &Menu.Parent, &Menu.Icon, &Menu.DisplayOrder)
		if err != nil {
			return nil, err
		}

		// Check if parent exists in map, otherwise create it
		if _, ok := menuMap[Menu.ID]; !ok {
			menuMap[Menu.ID] = &Menu
		} else {
			// If parent exists, add as child
			parent := menuMap[Menu.ID]
			parent.Children = append(parent.Children, Menu)
		}

		// Add top-level items to Menus
		if Menu.Parent == 0 {
			Menus = append(Menus, Menu)
		}
	}

	return Menus, nil
}
