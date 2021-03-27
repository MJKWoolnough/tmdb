package tmdb

import (
	"fmt"
	"net/url"
)

// CompanyDetails represents all of the details of a company
type CompanyDetails struct {
	Description   string          `json:"description"`
	Headquarters  string          `json:"headquarters"`
	Homepage      string          `json:"homepage"`
	ID            int64           `json:"id"`
	LogoPath      string          `json:"logo_path"`
	Name          string          `json:"name"`
	OriginCountry string          `json:"origin_country"`
	ParentCompany *CompanyDetails `json:"parent_company"`
}

// CompanyDetails retrieves company's details by id
func (t *TMDB) CompanyDetails(id int64) (*CompanyDetails, error) {
	c := new(CompanyDetails)
	if err := t.get(c, fmt.Sprintf("/3/company/%d", id), url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}
