// migration/20240717_create_user_roles_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateUserRolesTable(db *sql.DB) {
	dropQuery := `
    DROP TABLE IF EXISTS user_roles;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS user_roles (
        user_id INT,
        role_id INT,
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (role_id) REFERENCES roles(id),
        PRIMARY KEY (user_id, role_id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop user_roles table: %v", err)
	}
	log.Println("Dropped user_roles table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create user_roles table: %v", err)
	}
	log.Println("Created user_roles table")
}
