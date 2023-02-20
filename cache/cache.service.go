package cache

import (
	"encoding/json"
	"log"
	"time"
)

func GetFromCache(key string, toValue any) error {
	val, err := RedisConnection.Client.Get(RedisConnection.Context, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), toValue)
}

func SaveToCache(key string, value interface{}, ttl time.Duration) {
	valueToSave, err := json.Marshal(value)
	if err != nil {
		log.Printf("PARSE JSON ERROR key %s ==> %v\n", key, err)
		return
	}

	_, err = RedisConnection.Client.Set(RedisConnection.Context, key, valueToSave, ttl).Result()
	if err != nil {
		log.Printf("REDIS ERROR [SET] key: %s value: %v ==> %v\n", key, value, err)
		return
	}
}
