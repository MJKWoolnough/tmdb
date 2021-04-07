package tmdb

import "net/url"

// TrendingDay retrieves a list of trending items over the last day
func (t *TMDB) TrendingDay() (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/trending/all/day", url.Values{}); err != nil {
		return nil, err
	}
	return t, nil
}
