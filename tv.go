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
