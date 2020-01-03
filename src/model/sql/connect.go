package sqlconnect

import (
	// "os"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mySQLConnString   = "root:root@tcp(localhost:3307)/URLShortner"
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

func SQLGet(conn *sql.DB, hash string){
	queryString := "SELECT SHORTENEDURL FROM URLShortner where hash= "
	queryString += hash
	selDB, err := conn.Query(queryString)


}

func SQLAdd(conn *sql.DB, originalUrl string){
	//Todo: Create hash here
	hash := createHash(originalUrl)
	queryString := "INSER INTO SHORTENEDURL(HASH, ORIGINALURL) VALUES"
	queryString += "(" + hash + "," + originalUrl + ")"
	selDB, err := conn.Query(queryString)
	print("selDB:", selDB)
}

func createHash(url string) string{
	hash := ""
	//take time stamp and create hash
	return hash
}
func mustPrepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Error when preparing statement %q: %s", query, err)
	}
	return stmt
}
