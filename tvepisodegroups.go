package tmdb

import (
	"fmt"
	"net/url"
)

// TVEpisodeGroup contains the details for an episode groups for a TV show
type TVEpisodeGroup struct {
	Description  string `json:"description"`
	EpisodeCount int64  `json:"episode_count"`
	GroupCount   int64  `json:"group_count"`
	Groups       []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Order    int64  `json:"order"`
		Episodes []struct {
			Episode
			Order int64 `json:"int64"`
		} `json:"episodes"`
		Locked bool `json:"locked"`
	} `json:"groups"`
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Network ProductionCompany `json:"network"`
	Type    int64             `json:"type"`
}

// TVEpisodeGroup retrieves the details for an episode group for a TV show
func (t *TMDB) TVEpisodeGroup(id int64, params ...option) (*TVEpisodeGroup, error) {
	tv := new(TVEpisodeGroup)
	if err := t.get(tv, fmt.Sprintf("/3/tv/episode_group/%d", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}
