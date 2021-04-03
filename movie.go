package tmdb

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ProductionCompany contains information about the production company of a movie/show
type ProductionCompany struct {
	Name          string  `json:"name"`
	ID            int64   `json:"id"`
	LogoPath      *string `json:"logo_path"`
	OriginCountry string  `json:"origin_country"`
}

// MovieDetails contains all of the details of a particular movie
type MovieDetails struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path"`
	BelongsToCollection json.RawMessage     `json:"belongs_to_collection"`
	Budget              int64               `json:"budget"`
	Genres              Genres              `json:"genres"`
	Homepage            *string             `json:"homepage"`
	ID                  int64               `json:"id"`
	IMDBID              *string             `json:"imdb_id"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalTitle       string              `json:"original_title"`
	Overview            *string             `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          *string             `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
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

// AlternativeTitles is a list of the alternative titles for a movie
type AlternativeTitles struct {
	ID     int64 `json:"id"`
	Titles []struct {
		Country string `json:"iso_3166_1"`
		Title   string `json:"title"`
		Type    string `json:"type"`
	} `json:"titles"`
}

// MovieAlternativeTitles returns all of the alternative titles for a movie
func (t *TMDB) MovieAlternativeTitles(id int64, params ...option) (*AlternativeTitles, error) {
	a := new(AlternativeTitles)
	if err := t.get(a, fmt.Sprintf("/3/movie/%d/alternative_titles", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return a, nil
}

// EntryChanges lists changes to a movie entry
type EntryChanges []struct {
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
func (t *TMDB) MovieChanges(id int64, params ...option) (*EntryChanges, error) {
	e := new(EntryChanges)
	if err := t.get(e, fmt.Sprintf("/3/movie/%d/changes", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return e, nil
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
}

// CastCredit represents the credit information for a cast member
type CastCredit struct {
	CreditShared
	CreditID  int64  `json:"credit_id"`
	Character string `json:"character"`
	CastID    string `json:"cast_id"`
	Order     int64  `json:"order"`
}

// CrewCredit represents the credit information for a crew memeber
type CrewCredit struct {
	CreditShared
	CreditID   int64  `json:"credit_id"`
	Department string `json:"department"`
	Job        string `json:"job"`
}

// Credits contains all of the credits for a movie
type Credits struct {
	ID   int64
	Cast []CastCredit `json:"cast"`
	Crew []CrewCredit `json:"crew"`
}

// MovieCredits retrieves all of the credits for a movie
func (t *TMDB) MovieCredits(id int64, params ...option) (*Credits, error) {
	c := new(Credits)
	if err := t.get(c, fmt.Sprintf("/3/movie/%d/credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
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

// Images contains all linked images for a movie
type Images struct {
	ID        int64   `json:"id"`
	Backdrops []Image `json:"backdrops"`
	Posters   []Image `json:"posters"`
}

// MovieImages retrieves all linked images for an entry
func (t *TMDB) MovieImages(id int64, params ...option) (*Images, error) {
	i := new(Images)
	if err := t.get(i, fmt.Sprintf("/3/movie/%d/images", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return i, nil
}

// Keywords contains all of an entry's keywords
type Keywords struct {
	ID       int64     `json:"id"`
	Keywords []Keyword `json:"keywords"`
}

// MovieKeywords retrieves all of the keywords for a movie
func (t *TMDB) MovieKeywords(id int64) (*Keywords, error) {
	k := new(Keywords)
	if err := t.get(k, fmt.Sprintf("/3/movie/%d/keywords", id), url.Values{}); err != nil {
		return nil, err
	}
	return k, nil
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

// MovieRecommendations retrieves a list of recommended movies for a movie
func (t *TMDB) MovieRecommendations(id int64, params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, fmt.Sprintf("/3/movie/%d/recommendations", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// MovieReleaseDates contains a list of release dates, by country, for a movie
type MovieReleaseDates struct {
	ID      int64 `json:"id"`
	Results []struct {
		Country      string `json:"iso_3166_1"`
		ReleaseDates []struct {
			Certification string `json:"certification"`
			Language      string `json:"iso_639_1"`
			Type          int64  `json:"type"`
			Note          string `json:"note"`
		} `json:"release_dates"`
	} `json:"results"`
}

// MovieReleaseDates retrieves a list of release dates, by country, for the given movie
func (t *TMDB) MovieReleaseDates(id int64) (*MovieReleaseDates, error) {
	m := new(MovieReleaseDates)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/release_dates", id), url.Values{}); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieReviews contains user reviews for a movie
type MovieReviews struct {
	ID int64 `json:"id"`
	Search
	Results []struct {
		Author        string `json:"author"`
		AuthorDetails struct {
			Name       string  `json:"name"`
			Username   string  `json:"username"`
			AvatarPath *string `json:"avatar_path"`
			Rating     *int64  `json:"rating"`
		} `json:"author_details"`
		Content   string `json:"content"`
		CreatedAt string `json:"created_at"`
		ID        string `json:"id"`
		UpdatedAt string `json:"updated_at"`
		URL       string `json:"url"`
	} `json:"results"`
}

// MovieReviews returns a list of reviews for the specified movie
func (t *TMDB) MovieReviews(id int64, params ...option) (*MovieReviews, error) {
	m := new(MovieReviews)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/reviews", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieSimilar retrieves a list of recommended movies for a movie
func (t *TMDB) MovieSimilar(id int64, params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, fmt.Sprintf("/3/movie/%d/similar", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// MovieTranslations retrieves all of the translations for a movie
func (t *TMDB) MovieTranslations(id int64) (*Translations, error) {
	tr := new(Translations)
	if err := t.get(t, fmt.Sprintf("/3/movie/%d/translations", id), url.Values{}); err != nil {
		return nil, err
	}
	return tr, nil
}

// MovieVideos contains all of the videos related to a movie
type MovieVideos struct {
	ID      int64 `json:"id"`
	Results []struct {
		ID       string `json:"id"`
		Language string `json:"iso_639_1"`
		Country  string `json:"iso_3166_1"`
		Key      string `json:"key"`
		Name     string `json:"name"`
		Site     string `json:"site"`
		Size     int64  `json:"size"`
		Type     string `json:"type"`
	} `json:"results"`
}

// MovieVideos retrieves all of the videos related to a movie
func (t *TMDB) MovieVideos(id int64, params ...option) (*MovieVideos, error) {
	m := new(MovieVideos)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/videos", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchProviderData contains information about the provider of a video service
type WatchProviderData struct {
	DisplayPriority int64  `json:"display_priority"`
	LogoPath        string `json:"logo_path"`
	ProviderID      int64  `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
}

