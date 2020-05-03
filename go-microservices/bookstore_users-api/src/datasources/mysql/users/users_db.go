package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlDbUsername = "mysql_db_username"
	mysqlDbPassword = "mysql_db_password"
	mysqlDbHost     = "mysql_db_host"
	mysqlDbSchema   = "mysql_db_schema"
)

var (
	// Client mysql
	Client *sql.DB

	username = os.Getenv(mysqlDbUsername)
	password = os.Getenv(mysqlDbPassword)
	host     = os.Getenv(mysqlDbHost)
	schema   = os.Getenv(mysqlDbSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
