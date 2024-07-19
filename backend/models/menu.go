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

func GetMenuData() ([]Menu, error) {

	db := config.GetDB()

	rows, err := db.Query("SELECT id, name, link, parent, icon, display_order FROM menus ORDER BY parent, display_order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Menu
	itemMap := make(map[int]*Menu)

	for rows.Next() {
		var item Menu
		if err := rows.Scan(&item.ID, &item.Name, &item.Link, &item.Parent, &item.Icon, &item.DisplayOrder); err != nil {
			return nil, err
		}

		itemMap[item.ID] = &item
		if item.Parent == 0 {
			items = append(items, item)
		} else {
			parentItem := itemMap[item.Parent]
			parentItem.Children = append(parentItem.Children, item)
		}
	}
	return items, nil
}
