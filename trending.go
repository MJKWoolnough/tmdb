package tmdb

import "net/url"

// TrendingDay retrieves a list of trending items over the last day.
func (t *TMDB) TrendingDay() (*SearchMovie, error) {
	s := new(SearchMovie)

	if err := t.get(s, "/3/trending/all/day", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingWeek retrieves a list of trending items over the last week.
func (t *TMDB) TrendingWeek() (*SearchMovie, error) {
	s := new(SearchMovie)

	if err := t.get(s, "/3/trending/all/week", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingMoviesDay retrieves a list of trending movies over the last day.
func (t *TMDB) TrendingMoviesDay() (*SearchMovie, error) {
	s := new(SearchMovie)

	if err := t.get(s, "/3/trending/movie/day", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingMoviesWeek retrieves a list of trending movies over the last week.
func (t *TMDB) TrendingMoviesWeek() (*SearchMovie, error) {
	s := new(SearchMovie)

	if err := t.get(s, "/3/trending/movie/week", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingTVDay retrieves a list of trending TV shows over the last day.
func (t *TMDB) TrendingTVDay() (*SearchTV, error) {
	s := new(SearchTV)

	if err := t.get(s, "/3/trending/tv/day", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingTVWeek retrieves a list of trending tv shows over the last week.
func (t *TMDB) TrendingTVWeek() (*SearchTV, error) {
	s := new(SearchTV)

	if err := t.get(s, "/3/trending/tv/week", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingPeopleDay retrieves a list of trending people over the last day.
func (t *TMDB) TrendingPeopleDay() (*SearchPeople, error) {
	s := new(SearchPeople)

	if err := t.get(s, "/3/trending/person/day", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}

// TrendingPeopleWeek retrieves a list of trending people over the last week.
func (t *TMDB) TrendingPeopleWeek() (*SearchPeople, error) {
	s := new(SearchPeople)

	if err := t.get(s, "/3/trending/person/week", url.Values{}); err != nil {
		return nil, err
	}

	return s, nil
}
