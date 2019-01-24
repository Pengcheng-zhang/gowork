package services

import(
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

var redisConnect redis.Conn
func init() {
	redisConnect, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		Debug("connect to redis error", err)
		return
	}
	defer redisConnect.Close()
}

func GetRedisInstance() redis.Conn{
	return redisConnect;
}

func Set(key, value string) error{
	_, err := redisConnect.Do("SET", key, value)
	if err != nil {
		Debug("redis set failed:", key, err)
	}
	return nil
}

func SetNX(key string, value interface{}) error{
	jsonValue, _ := json.Marshal(value)

	n, err := redisConnect.Do("SETNX", key, jsonValue)
	if err != nil {
		Debug("redis setnx failed:", err)
	}

	if n == 1 {
		Debug("redis setnx success")
	}
	return nil
}

func Get(key string) (string, error) {
	value, err := redis.String(redisConnect.Do("GET", key))
	if err != nil {
		Debug("redis get failed:", key, err)
	}
	return value, err
}

func GetNX(key string) (interface{}, error){
	valueGet, err := redis.Bytes(redisConnect.Do("GET", key))
	if err != nil {
		Debug("redis getnx failed:", key, err)
	}
	var imapGet map[string]string
	errSha1 := json.Unmarshal(valueGet, &imapGet)
	if errSha1 != nil {
		Debug("redis getnx unmarsha1 failed:", valueGet)
	}
	return imapGet, nil
}

func Del(key string) error {
	_, err := redisConnect.Do("DEL", key)
	if err != nil {
		Debug("redis delete failed", key, err)
	}
	return nil
}