// WatchProviders contains all of the information about where and how to watch a movie
type WatchProviders struct {
	ID      int64 `json:"id"`
	Results struct {
		AR, AT, AU, BE, BR, CA, CH, CL, CO, CZ, DE, DK, EC, EE, ES, FI, FR, GB, GR, HU, ID, IE, IN, IT, JP, KR, LT, LV, MX, MY, NL, NO, NZ, PE, PH, PL, PT, RO, RU, SE, SG, TH, TR, US, VE, ZA struct {
			Link     string            `json:"link"`
			Flatrate WatchProviderData `json:"flatrate"`
			Rent     WatchProviderData `json:"rent"`
			But      WatchProviderData `json:"buy"`
		}
	} `json:"results"`
}

// MovieWatchProviders retrieves all of the ways to watch a movie
func (t *TMDB) MovieWatchProviders(id int64) (*WatchProviders, error) {
	m := new(WatchProviders)
	if err := t.get(m, fmt.Sprintf("/3/movie/%d/watch/providers", id), url.Values{}); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieLatest returns the latest movie added to the database
func (t *TMDB) MovieLatest(params ...option) (*MovieDetails, error) {
	m := new(MovieDetails)
	if err := t.get(m, "/3/movie/latest", url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MoviesWithDates contains a list of movies and the minimum and maximum play dates
type MoviesWithDates struct {
	Search
	Results []MovieResult `json:"results"`
	Dates   struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	}
}

// MovieNowPlaying retrives a list of currently playing movies
func (t *TMDB) MovieNowPlaying(params ...option) (*MoviesWithDates, error) {
	m := new(MoviesWithDates)
	if err := t.get(m, "/3/movie/now_playing", url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}

// MoviePopular retrieves a list of popular movies
func (t *TMDB) MoviePopular(params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/movie/popular", url.Values{}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// MovieTopRated retrieves a list of top-rated movies
func (t *TMDB) MovieTopRated(params ...option) (*SearchMovie, error) {
	s := new(SearchMovie)
	if err := t.get(s, "/3/movie/top_rated", url.Values{}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// MovieUpcoming retrieves a list of upcoming movies
func (t *TMDB) MovieUpcoming(params ...option) (*MoviesWithDates, error) {
	m := new(MoviesWithDates)
	if err := t.get(m, "/3/movie/upcoming", url.Values{}, params...); err != nil {
		return nil, err
	}
	return m, nil
}
