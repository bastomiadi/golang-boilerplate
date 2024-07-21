<p align="center"><img src="screenshot/golang.png?raw=true" width="200"></p>

Golang Boilerplate
=====================================
This package for Golang Boilerplate serves as a basic platform for quickly creating a back-office application. It includes profile creation and management, user management, roles, permissions and a dynamically-generated menu.

Feature
-------
* Configurable backend theme AdminLTE 3
* Ajax Crud With Datatable, Modal and Sweetalert
* backend, Frontend, Api (With Versioning) Modular
* etc.

This project is still early in its development... please feel free to contribute!
------------------------------------------------------------
Screenshoot |
-------------------------------------------------------------------------------
![Login](screenshot/web/login.png?raw=true)
![Register](screenshot/web/register.png?raw=true)
![Dashboard](screenshot/web/dashboard.png?raw=true)

Installation
------------

**1.** Get The Repository

```bash
E:\Bastomi\Projects\Golang\project\>git clone https://github.com/bastomiadi/golang-boilerplate.git
```

**2.** Run development server:

```bash
E:\Bastomi\Projects\Golang\project\golang-boilerplate>go run main.go
2024/07/20 23:58:08 Connected to the database
2024/07/20 23:58:08 Dropped categories table if it existed
2024/07/20 23:58:08 Created categories table
2024/07/20 23:58:08 Dropped products table if it existed
2024/07/20 23:58:08 Created products table
2024/07/20 23:58:08 Dropped user_roles table if it existed
2024/07/20 23:58:08 Dropped users table if it existed
2024/07/20 23:58:08 Created users table
2024/07/20 23:58:08 Dropped menus table if it existed
2024/07/20 23:58:08 Created menus table
2024/07/20 23:58:08 Dropped roles permissions table if it existed
2024/07/20 23:58:08 Dropped roles table if it existed
2024/07/20 23:58:08 Created roles table
2024/07/20 23:58:09 Dropped permissions table if it existed
2024/07/20 23:58:09 Created permissions table
2024/07/20 23:58:09 Dropped role_permissions table if it existed
2024/07/20 23:58:09 Created role_permissions table
2024/07/20 23:58:09 Dropped user_roles table if it existed
2024/07/20 23:58:09 Created user_roles table
2024/07/20 23:58:09 Seeded categories table
2024/07/20 23:58:09 Seeded products table
2024/07/20 23:58:09 Seeded users table with sample data
2024/07/20 23:58:09 Seeded menus table
2024/07/20 23:58:09 Seeded roles table
2024/07/20 23:58:09 Seeded permissions table
2024/07/20 23:58:09 Seeded user roles table
2024/07/20 23:58:09 Seeded role permissions table
2024/07/20 23:58:09 Server started on: http://localhost:8080

```

**3.** Open in browser http://localhost:8080/backend/login
```bash
Default user and password
+----+------------------+-------------+
| No |       User       |   Password  |
+----+------------------+-------------+
| 1  | admin@admin.com  |   password  |
| 2  | user@user.com    |   password  |
+----+------------------+-------------+
```

Usage
-----
You can find how it works with the read code, controller and views etc. Finnally... Happy Coding!
..

Restful Api and Docs Work in progress.. :  


Changelog
--------
Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

Contributing
------------
Contributions are very welcome.

License
-------

This package is free software distributed under the terms of the [MIT license](LICENSE.md).