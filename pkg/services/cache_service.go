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

func SetCache(key string, i interface{}) {
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
}

func GetCache(cacheKey string, i interface{}) string {
	cacheData, err := cache.Get(ctx, cacheKey).Result()
	if err != nil && err != redis.Nil {
		logger.Error("Error getting cache", err)
	}

	if cacheData != "" {
		err = json.Unmarshal([]byte(cacheData), i)
		if err != nil {
			logger.Error("Error unmarshalling cache", err)
		}
	}
	return cacheData
}

func DeleteCache(cacheKeys []string, flushall bool) {
	for _, cacheKey := range cacheKeys {
		err := cache.Del(ctx, cacheKey).Err()
		if err != nil && err != redis.Nil {
			logger.Error("Error deleting cache", err)
		}
	}

	if flushall {
		err := cache.FlushAll(ctx).Err()
		if err != nil && err != redis.Nil {
			logger.Error("Error flushing cache", err)
		}
	}
}
