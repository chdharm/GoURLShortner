package sqlconnect

import (
	"os"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mySQLConnString   = "root:root@tcp(localhost:3307)/goLangExperiment"
	mySQLMaxConnCount = 40
	mySQLMaxIdleConnCount = 40
)

var (
	sqldb *sql.DB
)

func Connect() *sql.DB{
	var sqlerr error
	if sqldb, sqlerr = sql.Open("mysql", mySQLConnString); sqlerr != nil {
		log.Fatalf("Error opening database: %s", sqlerr)
	}
	if sqlerr = sqldb.Ping(); sqlerr != nil {
		log.Fatalf("Cannot connect to db: %s", sqlerr)
	}else{
		fmt.Println("MySQL DB connected")
	}
	return sqldb
}