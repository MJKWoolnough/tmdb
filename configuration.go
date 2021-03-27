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
func (t *TMDB) Configuration() (*Configuration, error) {
	c := new(Configuration)
	if err := t.get(c, "/3/configuration", url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}

// Countries represents the list of ISO3166-1 country tags
type Countries []struct {
	Country     string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
}

// Countries retrieves the list of ISO3166-1 country tags
func (t *TMDB) Countries() (*Countries, error) {
	c := new(Countries)
	if err := t.get(c, "/3/countries", url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}
