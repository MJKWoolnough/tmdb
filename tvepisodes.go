package tmdb

import (
	"fmt"
	"net/url"
)

// TVEpisode retrieves the details for an episode of a TV show
func (t *TMDB) TVEpisode(id int64, season int64, episode int64, params ...option) (*TVEpisode, error) {
	tv := new(TVEpisode)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d", id, season, episode), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}
