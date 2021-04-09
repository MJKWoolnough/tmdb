package tmdb

import (
	"encoding/json"
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

// PersonChanges contains informaton about changes made to a person's profile
type PersonChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string          `json:"id"`
			Action        string          `json:"action"`
			Time          string          `json:"time"`
			OriginalValue json.RawMessage `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// PersonChanges retreives the information about changes made to a person's profile
func (t *TMDB) PersonChanges(id int64, params ...option) (*PersonChanges, error) {
	p := new(PersonChanges)
	if err := t.get(p, fmt.Sprintf("/3/person/%d/changes", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return p, nil
}
