package models

import (
	"os"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/logger"
	"github.com/gomodule/redigo/redis"
)

var RedisServer = os.Getenv("REDIS_SERVER")
var RedisPassword = envy.Get("REDIS_PASSWORD", "")
var appLogger logger.FieldLogger
var pool *redis.Pool

func init() {
	pool = NewPool()
}

func NewPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisServer)
			if err != nil {
				panic(err.Error())
			}

			if RedisPassword != "" {
				_, err = c.Do("AUTH", RedisPassword)
				if err != nil {
					panic(err.Error())
				}
			}

			return c, err
		},
	}
}

func GetConn() redis.Conn {
	conn := pool.Get()
	return conn
}

// SetInt function to save integer value in redis key
func SetInt(key string, value int) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetInt function to fetch integer value from redis key
func GetInt(key string) int {
	conn := GetConn()
	defer conn.Close()

	value, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		return 0
	}

	return value
}

// SetInt64 function to save int64 value in redis key
func SetInt64(key string, value int64) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetInt64 function to fetch int64 value from redis key
func GetInt64(key string) int64 {
	conn := GetConn()
	defer conn.Close()

	value, err := redis.Int64(conn.Do("GET", key))
	if err != nil {
		return 0
	}

	return value
}

// SetString function to save string value in redis key
func SetString(key string, value string) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetString function to fetch string value from redis key
func GetString(key string) string {
	conn := GetConn()
	defer conn.Close()

	value, _ := redis.String(conn.Do("GET", key))

	return value
}

// AddMember function to add value to redis array
func AddMember(key string, value string) error {
	conn := GetConn()
	defer conn.Close()
	_, err := conn.Do("SADD", key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetMembers function to fetch array of values
func GetMembers(key string) []string {
	conn := GetConn()
	defer conn.Close()
	values, _ := redis.Strings(conn.Do("SMEMBERS", key))

	return values
}

// RemoveMember function to delete a value from array
func RemoveMember(key string, value string) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do("SREM", key, value)
	if err != nil {
		return err
	}
	return nil
}

// IsMember function to check whether value present in the array or not
func IsMember(key string, value string) bool {
	conn := GetConn()
	defer conn.Close()

	result, err := redis.Bool(conn.Do("SISMEMBER", key, value))
	if err != nil {
		return false
	}
	return result
}

// DeleteKey function to delete mentioned key from redis server
func DeleteKey(key string) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do("DEL")
	if err != nil {
		return err
	}
	return nil
}
