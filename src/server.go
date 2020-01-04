package main

import (
	"os"
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	// routes "../src/controller"
	// response "../src/views"
	sqlconnect "../src/model/sql"
	redisconnect "../src/model/redis"
	cachehandler "../src/cachehandler"
	// "reflect"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"
	"strconv"
)

const (
	mySQLConnString   = "root:root@tcp(localhost:3307)/goLangExperiment"
	mySQLMaxConnCount = 40
)

func main() {
	// initialize redis pool and bootstrap redis
	redisconnect.InitRedis()
  
  	// get value which exists
	log.Printf(redisconnect.Get("00000E"))
  
  	// get value which does not exists
	log.Printf(redisconnect.Get("0000E"))

  	// add members to set
	redisconnect.Sadd("mystiko", "0000E")
	redisconnect.Sadd("mystiko", "0000D")
  
  	// get memebers of set
	s, _ := redisconnect.Smembers("mystiko")
	log.Println("%v", s)

	// /*------- Server Port ----------*/
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8082"
	}
	/*---------------------------------*/

	routee := fasthttprouter.New()

	// routee.GET("/", routes.Index)
	// routee.POST("/sht/", routes.GetShortenedURL)
	// routee.GET("/ext/:id" , routes.GetExtendedURL)

	routee.GET("/", Index)
	routee.POST("/sht/", GetShortenedURL)
	routee.GET("/ext/:id" , GetExtendedURL)

	if e := fasthttp.ListenAndServe(serverPort, routee.Handler); e != nil {
		fmt.Println(e.Error())
	}
}

func Index(ctx *fasthttp.RequestCtx){
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("It's working fine !"))
}

func GetExtendedURL(ctx *fasthttp.RequestCtx){
	_hash := "dhegehgefeffe"
	_response = cachehandler.GetConfiguration(_hash)
	print(_response)
}

func GetShortenedURL(ctx *fasthttp.RequestCtx){
	fmt.Println("-----333333-------")
	// fmt.Println(ctx)
	// req := &ctx.Request
	// fmt.Println(ctx.PostBody())
	// fmt.Println(reflect.TypeOf(req))

	originalURL := "http://www.workindia.in"

	var (
		short_url string
	)
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	for loopI := 0; loopI < 5; loopI++ {
		_hash := createHashString(originalURL)
		_obj = cachehandler.GetConfiguration(_hash)
		if _obj == nil{
			sqldb := sqlconnect.SQLConnect()
			defer sqldb.Close()
			sqlconnect.SQLAdd(sqldb, _hash, originalURL)
			cachehandler.SetConfiguration(_hash, originalURL)
			short_url = _hash
			break
		}
	}
	ctx.SetContentType("application/json")
	if short_url != nil{
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte(short_url))
	}
	else{
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}

func createHashString(originalURL string) string{
	hashHelper := sha1.New()
	hashHelper.Write([]byte(originalURL))
	sha1_hash := hex.EncodeToString(hashHelper.Sum(nil))
	sha1_hash = takerandonSubstring(sha1_hash)
	return sha1_hash
}

func takerandonSubstring(str string) string{
	_start := rand.Intn(4)
	_end := _start + 5
	str = str[_start:_end]
	_datetime_now := time.Now()
	str = strconv.Itoa(int(_datetime_now.Unix())) + str
	return str
}