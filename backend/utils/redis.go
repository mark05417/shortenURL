package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

type ShortURLStoreRedis struct {
	DB  *redis.Client
	Ctx context.Context
}

func (s *ShortURLStoreRedis) Save(url URL) string {
	short := GenerateRandomShortURL()
	url.Short = short
	s.DB.Set(short, url.Original, 0)
	return short
}

func (s *ShortURLStoreRedis) Retrieve(short string) (URL, bool) {
	original, err := s.DB.Get(short).Result()
	fmt.Println(short, original, err)
	if err != nil {
		return URL{Short: short}, false
	}
	return URL{Short: short, Original: original}, true
}

func (s *ShortURLStoreRedis) ListURLs() (data []URL) {
	keys, _ := s.DB.Keys("*").Result()
	for _, key := range keys {
		original, _ := s.DB.Get(key).Result()
		data = append(data, URL{Short: key, Original: original})
	}
	return
}

func (s *ShortURLStoreRedis) DeleteURL(short string) {
	s.DB.Del(short)
	return
}

func (s *ShortURLStoreRedis) DeleteURLs() {
	s.DB.FlushAll()
	return
}

// func (s *ShortURLStoreRedis) redisDecreaseByOne(c *gin.Context) {
// 	lock := redisLock.NewRedisLock(s.DB, "rlkey", "rlkey")
// 	lock.Lock()
// 	defer lock.Unlock()

// 	val, _ := s.DB.Get("stock").Result()
// 	stockVal, _ := strconv.Atoi(val)
// 	if stockVal > 0 {
// 		stockVal -= 1
// 		s.DB.Set("stock", stockVal, 0)
// 		fmt.Println("Success decrease stock by 1, it becomes ", stockVal)
// 	} else {
// 		fmt.Println("Fail to decrease stock by 1, no enough stocks.")
// 	}
// }
