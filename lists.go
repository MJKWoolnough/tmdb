package tmdb

import (
	"net/url"
	"strconv"
)

// ListDetails is the details of a list
type ListDetails struct {
	CreatedBy     string        `json:"created_by"`
	Description   string        `json:"description"`
	FavoriteCount int64         `json:"favorite_count"`
	ID            string        `json:"id"`
	Items         []MovieResult `json:"items"`
	ItemCount     int64         `json:"item_count"`
	Language      string        `json:"iso_639_1"`
	Name          string        `json:"name"`
	PosterPath    *string       `json:"poster_path"`
}

// ListDetails retrieves the details of a list
func (t *TMDB) ListDetails(id string, params ...option) (*ListDetails, error) {
	l := new(ListDetails)
	if err := t.get(l, "/3/list/"+id, url.Values{}, params...); err != nil {
		return nil, err
	}
	return l, nil
}

// ListItemPresent states whether a move exists within a list
type ListItemPresent struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// ListItemPresent checks whether a list contains a movie
func (t *TMDB) ListItemPresent(listID string, movieID int64) (*ListItemPresent, error) {
	l := new(ListItemPresent)
	if err := t.get(l, "/3/list/"+listID+"/item_status", url.Values{"movie_id": []string{strconv.FormatInt(movieID, 10)}}, params...); err != nil {
		return nil, err
	}
	return l, nil
}
