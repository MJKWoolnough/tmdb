package tmdb

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// MovieDetails contains all of the details of a particular movie
type MovieDetails struct {
	Adult               bool            `json:"adult"`
	BackdropPath        *string         `json:"backdrop_path"`
	BelongsToCollection json.RawMessage `json:"belongs_to_collection"`
	Budget              int64           `json:"budget"`
	Genres              []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            *string `json:"homepage"`
	ID                  int64   `json:"id"`
	IMDBID              *string `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            *string `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          *string `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string  `json:"name"`
		ID            int64   `json:"id"`
		LogoPath      *string `json:"logo_path"`
		OriginCountry string  `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Country string `json:"iso_3166_1"`
		Name    string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int64  `json:"revenue"`
	Runtime         *int64 `json:"runtime"`
	SpokenLanguages []struct {
		Language string `json:"iso_649_1"`
		Name     string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     *string `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
}

// MovieDetails retrieves movie information based on the ID specified
func (t *TMDB) MovieDetails(id int64, params ...option) (*MovieDetails, error) {
	m := new(MovieDetails)
	if err := t.get(fmt.Sprintf("/3/movie/%d", id), url.Values{}, params, m); err != nil {
		return nil, err
	}
	return m, nil
}
