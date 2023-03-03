package utils

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

type ShortURLStoreRedis struct {
	DB  *redis.Client
	Ctx context.Context
}

func (s *ShortURLStoreRedis) Save(url URL) string {
	short := GenerateRandomShortURL()
	url.Short = short
	s.DB.Set("url:"+short, url.Original, 0)
	s.DB.Set("count:"+short, 0, 0)
	return short
}

func (s *ShortURLStoreRedis) Retrieve(short string) (URL, bool) {
	original, err := s.DB.Get("url:" + short).Result()
	if err != nil {
		return URL{Short: short}, false
	}
	return URL{Short: short, Original: original}, true
}

func (s *ShortURLStoreRedis) IncreaseCount(short string) {
	countStr, _ := s.DB.Get("count:" + short).Result()
	count, _ := strconv.Atoi(countStr)
	count++
	s.DB.Set("count:"+short, count, 0)
	return
}

func (s *ShortURLStoreRedis) ListURLs() (data []URL) {
	keys, _ := s.DB.Keys("url:*").Result()
	for _, key2 := range keys {
		key := strings.Split(key2, "url:")[1]

		original, _ := s.DB.Get("url:" + key).Result()
		countStr, _ := s.DB.Get("count:" + key).Result()
		count, _ := strconv.Atoi(countStr)
		data = append(data, URL{Short: key, Original: original, Count: count})
	}
	return
}

func (s *ShortURLStoreRedis) DeleteURL(short string) {
	s.DB.Del("url:" + short)
	s.DB.Del("count:" + short)
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
