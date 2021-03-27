package tmdb

import "net/url"

// Configuration contains system-wide configuration information
type Configuration struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	}
	ChangeKeys []string `json:"change_keys"`
}

// Configuration retrieves the system-wide configuration information
func (t *TMDB) Configuration() (*configuration, error) {
	c := new(Configuration)
	if err := t.get(c, "/3/confiruation", url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}
