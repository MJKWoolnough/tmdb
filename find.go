package tmdb

import (
	"net/url"
)

// FindResults contains the results from the Find call.
type FindResults struct {
	MovieResults  []MovieResult  `json:"movie_results"`
	PersonResults []PeopleResult `json:"person_results"`
	TVResults     []TVResult     `json:"tv_result"`
}

// Find searches the database for movies, people, or TV shows that reference the given external ID.
func (t *TMDB) Find(id, source string, params ...option) (*FindResults, error) {
	f := new(FindResults)

	if err := t.get(f, "/3/find/"+id, url.Values{"external_source": []string{source}}, params...); err != nil {
		return nil, err
	}

	return f, nil
}
