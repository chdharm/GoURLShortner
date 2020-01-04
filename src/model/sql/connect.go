package main

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

func SQLGet(conn *sql.DB, _hash string) (string, error){
	queryString := "SELECT * FROM URLShortner where hash=?;"
	selDB := conn.QueryRow(queryString, _hash)
	var (
		hash string
		originalurl string
	)
	_err := selDB.Scan(&hash, &originalurl)
	if _err != nil {
		return "", _err
	}
	fmt.Println(hash)
	fmt.Println(originalurl)
	return originalurl, nil
}

func SQLAdd(conn *sql.DB, originalUrl string, _hash string) error{
	queryString := "INSERT INTO URLShortner VALUES"
	queryString += "('" + _hash + "','" + originalUrl + "');"
	fmt.Println(queryString)
	selDB, err := conn.Query(queryString)
	fmt.Println(selDB)
	if err != nil {
		return err
	}
	return nil
}

func main(){
	fmt.Println("hiiiis")
	SQLConnect()
	//SQLAdd(sqldb,"http://dharmpal.com/IN", "fhjmnfvdghnbBFVNSVB V SNBVSBNSV NVNBvdshds")
	SQLGet(sqldb, "fhjmnfvdghnbvdshds")

}