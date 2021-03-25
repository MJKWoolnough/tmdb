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