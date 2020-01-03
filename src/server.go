package main

import (
	"os"
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"
	// "routes"
)

const (
	mySQLConnString   = "root:root@tcp(localhost:3307)/goLangExperiment"
	mySQLMaxConnCount = 40
)

var (
	sqldb *sql.DB
)

func main() {

	/*------- Redis Config ----------*/
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisPong, redisErr := redisClient.Ping().Result()
	if redisErr != nil {
		panic(redisErr)
	}
	if redisPong == "PONG" {
		fmt.Println("Redis client connected")
	}
	/*------------------------------------*/


	/*------- MySQL Config ----------*/
	var sqlerr error
	if sqldb, sqlerr = sql.Open("mysql", mySQLConnString); sqlerr != nil {
		log.Fatalf("Error opening database: %s", sqlerr)
	}
	if sqlerr = sqldb.Ping(); sqlerr != nil {
		log.Fatalf("Cannot connect to db: %s", sqlerr)
	}else{
		fmt.Println("MySQL DB connected")
	}
	/*--------------------------------*/


	/*------- Server Port ----------*/
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8081"
	}
	/*---------------------------------*/

	routee := fasthttprouter.New()
	routee.GET("/", index)
	routee.GET("/ext/:id" , getExtendedURL)
	routee.POST("/sht/", getShortenedURL)

	// routee.GET("/", routes.index)
	// routee.GET("/ext/:id" , routes.getExtendedURL)
	// routee.POST("/sht/", routes.getShortenedURL)

	if e := fasthttp.ListenAndServe(serverPort, routee.Handler); e != nil {
		fmt.Println(e.Error())
	}
}

func index(ctx *fasthttp.RequestCtx){
	ctx.SetContentType("foo/bar")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "this is the first part of body\n")
	ctx.Response.Header.Set("Foo-Bar", "baz")
	fmt.Fprintf(ctx, "this is the 1111111 second part of body\n")
	ctx.SetBody([]byte("this is completely 1111111 new body contents"))
	ctx.SetStatusCode(fasthttp.StatusNotFound)
}
func getExtendedURL(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("foo/bar")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "this is the first part of body\n")
	ctx.Response.Header.Set("Foo-Bar", "baz")
	fmt.Fprintf(ctx, "this is the 22222222 second part of body\n")
	ctx.SetBody([]byte("this is completely 22222222 new body contents"))
	ctx.SetStatusCode(fasthttp.StatusNotFound)
}
func getShortenedURL(ctx *fasthttp.RequestCtx){
	ctx.SetContentType("foo/bar")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "this is the first part of body\n")
	ctx.Response.Header.Set("Foo-Bar", "baz")
	fmt.Fprintf(ctx, "this is the 3333333  second part of body\n")
	ctx.SetBody([]byte("this is completely 3333333 new body contents"))
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	
}