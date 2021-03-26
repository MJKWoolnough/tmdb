package tmdb

import (
	"net/url"
)

// SearchCompany is the results returned from a Company search
type SearchCompany struct {
	Page    uint16 `json:"page"`
	Results []struct {
		ID       int64   `json:"id"`
		LogoPath *string `json:"logo_path"`
		Name     string  `json:"name"`
	} `json:"results"`
	TotalResults uint64 `json:"total_results"`
	TotalPages   uint64 `json:"total_pages"`
}

// SearchCompany searches the TMDB Movie database for the name given
func (t *TMDB) SearchCompany(query string, params ...option) (*SearchCompany, error) {
	s := new(SearchCompany)
	if err := t.get(s, "/3/search/company", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// SearchMovie is the results returned from a Movie search
type SearchMovie struct {
	Page    uint16 `json:"page"`
	Results []struct {
		PosterPath       *string `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float64 `json:"popularity"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
	} `json:"results"`
	TotalResults uint64 `json:"total_results"`
	TotalPages   uint64 `json:"total_pages"`
}

// SearchMovie searches the TMDB Movie database for the name given
func (t *TMDB) SearchMovie(query string, params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/search/movie", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

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
	} `json:"results"`
	TotalResults uint64 `json:"total_results"`
	TotalPages   uint64 `json:"total_pages"`
}

// SearchTV searches the TMDB TV show database for the name given
func (t *TMDB) SearchTV(query string, params ...option) (*SearchTV, error) {
	s := new(SearchTV)
	if err := t.get(s, "/3/search/tv", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}
	return s, nil
}
