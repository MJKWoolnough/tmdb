package tmdb

import (
	"fmt"
	"net/http"
	"net/url"
)

var (
	contentType   = []string{"application/json;charset=utf-8"}
	defaultURL, _ = url.Parse("https://api.themoviedb.org/")
)

// TMDB holds an APIv4 key for themoviedb.org
type TMDB struct {
	headers http.Header
}

// New takes a TMDB APIv4 key
func New(v4key string) *TMDB {
	return &TMDB{
		headers: http.Header{
			"Authorization": []string{"Bearer " + v4key},
			"Content-Type":  contentType,
		},
	}
}

func (t *TMDB) get(path string, query url.Values) (*http.Response, error) {
	url := defaultURL
	url.Path = path
	url.RawQuery = query.Encode()
	fmt.Println(url)
	return http.DefaultClient.Do(&http.Request{
		URL:    url,
		Header: t.headers,
	})
}
