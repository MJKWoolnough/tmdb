package tmdb

import (
	"fmt"
	"net/url"
)

// NetworkDetails gets the details of a network.
func (t *TMDB) NetworkDetails(id int64) (*CompanyDetails, error) {
	c := new(CompanyDetails)

	if err := t.get(c, fmt.Sprintf("/3/network/%d", id), url.Values{}); err != nil {
		return nil, err
	}

	return c, nil
}

// NetworkAlternativeNames gets the alternative names of a network.
func (t *TMDB) NetworkAlternativeNames(id int64) (*CompanyAlternativeNames, error) {
	c := new(CompanyAlternativeNames)

	if err := t.get(c, fmt.Sprintf("/3/network/%d/alternative_names", id), url.Values{}); err != nil {
		return nil, err
	}

	return c, nil
}

// NetworkImages retrieves the logos for a network.
func (t *TMDB) NetworkImages(id int64) (*CompanyImages, error) {
	c := new(CompanyImages)

	if err := t.get(c, fmt.Sprintf("/3/network/%d/images", id), url.Values{}); err != nil {
		return nil, err
	}

	return c, nil
}
