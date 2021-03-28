package tmdb

import (
	"fmt"
	"net/url"
)

// Credit represents a movie or TV credit details
type Credit struct {
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	Media      struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		OriginalName string `json:"original_name"`
		Character    string `json:"character"`
		Episodes     []struct {
			AirDate       string `json:"air_date"`
			EpisodeNumber int64  `json:"episode_number"`
			Name          string `json:"name"`
			Overview      string `json:"overview"`
			SeasonNumber  int64  `json:"season_number"`
			StillPath     string `json:"still_path"`
		} `json:"episodes"`
		Seasons []struct {
			AirDate      string `json:"ait_date"`
			PosterPath   string `json:"poster_path"`
			SeasonNumber int64  `json:"season_number"`
		} `json:"seasons"`
	} `json:"media"`
	MediaType string `json:"media_type"`
	ID        string `json:"id"`
	Person    struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"person"`
}

// Credit retrieves a movie or TV credit
func (t *TMDB) Credit(id int64) (*Credit, error) {
	c := new(Credit)
	if err := t.get(c, fmt.Sprintf("/3/credit/%d", id), url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}
