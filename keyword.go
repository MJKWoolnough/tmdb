package tmdb

import (
	"fmt"
	"net/url"
)

// Keyword contains the ID and Name of a keyword
type Keyword struct {
	ID   int64  `json:"id"`
	Name string `json:"string"`
}

// Keyword returns the Keyword specified by the ID given
func (t *TMDB) Keyword(id int64) (*Keyword, error) {
	k := new(Keyword)
	if err := t.get(k, fmt.Sprintf("/3/keyword/%d", id), url.Values{}); err != nil {
		return nil, err
	}
	return k, nil
}
