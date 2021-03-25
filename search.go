package tmdb

import (
	"net/url"
)

// SearchTV is the results returned from a TV search
type SearchTV struct {
	Page    uint16 `json:"page"`
	Results []struct {
		PosterPath       *string  `json:"poster_path"`
		Popularity       float64  `json:"popularity"`
		ID               int64    `json:"id"`
		BackdropPath     string   `json:"backdrop_path"`
		VoteAverage      float64  `json:"vote_average"`
		Overview         string   `json:"overview"`
		FirstAirDate     string   `json:"first_air_date"`
		OriginCountry    []string `json:"origin_country"`
		GenreIDs         []int64  `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		VoteCount        int64    `json:"vote_count"`
		Name             string   `json:"name"`
		OriginalName     string   `json:"original_name"`
	}
	TotalResults uint64 `json:"total_results"`
	TotalPages   uint64 `json:"total_pages"`
}

// SearchTV searches the TMDB TV show database for the name given
func (t *TMDB) SearchTV(name string, params ...option) (*SearchTV, error) {
	s := new(SearchTV)
	if err := t.get("/3/search/tv", url.Values{"query": []string{name}}, params, s); err != nil {
		return nil, err
	}
	return s, nil

}
