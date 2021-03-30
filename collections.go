package tmdb

import (
	"fmt"
	"net/url"
)

// CollectionDetails stores a collection
type CollectionDetails struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	Overview     string        `json:"overview"`
	BackdropPath string        `json:"backdrop_path"`
	Parts        []MovieResult `json:"parts"`
}

// CollectionDetails retrieves collection details by id
func (t *TMDB) CollectionDetails(id int64, params ...option) (*CollectionDetails, error) {
	c := new(CollectionDetails)
	if err := t.get(c, fmt.Sprintf("/3/collection/%d", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// Image stores the data of a single image
type Image struct {
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
	ID        int64 `json:"id"`
	Backdrops Image `json:"backdrops"`
	Posters   Image `json:"posters"`
}

// CollectionImages retrieves collection images details by id
func (t *TMDB) CollectionImages(id int64, params ...option) (*CollectionImages, error) {
	c := new(CollectionImages)
	if err := t.get(c, fmt.Sprintf("/3/collection/%d/images", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// CollectionTranslations represents collection translations data
type CollectionTranslations struct {
	ID           int64 `json:"id"`
	Translations []struct {
		Country     string `json:"iso_3166_1"`
		Language    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title    string `json:"title"`
			Overview string `json:"overview"`
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}

// CollectionTranslations gets the list translations for a collection by ID
func (t *TMDB) CollectionTranslations(id int64, params ...option) (*CollectionTranslations, error) {
	c := new(CollectionTranslations)
	if err := t.get(c, fmt.Sprintf("/3/collection/%d/translations", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}
