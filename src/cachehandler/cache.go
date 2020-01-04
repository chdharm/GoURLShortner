package main

import (
	sqlconnect "../../src/model/sql"
	redisconnect "../../src/model/redis"
	"fmt"
)

func SetConfiguration(_hash string, originalURL string){
	redisconnect.Sadd(_hash, originalURL)
}

func GetConfiguration(_hash string){
	_obj, _err := redisconnect.Get(_hash)
	if _err != nil{
		fmt.Println("_err===", _err)
		// return _err
	}
	print(_obj)
// 	if _obj != nil{
// 		return _obj
// 	}
	sqldb := sqlconnect.SQLConnect()
	defer sqldb.Close()
// 	_response = sqlconnect.SQLGet(sqldb, _hash)
// 	originalURL = response
// 	redisconnect.Sadd(_hash, originalURL)
// 	return _response
}

// func InvalidateCache(_hash string){
// 	redisconnect.Del(_hash)
// }

func main(){
	fmt.Println("Hi=====")
	redisconnect.InitRedis()
	SetConfiguration("Ram", "Sita")
	GetConfiguration("Ram")
	// InvalidateCache("Ram")
}