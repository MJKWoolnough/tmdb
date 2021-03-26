package tmdb

import "net/url"

// Certification represents a single certification description
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int64  `json:"order"`
}

// Certifications represents the certifications founds in numerous countries
type Certifications struct {
	US, CA, AU, FR, RU, DE, TH, KR, GB, BR []Certification
}

// CertificationMovie retrieves all movie certifications
func (t *TMDB) CertificationMovie() (*Certification, error) {
	c := new(Certification)
	if err := t.get("/3/certification/movie/list", url.Values{}, []option{}, c); err != nil {
		return nil, err
	}
	return c, nil
}

// CertificationTV retrieves all tv certifications
func (t *TMDB) CertificationTV() (*Certification, error) {
	c := new(Certification)
	if err := t.get("/3/certification/tv/list", url.Values{}, []option{}, c); err != nil {
		return nil, err
	}
	return c, nil
}
