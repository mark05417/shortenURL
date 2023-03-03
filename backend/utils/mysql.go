package utils

import (
	"database/sql"
	"fmt"
)

type ShortURLStoreMysql struct {
	DB *sql.DB
}

func (s *ShortURLStoreMysql) Save(url URL) string {
	short := GenerateRandomShortURL()
	url.Short = short
	s.DB.Exec("INSERT INTO url_mapping (short, original) VALUES (?, ?)", short, url.Original)
	return short
}

func (s *ShortURLStoreMysql) Retrieve(short string) (URL, bool) {
	var originalURL string
	row := s.DB.QueryRow("SELECT original from mydatabase.url_mapping where short = ?", short)
	row.Scan(&originalURL)

	if originalURL == "" {
		return URL{Short: short}, false
	}
	return URL{Short: short, Original: originalURL}, true
}

func (s *ShortURLStoreMysql) ListURLs() (data []URL) {
	rows, err := s.DB.Query("SELECT short, original from mydatabase.url_mapping")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var short, originalURL string
		rows.Scan(&short, &originalURL)
		data = append(data, URL{Original: originalURL, Short: short})
	}
	return
}

func (s *ShortURLStoreMysql) DeleteURL(short string) {
	s.DB.Exec("delete from mydatabase.url_mapping where short = ?", short)
	return
}

func (s *ShortURLStoreMysql) DeleteURLs() {
	s.DB.Exec("delete from mydatabase.url_mapping")
	return
}
