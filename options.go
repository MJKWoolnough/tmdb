package tmdb

import (
	"net/url"
	"strconv"
)

type option interface {
	setParam(url.Values)
}

// Language is a ISO 639-1 value to display translated data for the fields that support it
type Language string

func (l Language) setParam(v url.Values) {
	v.Set("language", string(l))
}

// Page specifies which page of results to query
type Page uint16

func (p Page) setParam(v url.Values) {
	v.Set("page", strconv.FormatUint(uint64(p), 10))
}

// IncludeAdult determines whether adult content is included in search results
type IncludeAdult bool

func (i IncludeAdult) setParam(v url.Values) {
	v.Set("include_adult", strconv.FormatBool(bool(i)))
}

// FirstAirDateYear limits a search to the year specified
type FirstAirDateYear int64

func (f FirstAirDateYear) setParam(v url.Values) {
	v.Set("first_air_date_year", strconv.FormatInt(int64(f), 10))
}

// Region specified a ISO 3166-1 code to filter release dates. Must be uppercase
type Region string

func (r Region) setParam(v url.Values) {
	v.Set("region", string(r))
}

// Year filters the results by the specified year
type Year int64

func (y Year) setParam(v url.Values) {
	v.Set("year", strconv.FormatInt(int64(y), 10))
}

// PrimaryReleaseYear filters the results by the specified year
type PrimaryReleaseYear int64

func (p PrimaryReleaseYear) setParam(v url.Values) {
	v.Set("primary_release_year", strconv.FormatInt(int64(p), 10))
}

// Country filters a list by country
type Country string

func (c Country) setParam(v url.Values) {
	v.Set("country", string(c))
}

// StartDate limits a search to a starting date
type StartDate string

func (s StartDate) setParam(v url.Values) {
	v.Set("start_date", string(s))
}

// EndDate limits a search to a end date
type EndDate string

func (e EndDate) setParam(v url.Values) {
	v.Set("end_date", string(e))
}

// SortBy sorts an applicable list bythe specified key and direction
type SortBy string

func (s SortBy) setParam(v url.Values) {
	v.Set("sort_by", string(s))
}

// AirDateGTE filters and only include TV shows that have an airy date greater than (or equal to) that specified
type AirDateGTE string

func (a AirDateGTE) setParam(v url.Values) {
	v.Set("air_date.gte", string(a))
}

// AirDateLTE filters and only include TV shows that have any air date less than (or equal to) that specified
type AirDateLTE string

func (a AirDateLTE) setParam(v url.Values) {
	v.Set("air_date.gte", string(a))
}

// FirstAirDateGTE filters and only include TV shows that have a first air date greater than (or equal to) that specified
type FirstAirDateGTE string

func (f FirstAirDateGTE) setParam(v url.Values) {
	v.Set("air_date.lte", string(f))
}

// FirstAirDateLTE filters and only include TV shows that have a first air date less than (or equal to) that specified
type FirstAirDateLTE string

func (f FirstAirDateLTE) setParam(v url.Values) {
	v.Set("first_air_date.lte", string(f))
}

// Timezone is used to set the correct timezone for date limiting filters
type Timezone string

func (t Timezone) setParam(v url.Values) {
	v.Set("timezone", string(t))
}

// VoteAverageGTE filters by only allowing results with a vote average greater than (or equal to) that specified
type VoteAverageGTE float64

func (va VoteAverageGTE) setParam(v url.Values) {
	v.Set("vote_average.gte", strconv.FormatFloat(float64(va), 'f', 0, 64))
}

// VoteAverageLTE filters by only allowing results with a vote average less than (or equal to) that specified
type VoteAverageLTE float64

func (va VoteAverageLTE) setParam(v url.Values) {
	v.Set("vote_average.lte", strconv.FormatFloat(float64(va), 'f', 0, 64))
}

// WithGenres specifies a comma seperated list of genre ids to limit the results to
type WithGenres string

func (w WithGenres) setParam(v url.Values) {
	v.Set("with_genres", string(w))
}

// WithNetworks specifies a comma seperated list of network IDs to limit the results to
type WithNetworks string

func (w WithNetworks) setParam(v url.Values) {
	v.Set("with_networks", string(w))
}

// WithoutGenres specifies a comma seperated list of genre ids to filter from the results
type WithoutGenres string

func (w WithoutGenres) setParam(v url.Values) {
	v.Set("without_genres", string(w))
}

// WithRuntimeGTE filters by only allowing results with a runtime greater than (or equal to) that specified
type WithRuntimeGTE int64

func (w WithRuntimeGTE) setParam(v url.Values) {
	v.Set("with_runtime.gte", strconv.FormatFloat(float64(w), 'f', 0, 64))
}

// WithRuntimeLTE filters by only allowing results with a runtime less than (or equal to) that specified
type WithRuntimeLTE int64

func (w WithRuntimeLTE) setParam(v url.Values) {
	v.Set("with_runtime.lte", strconv.FormatFloat(float64(w), 'f', 0, 64))
}

// IncludeNullFirstAirDates specifies that results without a first air date should be included when filtering by air date
type IncludeNullFirstAirDates bool

func (i IncludeNullFirstAirDates) setParam(v url.Values) {
	v.Set("include_null_first_air_dates", strconv.FormatBool(bool(i)))
}

// WithOriginalLanguage specifies and ISO 639-1 string to filter results by original language
type WithOriginalLanguage string

func (w WithOriginalLanguage) setParam(v url.Values) {
	v.Set("with_original_language", string(w))
}

// WithoutKeywords filters results by the specified keywords
type WithoutKeywords string

func (w WithoutKeywords) setParam(v url.Values) {
	v.Set("without_keywords", string(w))
}

// ScreenedTheatrically filters results to only include those with a theatrical release
type ScreenedTheatrically bool

func (s ScreenedTheatrically) setParam(v url.Values) {
	v.Set("screened_theatrically", strconv.FormatBool(bool(s)))
}

// WithCompanies specifies a comma seperated list of company ids to limit the results to
type WithCompanies string

func (w WithCompanies) setParam(v url.Values) {
	v.Set("with_companies", string(w))
}

// WithKeywords filters results by the specified keywords
type WithKeywords string

func (w WithKeywords) setParam(v url.Values) {
	v.Set("with_keywords", string(w))
}

// WithWatchProviders filters results by the specified watch proviers
type WithWatchProviders string

func (w WithWatchProviders) setParam(v url.Values) {
	v.Set("with_watch_providers", string(w))
}

// WatchRegion filters results by the specified regions
type WatchRegion string

func (w WatchRegion) setParam(v url.Values) {
	v.Set("watch_region", string(w))
}
