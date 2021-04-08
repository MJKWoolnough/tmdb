package tmdb

import "net/url"

// TrendingDay retrieves a list of trending items over the last day
func (t *TMDB) TrendingDay() (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/trending/all/day", url.Values{}); err != nil {
		return nil, err
	}
	return s, nil
}

// TrendingWeek retrieves a list of trending items over the last week
func (t *TMDB) TrendingWeek() (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/trending/all/week", url.Values{}); err != nil {
		return nil, err
	}
	return s, nil
}

// TrendingMoviesDay retrieves a list of trending movies over the last day
func (t *TMDB) TrendingMoviesDay() (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/trending/movie/day", url.Values{}); err != nil {
		return nil, err
	}
	return s, nil
}
