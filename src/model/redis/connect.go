package redisconnect

import (
	"log"
	"os"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func InitRedis() {
	// init redis connection pool
	InitPool()

	// bootstramp some data to redis
	//InitStore()
}

func InitPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func InitStore() {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	macs := []string{"Testkey1  Cisco", "Testkey2  FIBRONICS", "Testkey3  Fujitsu",
		"Testkey4  Next", "Testkey5  Hughes"}
	for _, mac := range macs {
		pair := strings.Split(mac, "  ")
		Set(pair[0], pair[1])
	}
}

func Ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		os.Exit(1)
	}
}

func Set(key string, val string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("GET", key)
	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return "", err
	}

	return s, nil
}

func Sadd(key string, val string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", key, val)
	if err != nil {
		log.Printf("ERROR: fail add val %s to set %s, error %s", val, key, err.Error())
		return err
	}

	return nil
}

func Smembers(key string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Printf("ERROR: fail get set %s , error %s", key, err.Error())
		return nil, err
	}

	return s, nil
}