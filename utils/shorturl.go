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
	Store map[string]URL
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
