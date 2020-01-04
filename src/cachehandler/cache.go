package cachehandler

import (
	sqlconnect "../../src/model/sql"
	redisconnect "../../src/model/redis"
	"fmt"
)

func SetConfiguration(_hash string, originalURL string){
	redisconnect.Sadd(_hash, originalURL)
}

func GetConfiguration(_hash string) string{
	_obj, _err := redisconnect.Get(_hash)
	fmt.Println(_obj)
	fmt.Println(_err)
	if _err == nil{
		return _obj
	}
	sqldb := sqlconnect.SQLConnect()
	defer sqldb.Close()
	_sqlobj, _sqlerr := sqlconnect.SQLGet(sqldb, _hash)
	if _sqlerr != nil{
		return ""
	}
	redisconnect.Sadd(_hash, _sqlobj)
	return _sqlobj
}

// func InvalidateCache(_hash string){
// 	redisconnect.Del(_hash)
// }

// func main(){
// 	fmt.Println("Hi=====")
// 	redisconnect.InitRedis()
// 	SetConfiguration("QAaaaaaaaa", "Sita")
// 	GetConfiguration("QAaaaaaaaa")
// 	// // InvalidateCache("Ram")
// }