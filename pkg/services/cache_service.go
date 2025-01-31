package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/configs/logger"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var cache *redis.Client

func StartCache() {
	cacheConf := config.GetConfig().CACHE

	url := fmt.Sprintf("redis://%s:%s", cacheConf.Addr, cacheConf.Port)
	newCache, err := redis.ParseURL(url)
	if err != nil {
		logger.Error("Error parsing cache URL", err)
	}

	cache = redis.NewClient(newCache)

	if err := cache.Ping(ctx).Err(); err != nil {
		logger.Error("Error connecting to cache", err)
		return
	}

	log.Printf("Connected to cache at port: %s", cacheConf.Port)
}

func SetCache(key string, i interface{}) interface{} {
	var ttl = config.GetConfig().CACHE.ExpiresIn

	cacheValue, err := json.Marshal(i)
	if err != nil {
		logger.Error("Error marshalling cache", err)
	} else {
		err := cache.Set(ctx, key, string(cacheValue), ttl*time.Second).Err()
		if err != nil {
			logger.Error("Error setting cache", err)
		}
	}

	return cacheValue
}

func GetCache(key string, i interface{}) string {
	cacheData, err := cache.Get(ctx, key).Result()

	if err != nil {
		if err != redis.Nil {
			logger.Error("Error getting cache", err)
		}
	}

	if cacheData != "" {
		err = json.Unmarshal([]byte(cacheData), i)
		if err != nil {
			logger.Error("Error unmarshalling cache", err)
		}
	}
	return cacheData
}

func DeleteCache(prefix string, key string, flushall bool) {
	err := cache.Del(ctx, prefix+key).Err()
	if err != nil {
		logger.Error("Error deleting cache", err)
	}

	if flushall {
		err := cache.Del(ctx, prefix+"all").Err()
		if err != nil {
			logger.Error("Error flushing cache", err)
		}
	}
}
