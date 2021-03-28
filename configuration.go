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
	if err := t.get(c, "/3/configuration/countries", url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}

// Jobs represents the list of jobs/departments
type Jobs []struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// Jobs retrieves the list of jobs/departments
func (t *TMDB) Jobs() (*Jobs, error) {
	j := new(Jobs)
	if err := t.get(j, "/3/configuration/jobs", url.Values{}); err != nil {
		return nil, err
	}
	return j, nil
}

// Languages represents the list of languages
type Languages []struct {
	Language    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

// Languages retrieves the list of languages
func (t *TMDB) Languages() (*Languages, error) {
	l := new(Languages)
	if err := t.get(l, "/3/configurationslanguages", url.Values{}); err != nil {
		return nil, err
	}
	return l, nil
}

// PrimaryTranslations represents the list of primary translations
type PrimaryTranslations []string

// PrimaryTranslations retrieves the list of primary translations
func (t *TMDB) PrimaryTranslations() (*PrimaryTranslations, error) {
	p := new(PrimaryTranslations)
	if err := t.get(p, "/3/configuration/primary_translations", url.Values{}); err != nil {
		return nil, err
	}
	return p, nil
}
