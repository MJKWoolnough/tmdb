package tmdb

import (
	"fmt"
	"net/url"
)

// NetworkDetails gets the details of a network
func (t *TMDB) NetworkDetails(id int64) (*CompanyDetails, error) {
	c := new(CompanyDetails)
	if err := t.get(c, fmt.Sprintf("/3/network/%d", id), url.Values{}); err != nil {
		return nil, err
	}
	return c, nil
}
