package tmdb

import "net/url"

// Review contains a review for a movie or TV show
type Review struct {
	ID            string `json:"id"`
	Author        string `json:"author"`
	AuthorDetails struct {
		Name       string `json:"name"`
		Username   string `json:"username"`
		AvatarPath string `json:"avatar_path"`
		Rating     int64  `json:"rating"`
	} `json:"author_details"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	Language   string `json:"iso_639_1"`
	MediaID    int64  `json:"media_id"`
	MediaTitle string `json:"media_title"`
	UpdatedAt  string `json:"updated_at"`
	URL        string `json:"url"`
}

// Review retrieves information about a specific review
func (t *TMDB) Review(id string) (*Review, error) {
	r := new(Review)
	if err := t.get(r, "/3/review/"+id, url.Values{}); err != nil {
		return nil, err
	}
	return r, nil
}
