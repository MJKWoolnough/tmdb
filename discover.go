package tmdb

import "net/url"

// DiscoverMovie search the movie db filtered by the specified options.
func (t *TMDB) DiscoverMovie(params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)

	if err := t.get(s, "/3/discover/movie", url.Values{}, params...); err != nil {
		return nil, err
	}

	return s, nil
}

// DiscoverTV search the tv db filtered by the specified options.
func (t *TMDB) DiscoverTV(params ...option) (*SearchTV, error) {
	s := new(SearchTV)

	if err := t.get(s, "/3/discover/movie", url.Values{}, params...); err != nil {
		return nil, err
	}

	return s, nil
}
