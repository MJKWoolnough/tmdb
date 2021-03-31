package tmdb

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// MovieDetails contains all of the details of a particular movie
type MovieDetails struct {
	Adult               bool            `json:"adult"`
	BackdropPath        *string         `json:"backdrop_path"`
	BelongsToCollection json.RawMessage `json:"belongs_to_collection"`
	Budget              int64           `json:"budget"`
	Genres              []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            *string `json:"homepage"`
	ID                  int64   `json:"id"`
	IMDBID              *string `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            *string `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          *string `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string  `json:"name"`
		ID            int64   `json:"id"`
		LogoPath      *string `json:"logo_path"`
		OriginCountry string  `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Country string `json:"iso_3166_1"`
		Name    string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int64  `json:"revenue"`
	Runtime         *int64 `json:"runtime"`
	SpokenLanguages []struct {
		Language string `json:"iso_639_1"`
		Name     string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     *string `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
}

// MovieDetails retrieves movie information based on the ID specified
func (t *TMDB) MovieDetails(id int64, params ...option) (*MovieDetails, error) {
	m := new(MovieDetails)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieAlternativeTitles is a list of the alternative titles for a movie
type MovieAlternativeTitles struct {
	ID     int64 `json:"id"`
	Titles []struct {
		Country string `json:"iso_3166_1"`
		Title   string `json:"title"`
		Type    string `json:"type"`
	} `json:"titles"`
}

// MovieAlternativeTitles returns all of the alternative titles for a movie
func (t *TMDB) MovieAlternativeTitles(id int64, params ...option) (*MovieAlternativeTitles, error) {
	m := new(MovieAlternativeTitles)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/alternative_titles", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieChanges lists changes to a movie entry
type MovieChanges []struct {
	Key   string `json:"key"`
	Items []struct {
		ID            string `json:"id"`
		Action        string `json:"action"`
		Time          string `json:"time"`
		Language      string `json:"iso_639_1"`
		Value         string `json:"value"`
		OriginalValue string `json:"original_value"`
	}
}

// MovieChanges returns changes to a movie entry
func (t *TMDB) MovieChanges(id int64, params ...option) (*MovieChanges, error) {
	m := new(MovieChanges)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/changes", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// CreditShared represents the shared information for crediting either a cast or crew member
type CreditShared struct {
	Adult              bool    `json:"adult"`
	Gender             *int64  `json:"gender"`
	ID                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	CreditID           int64   `json:"credit_id"`
}

// CastCredit represents the credit information for a cast member
type CastCredit struct {
	CreditShared
	Character string `json:"character"`
	CastID    string `json:"cast_id"`
	Order     int64  `json:"order"`
}

// CrewCredit represents the credit information for a crew memeber
type CrewCredit struct {
	CreditShared
	Department string `json:"department"`
	Job        string `json:"job"`
}

// MovieCredits contains all of the credits for a movie
type MovieCredits struct {
	ID   int64
	Cast []CastCredit `json:"cast"`
	Crew []CrewCredit `json:"crew"`
}

// MovieCredits retrieves all of the credits for a movie
func (t *TMDB) MovieCredits(id int64, params ...option) (*MovieCredits, error) {
	m := new(MovieCredits)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieExternalIDs contains all known external IDs for a movie
type MovieExternalIDs struct {
	IMDB      *string `json:"imdb_id"`
	Facebook  *string `json:"facebook_id"`
	Instagram *string `json:"instagram_id"`
	Twitter   *string `json:"twitter_id"`
	ID        int64   `json:"id"`
}

// MovieExternalIDs retrieves all known external IDs for a movie
func (t *TMDB) MovieExternalIDs(id int64) (*MovieExternalIDs, error) {
	m := new(MovieExternalIDs)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/external_ids", id), url.Values{}); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieImages contains all linked images for a movie
type MovieImages struct {
	ID        int64   `json:"id"`
	Backdrops []Image `json:"backdrops"`
	Posters   []Image `json:"posters"`
}

// MovieImages retrieves all linked images for a movie
func (t *TMDB) MovieImages(id int64, params ...option) (*MovieImages, error) {
	m := new(MovieImages)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/images", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieKeywords contains all of a movies keywords
type MovieKeywords struct {
	ID       int64     `json:"id"`
	Keywords []Keyword `json:"keywords"`
}

// MovieKeywords retrieves all of the keywords for a movie
func (t *TMDB) MovieKeywords(id int64) (*MovieKeywords, error) {
	m := new(MovieKeywords)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/keywords", id), url.Values{}); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieLists contains a list of lists that belong to a specific movie
type MovieLists struct {
	ID int64 `json:"id"`
	Search
	Results []struct {
		Description   string  `json:"description"`
		FavoriteCount int64   `json:"favorite_count"`
		ID            int64   `json:"id"`
		ItemCount     int64   `json:"item_count"`
		Language      string  `json:"iso_639_1"`
		ListType      string  `json:"list_type"`
		Name          string  `json:"name"`
		PosterPath    *string `json:"poster_path"`
	} `json:"results"`
}

// MovieLists retrieves a list of lists that belong to the specified movie
func (t *TMDB) MovieLists(id int64, params ...option) (*MovieLists, error) {
	m := new(MovieLists)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/lists", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}
