package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Driver imported for mysql connection.
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

// Client is a database connection.
var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured.")
}
