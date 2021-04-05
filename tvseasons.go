package tmdb

import (
	"fmt"
	"net/url"
)

// TVSeason contains the details for a season of a TV show
type TVSeason struct {
	StrID    string `json:"_id"`
	AirDate  string `json:"air_date"`
	Episodes []struct {
		Episode
		Crew       []CrewCredit
		GuestStars []CastCredit
	} `json:"episodes"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	ID           int64   `json:"id"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int64   `json:"season_number"`
}

// TVSeason retrieves the details for a season of a TV show
func (t *TMDB) TVSeason(id int64, season int64, params ...option) (*TVSeason, error) {
	tv := new(TVSeason)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/season/%d", id, season), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVSeasonAggregateCredits retrieves the aggregate credits for a season of a TV show
func (t *TMDB) TVSeasonAggregateCredits(id int64, season int64, params ...option) (*TVAggregateCredits, error) {
	tv := new(TVAggregateCredits)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/season/%d/aggregate_credits", id, season), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVSeasonChanges retrieves changes for a season of a TV show
func (t *TMDB) TVSeasonChanges(id int64, season int64, params ...option) (*EntryChanges, error) {
	e := new(EntryChanges)
	if err := t.get(e, fmt.Sprintf("/3/tv/%d/season/%d/changes", id, season), url.Values{}, params...); err != nil {
		return nil, err
	}
	return e, nil
}

// TVSeasonCredits retrieves credits for a season of a TV show
func (t *TMDB) TVSeasonCredits(id int64, season int64, params ...option) (*Credits, error) {
	c := new(Credits)
	if err := t.get(c, fmt.Sprintf("/3/tv/%d/season/%d/credits", id, season), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// TVSeasonExternalIDs retrieves all of the external ids for a season of a TV show
func (t *TMDB) TVSeasonExternalIDs(id int64, season int64, params ...option) (*TVExternalIDs, error) {
	tv := new(TVExternalIDs)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/season/%d/external_ids", id, season), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}
