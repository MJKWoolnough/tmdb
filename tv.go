package tmdb

import (
	"fmt"
	"net/url"
)

// TVShow contains all of the information about a TV Show
type TVShow struct {
	BackdropPath *string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int64   `json:"id"`
		CreditID    string  `json:"credit_id"`
		Name        string  `json:"name"`
		Gender      int64   `json:"gender"`
		ProfilePath *string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime   []int64  `json:"episode_run_time"`
	FirstAirDate     string   `json:"first_air_date"`
	Genres           Genres   `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int64    `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int64   `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int64   `json:"season_number"`
		StillPath      *string `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
	} `json:"last_episode_to_air"`
	Name                string              `json:"name"`
	Networks            []ProductionCompany `json:"networks"`
	NumberOfEpisodes    int64               `json:"number_of_episodes"`
	NumberOfSeasons     int64               `json:"number_of_seasons"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalName        string              `json:"original_name"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          *string             `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []struct {
		Country string `json:"iso_3166_1"`
		Name    string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int64  `json:"episode_count"`
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int64  `json:"season_number"`
	} `json:"seasons"`
	SpokenLanguages Languages `json:"spoken_languages"`
	Status          string    `json:"status"`
	Tagline         string    `json:"tagline"`
	Type            string    `json:"type"`
	VoteAverage     float64   `json:"vote_average"`
	VoteCount       int64     `json:"vote_count"`
}

// TVShow retrieves the details of a TV Show
func (t *TMDB) TVShow(id int64, params ...option) (*TVShow, error) {
	tv := new(TVShow)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVAggregateCredits contains the aggregated credits for a TV show
type TVAggregateCredits struct {
	Cast []struct {
		CreditShared
		Roles []struct {
			CreditID     string `json:"credit_id"`
			Character    string `json:"character"`
			EpisodeCount int64  `json:"episode_count"`
		} `json:"roles"`
		TotalEpisodeCount int64 `json:"total_episode_count"`
		Order             int64 `json:"order"`
	} `json:"cast"`
	Crew []struct {
		CreditShared
		Jobs []struct {
			CreditID     string `json:"credit_id"`
			Job          string `json:"job"`
			EpisodeCount int64  `json:"episode_count"`
		} `jobs:"jobs"`
	} `json:"crew"`
	ID int64 `json:"id"`
}

// TVAggregateCredits retrieves the aggregated credits for a TV show
func (t *TMDB) TVAggregateCredits(id int64, params ...option) (*TVAggregateCredits, error) {
	tv := new(TVAggregateCredits)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/aggregate_credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVAlternativeTitles retrieves the alternative titles for a TV show
func (t *TMDB) TVAlternativeTitles(id int64, params ...option) (*AlternativeTitles, error) {
	a := new(AlternativeTitles)
	if err := t.get(a, fmt.Sprintf("/3/tv/%d/alternative_titles", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return a, nil
}

// TVChanges retrieves the list of changes for a TV show
func (t *TMDB) TVChanges(id int64, params ...option) (*EntryChanges, error) {
	e := new(EntryChanges)
	if err := t.get(e, fmt.Sprintf("/3/tv/%d/changes", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return e, nil
}

// TVContentRatings containts the content ratings for a TV show
type TVContentRatings struct {
	Results []struct {
		Country string `json:"iso_3166_1"`
		Rating  string `json:"rating"`
	} `json:"results"`
	ID int64 `json:"id"`
}

// TVContentRatings retrieves the content ratings for a TV show
func (t *TMDB) TVContentRatings(id int64, params ...option) (*TVContentRatings, error) {
	tv := new(TVContentRatings)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/content_ratings", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVCredits retrieves the credits for a TV show
func (t *TMDB) TVCredits(id int64, params ...option) (*Credits, error) {
	c := new(Credits)
	if err := t.get(c, fmt.Sprintf("/3/tv/%d/credits", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return c, nil
}

// TVEpisodeGroups contains all of the episode groups for a TV show
type TVEpisodeGroups struct {
	Results []struct {
		Description  string             `json:"description"`
		EpisodeCount int64              `json:"episode_count"`
		ID           string             `json:"id"`
		Name         string             `json:"name"`
		Network      *ProductionCompany `json:"network"`
		Type         int64              `json:"type"`
	} `json:"results"`
	ID int64 `json:"id"`
}

// TVEpisodeGroups retrieves all of the episode groups for a TV show
func (t *TMDB) TVEpisodeGroups(id int64, params ...option) (*TVEpisodeGroups, error) {
	tv := new(TVEpisodeGroups)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/episode_groups", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVExternalIDs contains all known external IDs for a TV show
type TVExternalIDs struct {
	IMDB        *string `json:"imdb_id"`
	FreebaseMID *string `json:"freebase_mid"`
	FreebaseID  *string `json:"freebase_id"`
	TVDB        *int64  `json:"tvdb_id"`
	TVRage      *int64  `json:"tbrage_id"`
	Facebook    *string `json:"facebook_id"`
	Instagram   *string `json:"instagram_id"`
	Twitter     *string `json:"twitter_ id"`
	ID          int64   `json:"id"`
}

// TVExternalIDs retrieves all of the external ids for a TV show
func (t *TMDB) TVExternalIDs(id int64, params ...option) (*TVExternalIDs, error) {
	tv := new(TVExternalIDs)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/external_ids", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVImages retrieves all of the images for a TV show
func (t *TMDB) TVImages(id int64, params ...option) (*Images, error) {
	i := new(Images)
	if err := t.get(i, fmt.Sprintf("/3/tv/%d/images", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return i, nil
}

// TVKeywords retrieves all of the keywords for a TV show
func (t *TMDB) TVKeywords(id int64) (*Keywords, error) {
	k := new(Keywords)
	if err := t.get(k, fmt.Sprintf("/3/tv/%d/keywords", id), url.Values{}); err != nil {
		return nil, err
	}
	return k, nil
}

// TVRecommendations retrieves all of the recommendations for a TV show
func (t *TMDB) TVRecommendations(id int64, params ...option) (*SearchTV, error) {
	s := new(SearchTV)
	if err := t.get(s, fmt.Sprintf("/3/tv/%d/recommendations", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// TVReviews retrieves reviews for a TV show
func (t *TMDB) TVReviews(id int64, params ...option) (*Reviews, error) {
	r := new(Reviews)
	if err := t.get(r, fmt.Sprintf("/3/tv/%d/reviews", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return r, nil
}

// TVScreenedTheatrically contains all of the episodes of a TV show that were screened theatrically
type TVScreenedTheatrically struct {
	ID      int64 `json:"id"`
	Results []struct {
		ID            int64 `json:"id"`
		EpisodeNumber int64 `json:"episode_number"`
		SeasonNumber  int64 `json:"season_number"`
	} `json:"results"`
}

// TVScreenedTheatrically retrieves all of the episodes of a TV show that were screened theatrically
func (t *TMDB) TVScreenedTheatrically(id int64) (*TVScreenedTheatrically, error) {
	tv := new(TVScreenedTheatrically)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/screen_theatrically", id), url.Values{}); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVSimilar retrieves all of the similar TV shows
func (t *TMDB) TVSimilar(id int64, params ...option) (*SearchTV, error) {
	s := new(SearchTV)
	if err := t.get(s, fmt.Sprintf("/3/tv/%d/similar", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return s, nil
}

// TVTranslations contains all of the translations that exist for a show
func (t *TMDB) TVTranslations(id int64) (*Translations, error) {
	tv := new(Translations)
	if err := t.get(tv, fmt.Sprintf("/3/tv/%d/translations", id), url.Values{}); err != nil {
		return nil, err
	}
	return tv, nil
}

// TVVideos retrieves all of the videos for a TV shows
func (t *TMDB) TVVideos(id int64, params ...option) (*Videos, error) {
	v := new(Videos)
	if err := t.get(v, fmt.Sprintf("/3/tv/%d/videos", id), url.Values{}, params...); err != nil {
		return nil, err
	}
	return v, nil
}
