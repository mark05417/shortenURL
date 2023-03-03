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

type Store interface {
	Save(url URL) string
	Retrieve(short string) (URL, bool)
	ListURLs() (data []URL)
	DeleteURLs()
}
