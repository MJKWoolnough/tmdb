package tmdb

// TMDB holds an APIv4 key for themoviedb.org
type TMDB struct {
	header string
}

// New takes a TMDB APIv4 key
func New(v4key string) *TMDB {
	return &TMDB{
		header: "Bearer " + v4key,
	}
}
