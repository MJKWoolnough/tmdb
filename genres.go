package tmdb

import "net/url"

// Genres is a list of genres Names and IDs
type Genres []struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// MovieGenres retrieves the official genres for movies
func (t *TMDB) MovieGenres(params ...option) (Genres, error) {
	g := new(Genres)
	if err := t.get(g, "/3/genre/movie/list", url.Values{}, params...); err != nil {
		return nil, err
	}
	return *g, nil
}
