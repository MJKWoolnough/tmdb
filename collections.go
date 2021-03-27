package tmdb

import (
	"fmt"
	"net/url"
)

// CollectionDetails stores a collection
type CollectionDetails struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	BackdropPath string `json:"backdrop_path"`
	Parts        []struct {
		Adult            bool    `json:"adult"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		PosterPath       string  `json:"poster_path"`
		Popularity       float64 `json:"popularity"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int64   `json:"vote_count"`
	} `json:"parts"`
}

// CollectionDetails retrieves collection details by id
func (t *TMDB) CollectionDetails(id int64, params ...option) (*CollectionDetails, error) {
	c := new(CollectionDetails)
	if err := t.get(c, fmt.Sprintf("/3/collection/%d", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// CollectionImage stores the data of a single image
type CollectionImage struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      uint64  `json:"height"`
	Language    *string `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	Width       int64   `json:"width"`
}

// CollectionImages stores the images details for a collection
type CollectionImages struct {
	ID        int64           `json:"id"`
	Backdrops CollectionImage `json:"backdrops"`
	Posters   CollectionImage `json:"posters"`
}

// CollectionImages retrieves collection images details by id
func (t *TMDB) CollectionImages(id int64, params ...option) (*CollectionImages, error) {
	c := new(CollectionImages)
	if err := t.get(c, fmt.Sprintf("/3/collection/%d/images", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}
