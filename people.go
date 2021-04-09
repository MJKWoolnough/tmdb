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

// PersonMovieCredits retreives the movie credits for a person
func (t *TMDB) PersonMovieCredits(id int64, params ...option) (*Credits, error) {
	c := new(Credits)
	if err := t.get(c, fmt.Sprintf("/3/person/%d/movie_credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// PersonTVCredits retreives the movie credits for a person
func (t *TMDB) PersonTVCredits(id int64, params ...option) (*Credits, error) {
	c := new(Credits)
	if err := t.get(c, fmt.Sprintf("/3/person/%d/tv_credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// PersonCombinedCredits retreives the movie credits for a person
func (t *TMDB) PersonCombinedCredits(id int64, params ...option) (*Credits, error) {
	c := new(Credits)
	if err := t.get(c, fmt.Sprintf("/3/person/%d/combined_credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// PersonExternalIDs retreives the external IDs for a person
func (t *TMDB) PersonExternalIDs(id int64, params ...option) (*ExternalIDs, error) {
	e := new(ExternalIDs)
	if err := t.get(e, fmt.Sprintf("/3/person/%d/external_ids", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return e, nil
}

// PersonImages contains all of the images for a person
type PersonImages struct {
	ID       int64   `json:"id"`
	Profiles []Image `json:"profiles"`
}

// PersonImages retreives all of the images for a person
func (t *TMDB) PersonImages(id int64) (*PersonImages, error) {
	p := new(PersonImages)
	if err := t.get(p, fmt.Sprintf("/3/person/%d/images", id), url.Values{}); err != nil {
		return nil, err
	}
	return p, nil
}
