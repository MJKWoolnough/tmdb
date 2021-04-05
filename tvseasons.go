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
