package tmdb

import (
	"fmt"
	"net/url"
)

// Person contains the information about a person
type Person struct {
	Birthday           *string  `json:"birthday"`
	KnownForDepartment string   `json:"known_for_department"`
	Deathday           *string  `json:"deathday"`
	ID                 int64    `json:"id"`
	AKA                []string `json:"also_known_as"`
	Gender             int64    `json:"gender"`
	Biography          string   `json:"biography"`
	Popularity         float64  `json:"popularity"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	ProfilePath        *string  `json:"profile_path"`
	Adult              bool     `json:"adult"`
	IMDBID             string   `json:"imdb_id"`
	Homepage           string   `json:"homepage"`
}

// PersonDetails retreives the details of the specified person
func (t *TMDB) PersonDetails(id int64, params ...option) (*Person, error) {
	p := new(Person)
	if err := t.get(p, fmt.Sprintf("/3/person/%d", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return p, nil
}
