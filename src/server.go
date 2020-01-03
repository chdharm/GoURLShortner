package main

import (
	"os"
	"fmt"
	// "log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/go-redis/redis"
	routes "../src/controller"
	// response "../src/views"
	sqlconnect "../src/model/sql" 
)

const (
	mySQLConnString   = "root:root@tcp(localhost:3307)/goLangExperiment"
	mySQLMaxConnCount = 40
)

var (
	sqldb *sql.DB
)

func main() {
	db := sqlconnect.SQLConnect()
	defer db.Close()

	// /*------- Redis Config ----------*/
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })
	// redisPong, redisErr := redisClient.Ping().Result()
	// if redisErr != nil {
	// 	panic(redisErr)
	// }
	// if redisPong == "PONG" {
	// 	fmt.Println("Redis client connected")
	// }
	// /*------------------------------------*/


	// /*------- MySQL Config ----------*/
	// var sqlerr error
	// if sqldb, sqlerr = sql.Open("mysql", mySQLConnString); sqlerr != nil {
	// 	log.Fatalf("Error opening database: %s", sqlerr)
	// }
	// if sqlerr = sqldb.Ping(); sqlerr != nil {
	// 	log.Fatalf("Cannot connect to db: %s", sqlerr)
	// }else{
	// 	fmt.Println("MySQL DB connected")
	// }
	// /*--------------------------------*/

	selDB, err := db.Query("SELECT SHORTENEDURL FROM URLShortner")

	/*------- Server Port ----------*/
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8081"
	}
	/*---------------------------------*/

	routee := fasthttprouter.New()

	routee.GET("/", routes.Index)
	routee.POST("/sht/", routes.GetShortenedURL)
	routee.GET("/ext/:id" , routes.GetExtendedURL)

	if e := fasthttp.ListenAndServe(serverPort, routee.Handler); e != nil {
		fmt.Println(e.Error())
	}
}