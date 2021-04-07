// Package tmdb is a simple interface to the themoviedb.org database
package tmdb // import "vimagination.zapto.org/tmdb"

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var (
	contentType   = []string{"application/json;charset=utf-8"}
	defaultURL, _ = url.Parse("https://api.themoviedb.org/")
)

// APIKey is used to set the API key for a request
type APIKey interface {
	setAPIKey(r *http.Request, params url.Values)
}

type v4Key struct {
	headers http.Header
}

func (v v4Key) setAPIKey(r *http.Request, _ url.Values) {
	r.Header = v.headers
}

// V4Key creates a new TMDB APIv4 key
func V4Key(v4key string) APIKey {
	return &v4Key{
		headers: http.Header{
			"Authorization": []string{"Bearer " + v4key},
			"Content-Type":  contentType,
		},
	}
}

// TMDB holds an API Key to allow connection to TMDB
type TMDB struct {
	apiKey APIKey
}

// New takes a TMDB API key to create a TMDB connection
func New(apiKey APIKey) *TMDB {
	return &TMDB{apiKey}
}

// Error represents an error repsonse (401, 404) from the API
type Error struct {
	Message string `json:"status_message"`
	Code    int64  `json:"status_code"`
}

func (e *Error) Error() string {
	return e.Message
}

func (t *TMDB) get(result interface{}, path string, query url.Values, params ...option) error {
	for _, param := range params {
		param.setParam(query)
	}
	var req http.Request
	t.apiKey.setAPIKey(&req, query)
	url := defaultURL
	url.Path = path
	url.RawQuery = query.Encode()
	req.URL = url
	r, err := http.DefaultClient.Do(&req)
	if err != nil {
		return err
	}
	if r.StatusCode == http.StatusUnauthorized || r.StatusCode == http.StatusNotFound {
		e := new(Error)
		err = json.NewDecoder(r.Body).Decode(result)
		r.Body.Close()
		if err != nil {
			return err
		}
		return e
	}
	err = json.NewDecoder(r.Body).Decode(result)
	r.Body.Close()
	return err
}
