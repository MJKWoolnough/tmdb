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

// TVEpisodeChanges retrieves changes for an episode of a TV show
func (t *TMDB) TVEpisodeChanges(id int64, season int64, episode int64, params ...option) (*EntryChanges, error) {
	e := new(EntryChanges)
	if err := t.get(e, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d/changes", id, season, episode), url.Values{}, params...); err != nil {
		return nil, err
	}
	return e, nil
}

// TVEpisodeCredits contains the credit information for an episode of a TV show
type TVEpisodeCredits struct {
	ID         int64
	Cast       []CastCredit `json:"cast"`
	Crew       []CrewCredit `json:"crew"`
	GuestStars []CastCredit `json:"guest_stars"`
}

// TVEpisodeCredits retrieves the credit information for an episode of a TV show
func (t *TMDB) TVEpisodeCredits(id int64, season int64, episode int64, params ...option) (*TVEpisodeCredits, error) {
	tv := new(TVEpisodeCredits)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d/credits", id, season, episode), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVEpisodeExternalIDs contains all of the known external IDs for an episode of a TV show
type TVEpisodeExternalIDs struct {
	IMDB        *string `json:"imdb_id"`
	FreebaseMID *string `json:"freebase_mid"`
	FreebaseID  *string `json:"freebase_id"`
	TVDB        *int64  `json:"tvdb_id"`
	TVRage      *int64  `json:"tbrage_id"`
	ID          int64   `json:"id"`
}

// TVEpisodeExternalIDs retrieves all of the known external IDs for an episode of a TV show
func (t *TMDB) TVEpisodeExternalIDs(id int64, season int64, episode int64, params ...option) (*TVEpisodeExternalIDs, error) {
	tv := new(TVEpisodeExternalIDs)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d/external_ids", id, season, episode), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVEpisodeImages retrieves all of the images for an episode of a TV show
func (t *TMDB) TVEpisodeImages(id int64, season int64, episode int64, params ...option) (*Images, error) {
	i := new(Images)
	if err := t.get(i, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d/images", id, season, episode), url.Values{}, params...); err != nil {
		return nil, err
	}
	return i, nil
}

// TVEpisodeTranslations retrieves all of the translations for an episode of a TV show
func (t *TMDB) TVEpisodeTranslations(id int64, season int64, episode int64, params ...option) (*Translations, error) {
	tr := new(Translations)
	if err := t.get(tr, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d/translations", id, season, episode), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tr, nil
}
