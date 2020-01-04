package sqlconnect

import (
	// "os"
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

func SQLConnect() *sql.DB{
	var sqlerr error
	if sqldb, sqlerr = sql.Open("mysql", mySQLConnString); sqlerr != nil {
		log.Fatalf("Error opening database: %s", sqlerr)
	}
	if sqlerr = sqldb.Ping(); sqlerr != nil {
		log.Fatalf("Cannot connect to db: %s", sqlerr)
	}else{
		fmt.Println("MySQL DB connected")
	}
	sqldb.SetMaxOpenConns(mySQLMaxConnCount)
	sqldb.SetMaxIdleConns(mySQLMaxIdleConnCount)
	return sqldb
}

func SQLGet(conn *sql.DB, hash string) *sql.Rows{
	queryString := "SELECT SHORTENEDURL FROM URLShortner where hash= "
	queryString += hash
	selDB, err := conn.Query(queryString)
	fmt.Println(err)
	return selDB
}

func SQLAdd(conn *sql.DB, originalUrl string, hash string){
	queryString := "INSER INTO SHORTENEDURL(HASH, ORIGINALURL) VALUES"
	queryString += "(" + hash + "," + originalUrl + ")"
	selDB, err := conn.Query(queryString)
	fmt.Println(selDB)
	fmt.Println(err)
}