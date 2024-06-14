package tmdb

import "net/url"

// Change represents the ID of an item change.
type Change struct {
	ID    int64 `json:"id"`
	Adult *bool `json:"adult"`
}

// Changes lists results from a change search.
type Changes struct {
	Results      []Change `json:"results"`
	Page         int64    `json:"page"`
	TotalPages   int64    `json:"total_pages"`
	TotalResults int64    `json:"total_results"`
}

// ChangesMovie returns all of the Movie changes within the last 24 hours, filtered by the given params.
func (t *TMDB) ChangesMovie(params ...option) (*Changes, error) {
	c := new(Changes)

	if err := t.get(c, "/3/movie/changes", url.Values{}, params...); err != nil {
		return nil, err
	}

	return c, nil
}

// ChangesTV returns all of the TV changes within the last 24 hours, filtered by the given params.
func (t *TMDB) ChangesTV(params ...option) (*Changes, error) {
	c := new(Changes)

	if err := t.get(c, "/3/tv/changes", url.Values{}, params...); err != nil {
		return nil, err
	}

	return c, nil
}

// ChangesPerson returns all of the Person changes within the last 24 hours, filtered by the given params.
func (t *TMDB) ChangesPerson(params ...option) (*Changes, error) {
	c := new(Changes)

	if err := t.get(c, "/3/person/changes", url.Values{}, params...); err != nil {
		return nil, err
	}

	return c, nil
}
