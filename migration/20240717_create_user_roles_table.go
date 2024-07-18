// migration/20240717_create_user_roles_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateUserRolesTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS user_roles (
        user_id INT,
        role_id INT,
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (role_id) REFERENCES roles(id),
        PRIMARY KEY (user_id, role_id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create user_roles table: %v", err)
	}
	log.Println("Created user_roles table")
}
