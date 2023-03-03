package utils

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

func (s *ShortURLStore) IncreaseCount(short string) {
	url, _ := s.Store[short]
	url.Count++
	s.Store[short] = url
	return
}

func (s *ShortURLStore) ListURLs() (data []URL) {
	for _, val := range s.Store {
		data = append(data, val)
	}
	return data
}

func (s *ShortURLStore) DeleteURL(short string) {
	delete(s.Store, short)
}

func (s *ShortURLStore) DeleteURLs() {
	s.Store = map[string]URL{}
}
