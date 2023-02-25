package utils

import (
	"math/rand"
	"strings"
	"time"
)

type URL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

type ShortURLStore struct {
	Store map[string]URL // 'MQgJX5' -> {original : 'http:google.com', short : 'MQgJX5'}
}

func (s *ShortURLStore) Save(url URL) string {
	short := GenerateRandomShortURL()
	url.Short = short
	s.Store[short] = url
	return short
}

func (s *ShortURLStore) Retrieve(short string) (URL, bool) {
	url, ok := s.Store[short]
	return url, ok
}

func (s *ShortURLStore) ListURLs() (data []URL) {
	for _, val := range s.Store {
		data = append(data, val)
	}
	return data
}

func (s *ShortURLStore) DeleteURLs() {
	s.Store = map[string]URL{}
}

func GenerateRandomShortURL() string {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const length = 6

	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	return sb.String()
}
