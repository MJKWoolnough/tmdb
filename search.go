package tmdb

import (
	"net/url"
)

// Search stores the pages fields for a returned search.
type Search struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
}

// CompanyResult stores a single result of a company search.
type CompanyResult struct {
	ID       int64   `json:"id"`
	LogoPath *string `json:"logo_path"`
	Name     string  `json:"name"`
}

// SearchCompany is the results returned from a Company search.
type SearchCompany struct {
	Search
	Results []CompanyResult `json:"results"`
}

// SearchCompany searches the TMDB Movie database for the name given.
func (t *TMDB) SearchCompany(query string, params ...option) (*SearchCompany, error) {
	s := new(SearchCompany)

	if err := t.get(s, "/3/search/company", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}

	return s, nil
}

// CollectionResult stores a single result of a collection search.
type CollectionResult struct {
	ID           int64   `json:"id"`
	BackdropPath *string `json:"backdrop_path"`
	Name         string  `json:"name"`
	PosterPath   *string `json:"poster_path"`
}

// SearchCollection is the results returned from a Collection search.
type SearchCollection struct {
	Search
	Results []CollectionResult `json:"results"`
}

// SearchCollection searches the TMDB Movie database for the name given.
func (t *TMDB) SearchCollection(query string, params ...option) (*SearchCollection, error) {
	s := new(SearchCollection)

	if err := t.get(s, "/3/search/collection", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}

	return s, nil
}

// SearchKeywords is the results returned from a Keyword search.
type SearchKeywords struct {
	Search
	Results []Keyword `json:"results"`
}

// SearchKeywords search the TMDB database for the tersm given.
func (t *TMDB) SearchKeywords(query string, params ...option) (*SearchKeywords, error) {
	s := new(SearchKeywords)

	if err := t.get(s, "/3/search/keyword", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}

	return s, nil
}

// MovieResult stores a single result of a movie search.
type MovieResult struct {
	PosterPath       *string `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIDs         []int64 `json:"genre_ids"`
	ID               int64   `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     *string `json:"backdrop_path"`
	Popularity       float64 `json:"popularity"`
	VoteCount        int64   `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
}

// SearchMovie is the results returned from a Movie search.
type SearchMovie struct {
	Search
	Results []MovieResult `json:"results"`
}

// SearchMovie searches the TMDB Movie database for the name given.
func (t *TMDB) SearchMovie(query string, params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)

	if err := t.get(s, "/3/search/movie", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}

	return s, nil
}

// TVOrMovie contains information which could be for either a TV show or a movie.
type TVOrMovie struct {
	PosterPath       *string  `json:"poster_path"`
	Adult            bool     `json:"adult"`
	Overview         string   `json:"overview"`
	ReleaseDate      string   `json:"release_date"`
	OriginalTitle    string   `json:"original_title"`
	GenreIDs         []int64  `json:"genre_ids"`
	ID               int64    `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	Title            string   `json:"title"`
	BackdropPath     *string  `json:"backdrop_path"`
	Popularity       float64  `json:"popularity"`
	VoteCount        int64    `json:"vote_count"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
}

// PeopleResult stores a single result of a people search.
type PeopleResult struct {
	ProfilePath *string `json:"profile_path"`
	Adult       bool    `json:"adult"`
	ID          int64   `json:"id"`
	KnownFor    struct {
		TVOrMovie
		MediaType string `json:"media_type"`
	} `json:"known_for"`
	Name       string  `json:"name"`
	Popularity float64 `json:"popularity"`
}

// SearchPeople searc the TMDB People database for the name given.
type SearchPeople struct {
	Search
	Results []PeopleResult `json:"results"`
}

// SearchPerson searches the TMDB people database for the name given.
func (t *TMDB) SearchPerson(query string, params ...option) (*SearchPeople, error) {
	s := new(SearchPeople)

	if err := t.get(s, "/3/search/person", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}

	return s, nil
}

// TVResult stores a single result of a TV search.
type TVResult struct {
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

// SearchTV is the results returned from a TV search.
type SearchTV struct {
	Search
	Results []TVResult `json:"results"`
}

// SearchTV searches the TMDB TV show database for the name given.
func (t *TMDB) SearchTV(query string, params ...option) (*SearchTV, error) {
	s := new(SearchTV)

	if err := t.get(s, "/3/search/tv", url.Values{"query": []string{query}}, params...); err != nil {
		return nil, err
	}

	return s, nil
}
