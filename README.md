# tmdb
--
    import "vimagination.zapto.org/tmdb"

Package tmdb is a simple interface to the themoviedb.org database

## Usage

#### type APIKey

```go
type APIKey interface {
	// contains filtered or unexported methods
}
```

APIKey is used to set the API key for a request

#### func  V3Key

```go
func V3Key(v3key string) APIKey
```
V3Key create new TVDM APIv3 key

#### func  V4Key

```go
func V4Key(v4key string) APIKey
```
V4Key creates a new TMDB APIv4 key

#### type AirDateGTE

```go
type AirDateGTE string
```

AirDateGTE filters and only include TV shows that have an airy date greater than
(or equal to) that specified

#### type AirDateLTE

```go
type AirDateLTE string
```

AirDateLTE filters and only include TV shows that have any air date less than
(or equal to) that specified

#### type AlternativeTitles

```go
type AlternativeTitles struct {
	ID     int64 `json:"id"`
	Titles []struct {
		Country string `json:"iso_3166_1"`
		Title   string `json:"title"`
		Type    string `json:"type"`
	} `json:"titles"`
}
```

AlternativeTitles is a list of the alternative titles for a movie

#### type CastCredit

```go
type CastCredit struct {
	CreditShared
	CreditID  int64  `json:"credit_id"`
	Character string `json:"character"`
	CastID    string `json:"cast_id"`
	Order     int64  `json:"order"`
}
```

CastCredit represents the credit information for a cast member

#### type Certification

```go
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int64  `json:"order"`
}
```

Certification represents a single certification description

#### type CertificationCountry

```go
type CertificationCountry string
```

CertificationCountry is used in conjustion fir the Certification filter to
specify a country with a valid certificate

#### type CertificationFilter

```go
type CertificationFilter string
```

CertificationFilter filters results with a vald certification from the
CertificationCountry value

#### type CertificationGTE

```go
type CertificationGTE string
```

CertificationGTE filters results to only include those with a certification
greater than (or equal to) that specified

#### type CertificationLTE

```go
type CertificationLTE string
```

CertificationLTE filters results to only include those with a certification less
than (or equal to) that specified

#### type Certifications

```go
type Certifications struct {
	US, CA, AU, FR, RU, DE, TH, KR, GB, BR []Certification
}
```

Certifications represents the certifications founds in numerous countries

#### type Change

```go
type Change struct {
	ID    int64 `json:"id"`
	Adult *bool `json:"adult"`
}
```

Change represents the ID of an item change

#### type Changes

```go
type Changes struct {
	Results      []Change `json:"results"`
	Page         int64    `json:"page"`
	TotalPages   int64    `json:"total_pages"`
	TotalResults int64    `json:"total_results"`
}
```

Changes lists results from a change search

#### type CollectionDetails

```go
type CollectionDetails struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	Overview     string        `json:"overview"`
	BackdropPath string        `json:"backdrop_path"`
	Parts        []MovieResult `json:"parts"`
}
```

CollectionDetails stores a collection

#### type CollectionImages

```go
type CollectionImages struct {
	ID        int64 `json:"id"`
	Backdrops Image `json:"backdrops"`
	Posters   Image `json:"posters"`
}
```

CollectionImages stores the images details for a collection

#### type CollectionResult

```go
type CollectionResult struct {
	ID           int64   `json:"id"`
	BackdropPath *string `json:"backdrop_path"`
	Name         string  `json:"name"`
	PosterPath   *string `json:"poster_path"`
}
```

CollectionResult stores a single result of a collection search

#### type CompanyAlternativeNames

```go
type CompanyAlternativeNames struct {
	ID      int64 `json:"id"`
	Results []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}
```

CompanyAlternativeNames represents a list of alternative name for a company

#### type CompanyDetails

```go
type CompanyDetails struct {
	Description   string          `json:"description"`
	Headquarters  string          `json:"headquarters"`
	Homepage      string          `json:"homepage"`
	ID            int64           `json:"id"`
	LogoPath      string          `json:"logo_path"`
	Name          string          `json:"name"`
	OriginCountry string          `json:"origin_country"`
	ParentCompany *CompanyDetails `json:"parent_company"`
}
```

CompanyDetails represents all of the details of a company

#### type CompanyImages

```go
type CompanyImages struct {
	ID    int64   `json:"id"`
	Logos []Image `json:"logos"`
}
```

CompanyImages represents a list of logos for a company

#### type CompanyResult

```go
type CompanyResult struct {
	ID       int64   `json:"id"`
	LogoPath *string `json:"logo_path"`
	Name     string  `json:"name"`
}
```

CompanyResult stores a single result of a company search

#### type Configuration

```go
type Configuration struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	}
	ChangeKeys []string `json:"change_keys"`
}
```

Configuration contains system-wide configuration information

#### type Countries

```go
type Countries []struct {
	Country     string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
}
```

Countries represents the list of ISO3166-1 country tags

#### type Country

```go
type Country string
```

Country filters a list by country

#### type Credit

```go
type Credit struct {
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	Media      struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		OriginalName string `json:"original_name"`
		Character    string `json:"character"`
		Episodes     []struct {
			AirDate       string `json:"air_date"`
			EpisodeNumber int64  `json:"episode_number"`
			Name          string `json:"name"`
			Overview      string `json:"overview"`
			SeasonNumber  int64  `json:"season_number"`
			StillPath     string `json:"still_path"`
		} `json:"episodes"`
		Seasons []struct {
			AirDate      string `json:"ait_date"`
			PosterPath   string `json:"poster_path"`
			SeasonNumber int64  `json:"season_number"`
		} `json:"seasons"`
	} `json:"media"`
	MediaType string `json:"media_type"`
	ID        string `json:"id"`
	Person    struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"person"`
}
```

Credit represents a movie or TV credit details

#### type CreditShared

```go
type CreditShared struct {
	Adult              *bool   `json:"adult"`
	Gender             *int64  `json:"gender"`
	ID                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}
```

CreditShared represents the shared information for crediting either a cast or
crew member

#### type Credits

```go
type Credits struct {
	ID   int64
	Cast []CastCredit `json:"cast"`
	Crew []CrewCredit `json:"crew"`
}
```

Credits contains all of the credits for a movie

#### type CrewCredit

```go
type CrewCredit struct {
	CreditShared
	CreditID   int64  `json:"credit_id"`
	Department string `json:"department"`
	Job        string `json:"job"`
}
```

CrewCredit represents the credit information for a crew memeber

#### type EndDate

```go
type EndDate string
```

EndDate limits a search to a end date

#### type EntryChanges

```go
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
```

EntryChanges lists changes to a movie entry

#### type Episode

```go
type Episode struct {
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
}
```

Episode contains basic inforamtion for an episode

#### type Error

```go
type Error struct {
	Message string `json:"status_message"`
	Code    int64  `json:"status_code"`
}
```

Error represents an error repsonse (401, 404) from the API

#### func (*Error) Error

```go
func (e *Error) Error() string
```

#### type ExternalIDs

```go
type ExternalIDs struct {
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
```

ExternalIDs contains all known external IDs for a TV show

#### type FindResults

```go
type FindResults struct {
	MovieResults  []MovieResult  `json:"movie_results"`
	PersonResults []PeopleResult `json:"person_results"`
	TVResults     []TVResult     `json:"tv_result"`
}
```

FindResults contains the results from the Find call

#### type FirstAirDateGTE

```go
type FirstAirDateGTE string
```

FirstAirDateGTE filters and only include TV shows that have a first air date
greater than (or equal to) that specified

#### type FirstAirDateLTE

```go
type FirstAirDateLTE string
```

FirstAirDateLTE filters and only include TV shows that have a first air date
less than (or equal to) that specified

#### type FirstAirDateYear

```go
type FirstAirDateYear int64
```

FirstAirDateYear limits a search to the year specified

#### type Genres

```go
type Genres []struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
```

Genres is a list of genres Names and IDs

#### type Image

```go
type Image struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      uint64  `json:"height"`
	Language    *string `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	Width       int64   `json:"width"`
}
```

Image stores the data of a single image

#### type Images

```go
type Images struct {
	ID        int64   `json:"id"`
	Backdrops []Image `json:"backdrops"`
	Posters   []Image `json:"posters"`
}
```

Images contains all linked images for a movie

#### type IncludeAdult

```go
type IncludeAdult bool
```

IncludeAdult determines whether adult content is included in search results

#### type IncludeImageLanguage

```go
type IncludeImageLanguage string
```

IncludeImageLanguage determines whether an the language of an image is returned
with the image data

#### type IncludeNullFirstAirDates

```go
type IncludeNullFirstAirDates bool
```

IncludeNullFirstAirDates specifies that results without a first air date should
be included when filtering by air date

#### type IncludeVideo

```go
type IncludeVideo bool
```

IncludeVideo is a filter to include/exclude videos

#### type Jobs

```go
type Jobs []struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}
```

Jobs represents the list of jobs/departments

#### type Keyword

```go
type Keyword struct {
	ID   int64  `json:"id"`
	Name string `json:"string"`
}
```

Keyword contains the ID and Name of a keyword

#### type Keywords

```go
type Keywords struct {
	ID       int64     `json:"id"`
	Keywords []Keyword `json:"keywords"`
}
```

Keywords contains all of an entry's keywords

#### type Language

```go
type Language string
```

Language is a ISO 639-1 value to display translated data for the fields that
support it

#### type Languages

```go
type Languages []struct {
	Language    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}
```

Languages represents the list of languages

#### type ListDetails

```go
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
```

ListDetails is the details of a list

#### type ListItemPresent

```go
type ListItemPresent struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}
```

ListItemPresent states whether a move exists within a list

#### type MovieDetails

```go
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
```

MovieDetails contains all of the details of a particular movie

#### type MovieExternalIDs

```go
type MovieExternalIDs struct {
	IMDB      *string `json:"imdb_id"`
	Facebook  *string `json:"facebook_id"`
	Instagram *string `json:"instagram_id"`
	Twitter   *string `json:"twitter_id"`
	ID        int64   `json:"id"`
}
```

MovieExternalIDs contains all known external IDs for a movie

#### type MovieLists

```go
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
```

MovieLists contains a list of lists that belong to a specific movie

#### type MovieReleaseDates

```go
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
```

MovieReleaseDates contains a list of release dates, by country, for a movie

#### type MovieResult

```go
type MovieResult struct {
	PosterPath       *string `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIDs         []int64 `json:"genre_ids"`
	ID               int64   `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     *string `json:"backdrop_path"`
	Popularity       float64 `json:"popularity"`
	VoteCount        int64   `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
}
```

MovieResult stores a single result of a movie search

#### type MoviesWithDates

```go
type MoviesWithDates struct {
	Search
	Results []MovieResult `json:"results"`
	Dates   struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	}
}
```

MoviesWithDates contains a list of movies and the minimum and maximum play dates

#### type Page

```go
type Page uint16
```

Page specifies which page of results to query

#### type PeopleResult

```go
type PeopleResult struct {
	ProfilePath *string `json:"profile_path"`
	Adult       bool    `json:"adult"`
	ID          int64   `json:"id"`
	KnownFor    struct {
		TVOrMovie
		MediaType string `json:"media_type"`
	} `json:"known_for"`
	Name       string  `json:"name"`
	Popularity float64 `json:"popularity"`
}
```

PeopleResult stores a single result of a people search

#### type Person

```go
type Person struct {
	Birthday           *string  `json:"birthday"`
	KnownForDepartment string   `json:"known_for_department"`
	Deathday           *string  `json:"deathday"`
	ID                 int64    `json:"id"`
	Name               string   `json:"name"`
	AKA                []string `json:"also_known_as"`
	Gender             int64    `json:"gender"`
	Biography          string   `json:"biography"`
	Popularity         float64  `json:"popularity"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	ProfilePath        *string  `json:"profile_path"`
	Adult              bool     `json:"adult"`
	IMDBID             string   `json:"imdb_id"`
	Homepage           string   `json:"homepage"`
}
```

Person contains the information about a person

#### type PersonChanges

```go
type PersonChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string          `json:"id"`
			Action        string          `json:"action"`
			Time          string          `json:"time"`
			OriginalValue json.RawMessage `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}
```

PersonChanges contains informaton about changes made to a person's profile

#### type PersonImages

```go
type PersonImages struct {
	ID       int64   `json:"id"`
	Profiles []Image `json:"profiles"`
}
```

PersonImages contains all of the images for a person

#### type PersonTaggedImages

```go
type PersonTaggedImages struct {
	ID int64 `json:"id"`
	Search
	Results []struct {
		Image
		ID        string    `json:"id"`
		ImageType string    `json:"image_type"`
		Media     TVOrMovie `json:"media"`
		MediaType string    `json:"media_type"`
	} `json:"results"`
}
```

PersonTaggedImages contains a list of all tagged images for a person

#### type PersonTranslations

```go
type PersonTranslations struct {
	Translations []struct {
		Language string `json:"iso_639_1"`
		Country  string `json:"iso_3166_1"`
		Name     string `json:"name"`
		Data     struct {
			Biography string `json:"biography"`
		} `json:"data"`
		EnglishName string `json:"english_name"`
	} `json:"translations"`
	ID int64 `json:"id"`
}
```

PersonTranslations contains the translations that have been created for a person

#### type PrimaryReleaseDateGTE

```go
type PrimaryReleaseDateGTE string
```

PrimaryReleaseDateGTE filters results to only include those with a primary
release date greater than (or equal to) that specified

#### type PrimaryReleaseDateLTE

```go
type PrimaryReleaseDateLTE string
```

PrimaryReleaseDateLTE filters results to only include those with a primary
release date less than (or equal to) that specified

#### type PrimaryReleaseYear

```go
type PrimaryReleaseYear int64
```

PrimaryReleaseYear filters the results by the specified year

#### type PrimaryTranslations

```go
type PrimaryTranslations []string
```

PrimaryTranslations represents the list of primary translations

#### type ProductionCompany

```go
type ProductionCompany struct {
	Name          string  `json:"name"`
	ID            int64   `json:"id"`
	LogoPath      *string `json:"logo_path"`
	OriginCountry string  `json:"origin_country"`
}
```

ProductionCompany contains information about the production company of a
movie/show

#### type Region

```go
type Region string
```

Region specified a ISO 3166-1 code to filter release dates. Must be uppercase

#### type ReleaseDateGTE

```go
type ReleaseDateGTE string
```

ReleaseDateGTE filters results to only include those with a release date greater
than (or equal to) that specified

#### type ReleaseDateLTE

```go
type ReleaseDateLTE string
```

ReleaseDateLTE filters results to only include those with a release date less
than (or equal to) that specified

#### type Review

```go
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
```

Review contains a review for a movie or TV show

#### type Reviews

```go
type Reviews struct {
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
```

Reviews contains user reviews for an entry

#### type ScreenedTheatrically

```go
type ScreenedTheatrically bool
```

ScreenedTheatrically filters results to only include those with a theatrical
release

#### type Search

```go
type Search struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
}
```

Search stores the pages fields for a returned search

#### type SearchCollection

```go
type SearchCollection struct {
	Search
	Results []CollectionResult `json:"results"`
}
```

SearchCollection is the results returned from a Collection search

#### type SearchCompany

```go
type SearchCompany struct {
	Search
	Results []CompanyResult `json:"results"`
}
```

SearchCompany is the results returned from a Company search

#### type SearchKeywords

```go
type SearchKeywords struct {
	Search
	Results []Keyword `json:"results"`
}
```

SearchKeywords is the results returned from a Keyword search

#### type SearchMovie

```go
type SearchMovie struct {
	Search
	Results []MovieResult `json:"results"`
}
```

SearchMovie is the results returned from a Movie search

#### type SearchPeople

```go
type SearchPeople struct {
	Search
	Results []PeopleResult `json:"results"`
}
```

SearchPeople searc the TMDB People database for the name given

#### type SearchTV

```go
type SearchTV struct {
	Search
	Results []TVResult `json:"results"`
}
```

SearchTV is the results returned from a TV search

#### type SortBy

```go
type SortBy string
```

SortBy sorts an applicable list bythe specified key and direction

#### type StartDate

```go
type StartDate string
```

StartDate limits a search to a starting date

#### type TMDB

```go
type TMDB struct {
}
```

TMDB holds an API Key to allow connection to TMDB

#### func  New

```go
func New(apiKey APIKey) *TMDB
```
New takes a TMDB API key to create a TMDB connection

#### func (*TMDB) CertificationMovie

```go
func (t *TMDB) CertificationMovie() (*Certification, error)
```
CertificationMovie retrieves all movie certifications

#### func (*TMDB) CertificationTV

```go
func (t *TMDB) CertificationTV() (*Certification, error)
```
CertificationTV retrieves all tv certifications

#### func (*TMDB) ChangesMovie

```go
func (t *TMDB) ChangesMovie(params ...option) (*Changes, error)
```
ChangesMovie returns all of the Movie changes within the last 24 hours, filtered
by the given params

#### func (*TMDB) ChangesPerson

```go
func (t *TMDB) ChangesPerson(params ...option) (*Changes, error)
```
ChangesPerson returns all of the Person changes within the last 24 hours,
filtered by the given params

#### func (*TMDB) ChangesTV

```go
func (t *TMDB) ChangesTV(params ...option) (*Changes, error)
```
ChangesTV returns all of the TV changes within the last 24 hours, filtered by
the given params

#### func (*TMDB) CollectionDetails

```go
func (t *TMDB) CollectionDetails(id int64, params ...option) (*CollectionDetails, error)
```
CollectionDetails retrieves collection details by id

#### func (*TMDB) CollectionImages

```go
func (t *TMDB) CollectionImages(id int64, params ...option) (*CollectionImages, error)
```
CollectionImages retrieves collection images details by id

#### func (*TMDB) CollectionTranslations

```go
func (t *TMDB) CollectionTranslations(id int64, params ...option) (*Translations, error)
```
CollectionTranslations gets the list translations for a collection by ID

#### func (*TMDB) CompanyAlternativeNames

```go
func (t *TMDB) CompanyAlternativeNames(id int64) (*CompanyAlternativeNames, error)
```
CompanyAlternativeNames retrieves the alternative name for a company

#### func (*TMDB) CompanyDetails

```go
func (t *TMDB) CompanyDetails(id int64) (*CompanyDetails, error)
```
CompanyDetails retrieves company's details by id

#### func (*TMDB) CompanyImages

```go
func (t *TMDB) CompanyImages(id int64) (*CompanyImages, error)
```
CompanyImages retrieves the logos for a company

#### func (*TMDB) Configuration

```go
func (t *TMDB) Configuration() (*Configuration, error)
```
Configuration retrieves the system-wide configuration information

#### func (*TMDB) Countries

```go
func (t *TMDB) Countries() (*Countries, error)
```
Countries retrieves the list of ISO3166-1 country tags

#### func (*TMDB) Credit

```go
func (t *TMDB) Credit(id int64) (*Credit, error)
```
Credit retrieves a movie or TV credit

#### func (*TMDB) DiscoverMovie

```go
func (t *TMDB) DiscoverMovie(params ...option) (*SearchMovie, error)
```
DiscoverMovie search the movie db filterd by the specified options

#### func (*TMDB) DiscoverTV

```go
func (t *TMDB) DiscoverTV(params ...option) (*SearchTV, error)
```
DiscoverTV search the tv db filterd by the specified options

#### func (*TMDB) Find

```go
func (t *TMDB) Find(id, source string, params ...option) (*FindResults, error)
```
Find searches the database for movies, people, or TV shows that reference the
given external ID

#### func (*TMDB) Jobs

```go
func (t *TMDB) Jobs() (*Jobs, error)
```
Jobs retrieves the list of jobs/departments

#### func (*TMDB) Keyword

```go
func (t *TMDB) Keyword(id int64) (*Keyword, error)
```
Keyword returns the Keyword specified by the ID given

#### func (*TMDB) KeywordMovies

```go
func (t *TMDB) KeywordMovies(id int64, params ...option) (*SearchMovie, error)
```
KeywordMovies retrieves a list of movies that share the keyword given

#### func (*TMDB) Languages

```go
func (t *TMDB) Languages() (*Languages, error)
```
Languages retrieves the list of languages

#### func (*TMDB) ListDetails

```go
func (t *TMDB) ListDetails(id string, params ...option) (*ListDetails, error)
```
ListDetails retrieves the details of a list

#### func (*TMDB) ListItemPresent

```go
func (t *TMDB) ListItemPresent(listID string, movieID int64) (*ListItemPresent, error)
```
ListItemPresent checks whether a list contains a movie

#### func (*TMDB) MovieAlternativeTitles

```go
func (t *TMDB) MovieAlternativeTitles(id int64, params ...option) (*AlternativeTitles, error)
```
MovieAlternativeTitles returns all of the alternative titles for a movie

#### func (*TMDB) MovieChanges

```go
func (t *TMDB) MovieChanges(id int64, params ...option) (*EntryChanges, error)
```
MovieChanges returns changes to a movie entry

#### func (*TMDB) MovieCredits

```go
func (t *TMDB) MovieCredits(id int64, params ...option) (*Credits, error)
```
MovieCredits retrieves all of the credits for a movie

#### func (*TMDB) MovieDetails

```go
func (t *TMDB) MovieDetails(id int64, params ...option) (*MovieDetails, error)
```
MovieDetails retrieves movie information based on the ID specified

#### func (*TMDB) MovieExternalIDs

```go
func (t *TMDB) MovieExternalIDs(id int64) (*MovieExternalIDs, error)
```
MovieExternalIDs retrieves all known external IDs for a movie

#### func (*TMDB) MovieGenres

```go
func (t *TMDB) MovieGenres(params ...option) (Genres, error)
```
MovieGenres retrieves the official genres for movies

#### func (*TMDB) MovieImages

```go
func (t *TMDB) MovieImages(id int64, params ...option) (*Images, error)
```
MovieImages retrieves all linked images for an entry

#### func (*TMDB) MovieKeywords

```go
func (t *TMDB) MovieKeywords(id int64) (*Keywords, error)
```
MovieKeywords retrieves all of the keywords for a movie

#### func (*TMDB) MovieLatest

```go
func (t *TMDB) MovieLatest(params ...option) (*MovieDetails, error)
```
MovieLatest returns the latest movie added to the database

#### func (*TMDB) MovieLists

```go
func (t *TMDB) MovieLists(id int64, params ...option) (*MovieLists, error)
```
MovieLists retrieves a list of lists that belong to the specified movie

#### func (*TMDB) MovieNowPlaying

```go
func (t *TMDB) MovieNowPlaying(params ...option) (*MoviesWithDates, error)
```
MovieNowPlaying retrives a list of currently playing movies

#### func (*TMDB) MoviePopular

```go
func (t *TMDB) MoviePopular(params ...option) (*SearchMovie, error)
```
MoviePopular retrieves a list of popular movies

#### func (*TMDB) MovieRecommendations

```go
func (t *TMDB) MovieRecommendations(id int64, params ...option) (*SearchMovie, error)
```
MovieRecommendations retrieves a list of recommended movies for a movie

#### func (*TMDB) MovieReleaseDates

```go
func (t *TMDB) MovieReleaseDates(id int64) (*MovieReleaseDates, error)
```
MovieReleaseDates retrieves a list of release dates, by country, for the given
movie

#### func (*TMDB) MovieReviews

```go
func (t *TMDB) MovieReviews(id int64, params ...option) (*Reviews, error)
```
MovieReviews returns a list of reviews for the specified movie

#### func (*TMDB) MovieSimilar

```go
func (t *TMDB) MovieSimilar(id int64, params ...option) (*SearchMovie, error)
```
MovieSimilar retrieves a list of recommended movies for a movie

#### func (*TMDB) MovieTopRated

```go
func (t *TMDB) MovieTopRated(params ...option) (*SearchMovie, error)
```
MovieTopRated retrieves a list of top-rated movies

#### func (*TMDB) MovieTranslations

```go
func (t *TMDB) MovieTranslations(id int64) (*Translations, error)
```
MovieTranslations retrieves all of the translations for a movie

#### func (*TMDB) MovieUpcoming

```go
func (t *TMDB) MovieUpcoming(params ...option) (*MoviesWithDates, error)
```
MovieUpcoming retrieves a list of upcoming movies

#### func (*TMDB) MovieVideos

```go
func (t *TMDB) MovieVideos(id int64, params ...option) (*Videos, error)
```
MovieVideos retrieves all of the videos related to a movie

#### func (*TMDB) MovieWatchProviders

```go
func (t *TMDB) MovieWatchProviders(id int64) (*WatchProviders, error)
```
MovieWatchProviders retrieves all of the ways to watch a movie

#### func (*TMDB) NetworkAlternativeNames

```go
func (t *TMDB) NetworkAlternativeNames(id int64) (*CompanyAlternativeNames, error)
```
NetworkAlternativeNames gets the alternative names of a network

#### func (*TMDB) NetworkDetails

```go
func (t *TMDB) NetworkDetails(id int64) (*CompanyDetails, error)
```
NetworkDetails gets the details of a network

#### func (*TMDB) NetworkImages

```go
func (t *TMDB) NetworkImages(id int64) (*CompanyImages, error)
```
NetworkImages retrieves the logos for a network

#### func (*TMDB) PersonChanges

```go
func (t *TMDB) PersonChanges(id int64, params ...option) (*PersonChanges, error)
```
PersonChanges retreives the information about changes made to a person's profile

#### func (*TMDB) PersonCombinedCredits

```go
func (t *TMDB) PersonCombinedCredits(id int64, params ...option) (*Credits, error)
```
PersonCombinedCredits retreives the movie credits for a person

#### func (*TMDB) PersonDetails

```go
func (t *TMDB) PersonDetails(id int64, params ...option) (*Person, error)
```
PersonDetails retreives the details of the specified person

#### func (*TMDB) PersonExternalIDs

```go
func (t *TMDB) PersonExternalIDs(id int64, params ...option) (*ExternalIDs, error)
```
PersonExternalIDs retreives the external IDs for a person

#### func (*TMDB) PersonImages

```go
func (t *TMDB) PersonImages(id int64) (*PersonImages, error)
```
PersonImages retreives all of the images for a person

#### func (*TMDB) PersonLatest

```go
func (t *TMDB) PersonLatest(params ...option) (*Person, error)
```
PersonLatest retreives the most recent person added to the database

#### func (*TMDB) PersonMovieCredits

```go
func (t *TMDB) PersonMovieCredits(id int64, params ...option) (*Credits, error)
```
PersonMovieCredits retreives the movie credits for a person

#### func (*TMDB) PersonPopular

```go
func (t *TMDB) PersonPopular(params ...option) (*SearchPeople, error)
```
PersonPopular retreives the most popular people in the database

#### func (*TMDB) PersonTVCredits

```go
func (t *TMDB) PersonTVCredits(id int64, params ...option) (*Credits, error)
```
PersonTVCredits retreives the movie credits for a person

#### func (*TMDB) PersonTaggedImages

```go
func (t *TMDB) PersonTaggedImages(id int64, params ...option) (*PersonTaggedImages, error)
```
PersonTaggedImages retreives all of the tagged images for the specified person

#### func (*TMDB) PersonTranslations

```go
func (t *TMDB) PersonTranslations(id int64, params ...option) (*PersonTranslations, error)
```
PersonTranslations retreives all of the translations that have been created for
a person

#### func (*TMDB) PrimaryTranslations

```go
func (t *TMDB) PrimaryTranslations() (*PrimaryTranslations, error)
```
PrimaryTranslations retrieves the list of primary translations

#### func (*TMDB) Review

```go
func (t *TMDB) Review(id string) (*Review, error)
```
Review retrieves information about a specific review

#### func (*TMDB) SearchCollection

```go
func (t *TMDB) SearchCollection(query string, params ...option) (*SearchCollection, error)
```
SearchCollection searches the TMDB Movie database for the name given

#### func (*TMDB) SearchCompany

```go
func (t *TMDB) SearchCompany(query string, params ...option) (*SearchCompany, error)
```
SearchCompany searches the TMDB Movie database for the name given

#### func (*TMDB) SearchKeywords

```go
func (t *TMDB) SearchKeywords(query string, params ...option) (*SearchKeywords, error)
```
SearchKeywords search the TMDB database for the tersm given

#### func (*TMDB) SearchMovie

```go
func (t *TMDB) SearchMovie(query string, params ...option) (*SearchMovie, error)
```
SearchMovie searches the TMDB Movie database for the name given

#### func (*TMDB) SearchPerson

```go
func (t *TMDB) SearchPerson(query string, params ...option) (*SearchPeople, error)
```
SearchPerson searches the TMDB people database for the name given

#### func (*TMDB) SearchTV

```go
func (t *TMDB) SearchTV(query string, params ...option) (*SearchTV, error)
```
SearchTV searches the TMDB TV show database for the name given

#### func (*TMDB) TVAggregateCredits

```go
func (t *TMDB) TVAggregateCredits(id int64, params ...option) (*TVAggregateCredits, error)
```
TVAggregateCredits retrieves the aggregated credits for a TV show

#### func (*TMDB) TVAiringToday

```go
func (t *TMDB) TVAiringToday(params ...option) (*SearchTV, error)
```
TVAiringToday retrieves the TV show airing today

#### func (*TMDB) TVAlternativeTitles

```go
func (t *TMDB) TVAlternativeTitles(id int64, params ...option) (*AlternativeTitles, error)
```
TVAlternativeTitles retrieves the alternative titles for a TV show

#### func (*TMDB) TVChanges

```go
func (t *TMDB) TVChanges(id int64, params ...option) (*EntryChanges, error)
```
TVChanges retrieves the list of changes for a TV show

#### func (*TMDB) TVContentRatings

```go
func (t *TMDB) TVContentRatings(id int64, params ...option) (*TVContentRatings, error)
```
TVContentRatings retrieves the content ratings for a TV show

#### func (*TMDB) TVCredits

```go
func (t *TMDB) TVCredits(id int64, params ...option) (*Credits, error)
```
TVCredits retrieves the credits for a TV show

#### func (*TMDB) TVEpisode

```go
func (t *TMDB) TVEpisode(id int64, season int64, episode int64, params ...option) (*TVEpisode, error)
```
TVEpisode retrieves the details for an episode of a TV show

#### func (*TMDB) TVEpisodeChanges

```go
func (t *TMDB) TVEpisodeChanges(id int64, season int64, episode int64, params ...option) (*EntryChanges, error)
```
TVEpisodeChanges retrieves changes for an episode of a TV show

#### func (*TMDB) TVEpisodeCredits

```go
func (t *TMDB) TVEpisodeCredits(id int64, season int64, episode int64, params ...option) (*TVEpisodeCredits, error)
```
TVEpisodeCredits retrieves the credit information for an episode of a TV show

#### func (*TMDB) TVEpisodeExternalIDs

```go
func (t *TMDB) TVEpisodeExternalIDs(id int64, season int64, episode int64, params ...option) (*TVEpisodeExternalIDs, error)
```
TVEpisodeExternalIDs retrieves all of the known external IDs for an episode of a
TV show

#### func (*TMDB) TVEpisodeGroup

```go
func (t *TMDB) TVEpisodeGroup(id int64, params ...option) (*TVEpisodeGroup, error)
```
TVEpisodeGroup retrieves the details for an episode group for a TV show

#### func (*TMDB) TVEpisodeGroups

```go
func (t *TMDB) TVEpisodeGroups(id int64, params ...option) (*TVEpisodeGroups, error)
```
TVEpisodeGroups retrieves all of the episode groups for a TV show

#### func (*TMDB) TVEpisodeImages

```go
func (t *TMDB) TVEpisodeImages(id int64, season int64, episode int64, params ...option) (*Images, error)
```
TVEpisodeImages retrieves all of the images for an episode of a TV show

#### func (*TMDB) TVEpisodeTranslations

```go
func (t *TMDB) TVEpisodeTranslations(id int64, season int64, episode int64, params ...option) (*Translations, error)
```
TVEpisodeTranslations retrieves all of the translations for an episode of a TV
show

#### func (*TMDB) TVEpisodeVideos

```go
func (t *TMDB) TVEpisodeVideos(id int64, season int64, episode int64, params ...option) (*Videos, error)
```
TVEpisodeVideos retrieves all of the videos for an episode of a TV show

#### func (*TMDB) TVExternalIDs

```go
func (t *TMDB) TVExternalIDs(id int64, params ...option) (*ExternalIDs, error)
```
TVExternalIDs retrieves all of the external ids for a TV show

#### func (*TMDB) TVGenres

```go
func (t *TMDB) TVGenres(params ...option) (Genres, error)
```
TVGenres retrieves the official genres for movies

#### func (*TMDB) TVImages

```go
func (t *TMDB) TVImages(id int64, params ...option) (*Images, error)
```
TVImages retrieves all of the images for a TV show

#### func (*TMDB) TVKeywords

```go
func (t *TMDB) TVKeywords(id int64) (*Keywords, error)
```
TVKeywords retrieves all of the keywords for a TV show

#### func (*TMDB) TVLatest

```go
func (t *TMDB) TVLatest() (*TVShow, error)
```
TVLatest retrieves the latest TV show added to the database

#### func (*TMDB) TVOnTheAir

```go
func (t *TMDB) TVOnTheAir(params ...option) (*SearchTV, error)
```
TVOnTheAir retrieves TV show that are airing within the next 7 days

#### func (*TMDB) TVPopular

```go
func (t *TMDB) TVPopular(params ...option) (*SearchTV, error)
```
TVPopular retrieves a list of popular TV shows

#### func (*TMDB) TVRecommendations

```go
func (t *TMDB) TVRecommendations(id int64, params ...option) (*SearchTV, error)
```
TVRecommendations retrieves all of the recommendations for a TV show

#### func (*TMDB) TVReviews

```go
func (t *TMDB) TVReviews(id int64, params ...option) (*Reviews, error)
```
TVReviews retrieves reviews for a TV show

#### func (*TMDB) TVScreenedTheatrically

```go
func (t *TMDB) TVScreenedTheatrically(id int64) (*TVScreenedTheatrically, error)
```
TVScreenedTheatrically retrieves all of the episodes of a TV show that were
screened theatrically

#### func (*TMDB) TVSeason

```go
func (t *TMDB) TVSeason(id int64, season int64, params ...option) (*TVSeason, error)
```
TVSeason retrieves the details for a season of a TV show

#### func (*TMDB) TVSeasonAggregateCredits

```go
func (t *TMDB) TVSeasonAggregateCredits(id int64, season int64, params ...option) (*TVAggregateCredits, error)
```
TVSeasonAggregateCredits retrieves the aggregate credits for a season of a TV
show

#### func (*TMDB) TVSeasonChanges

```go
func (t *TMDB) TVSeasonChanges(id int64, season int64, params ...option) (*EntryChanges, error)
```
TVSeasonChanges retrieves changes for a season of a TV show

#### func (*TMDB) TVSeasonCredits

```go
func (t *TMDB) TVSeasonCredits(id int64, season int64, params ...option) (*Credits, error)
```
TVSeasonCredits retrieves credits for a season of a TV show

#### func (*TMDB) TVSeasonExternalIDs

```go
func (t *TMDB) TVSeasonExternalIDs(id int64, season int64, params ...option) (*ExternalIDs, error)
```
TVSeasonExternalIDs retrieves all of the external ids for a season of a TV show

#### func (*TMDB) TVSeasonImages

```go
func (t *TMDB) TVSeasonImages(id int64, season int64, params ...option) (*Images, error)
```
TVSeasonImages retrieves all of the images for a season of a TV show

#### func (*TMDB) TVSeasonTranslations

```go
func (t *TMDB) TVSeasonTranslations(id int64, season int64, params ...option) (*Translations, error)
```
TVSeasonTranslations retrieves all of the translations for a season of a TV show

#### func (*TMDB) TVSeasonVideos

```go
func (t *TMDB) TVSeasonVideos(id int64, season int64, params ...option) (*Videos, error)
```
TVSeasonVideos retrieves all of the videos for a season of a TV show

#### func (*TMDB) TVShow

```go
func (t *TMDB) TVShow(id int64, params ...option) (*TVShow, error)
```
TVShow retrieves the details of a TV Show

#### func (*TMDB) TVSimilar

```go
func (t *TMDB) TVSimilar(id int64, params ...option) (*SearchTV, error)
```
TVSimilar retrieves all of the similar TV shows

#### func (*TMDB) TVTopRated

```go
func (t *TMDB) TVTopRated(params ...option) (*SearchTV, error)
```
TVTopRated retrieves a list of the top rated TV shows

#### func (*TMDB) TVTranslations

```go
func (t *TMDB) TVTranslations(id int64) (*Translations, error)
```
TVTranslations retrieves all of the translations that exist for a show

#### func (*TMDB) TVVideos

```go
func (t *TMDB) TVVideos(id int64, params ...option) (*Videos, error)
```
TVVideos retrieves all of the videos for a TV shows

#### func (*TMDB) TVWatchProviders

```go
func (t *TMDB) TVWatchProviders(id int64) (*WatchProviders, error)
```
TVWatchProviders retrieves all of the watch providers that exist for a show

#### func (*TMDB) Timezones

```go
func (t *TMDB) Timezones() (*Timezones, error)
```
Timezones retrieves the list of timezones

#### func (*TMDB) TrendingDay

```go
func (t *TMDB) TrendingDay() (*SearchMovie, error)
```
TrendingDay retrieves a list of trending items over the last day

#### func (*TMDB) TrendingMoviesDay

```go
func (t *TMDB) TrendingMoviesDay() (*SearchMovie, error)
```
TrendingMoviesDay retrieves a list of trending movies over the last day

#### func (*TMDB) TrendingMoviesWeek

```go
func (t *TMDB) TrendingMoviesWeek() (*SearchMovie, error)
```
TrendingMoviesWeek retrieves a list of trending movies over the last week

#### func (*TMDB) TrendingPeopleDay

```go
func (t *TMDB) TrendingPeopleDay() (*SearchPeople, error)
```
TrendingPeopleDay retrieves a list of trending people over the last day

#### func (*TMDB) TrendingPeopleWeek

```go
func (t *TMDB) TrendingPeopleWeek() (*SearchPeople, error)
```
TrendingPeopleWeek retrieves a list of trending people over the last week

#### func (*TMDB) TrendingTVDay

```go
func (t *TMDB) TrendingTVDay() (*SearchTV, error)
```
TrendingTVDay retrieves a list of trending TV shows over the last day

#### func (*TMDB) TrendingTVWeek

```go
func (t *TMDB) TrendingTVWeek() (*SearchTV, error)
```
TrendingTVWeek retrieves a list of trending tv shows over the last week

#### func (*TMDB) TrendingWeek

```go
func (t *TMDB) TrendingWeek() (*SearchMovie, error)
```
TrendingWeek retrieves a list of trending items over the last week

#### type TVAggregateCredits

```go
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
```

TVAggregateCredits contains the aggregated credits for a TV show

#### type TVContentRatings

```go
type TVContentRatings struct {
	Results []struct {
		Country string `json:"iso_3166_1"`
		Rating  string `json:"rating"`
	} `json:"results"`
	ID int64 `json:"id"`
}
```

TVContentRatings containts the content ratings for a TV show

#### type TVEpisode

```go
type TVEpisode struct {
	Episode
	Crew       []CrewCredit
	GuestStars []CastCredit
}
```

TVEpisode contains basic information about an episode of a TV show

#### type TVEpisodeCredits

```go
type TVEpisodeCredits struct {
	ID         int64
	Cast       []CastCredit `json:"cast"`
	Crew       []CrewCredit `json:"crew"`
	GuestStars []CastCredit `json:"guest_stars"`
}
```

TVEpisodeCredits contains the credit information for an episode of a TV show

#### type TVEpisodeExternalIDs

```go
type TVEpisodeExternalIDs struct {
	IMDB        *string `json:"imdb_id"`
	FreebaseMID *string `json:"freebase_mid"`
	FreebaseID  *string `json:"freebase_id"`
	TVDB        *int64  `json:"tvdb_id"`
	TVRage      *int64  `json:"tbrage_id"`
	ID          int64   `json:"id"`
}
```

TVEpisodeExternalIDs contains all of the known external IDs for an episode of a
TV show

#### type TVEpisodeGroup

```go
type TVEpisodeGroup struct {
	Description  string `json:"description"`
	EpisodeCount int64  `json:"episode_count"`
	GroupCount   int64  `json:"group_count"`
	Groups       []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Order    int64  `json:"order"`
		Episodes []struct {
			Episode
			Order int64 `json:"int64"`
		} `json:"episodes"`
		Locked bool `json:"locked"`
	} `json:"groups"`
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Network ProductionCompany `json:"network"`
	Type    int64             `json:"type"`
}
```

TVEpisodeGroup contains the details for an episode groups for a TV show

#### type TVEpisodeGroups

```go
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
```

TVEpisodeGroups contains all of the episode groups for a TV show

#### type TVOrMovie

```go
type TVOrMovie struct {
	PosterPath       *string  `json:"poster_path"`
	Adult            bool     `json:"adult"`
	Overview         string   `json:"overview"`
	ReleaseDate      string   `json:"release_date"`
	OriginalTitle    string   `json:"original_title"`
	GenreIDs         []int64  `json:"genre_ids"`
	ID               int64    `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	Title            string   `json:"title"`
	BackdropPath     *string  `json:"backdrop_path"`
	Popularity       float64  `json:"popularity"`
	VoteCount        int64    `json:"vote_count"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
}
```

TVOrMovie contains information which could be for either a TV show or a movie

#### type TVResult

```go
type TVResult struct {
	PosterPath       *string  `json:"poster_path"`
	Popularity       float64  `json:"popularity"`
	ID               int64    `json:"id"`
	BackdropPath     string   `json:"backdrop_path"`
	VoteAverage      float64  `json:"vote_average"`
	Overview         string   `json:"overview"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	GenreIDs         []int64  `json:"genre_ids"`
	OriginalLanguage string   `json:"original_language"`
	VoteCount        int64    `json:"vote_count"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
}
```

TVResult stores a single result of a TV search

#### type TVScreenedTheatrically

```go
type TVScreenedTheatrically struct {
	ID      int64 `json:"id"`
	Results []struct {
		ID            int64 `json:"id"`
		EpisodeNumber int64 `json:"episode_number"`
		SeasonNumber  int64 `json:"season_number"`
	} `json:"results"`
}
```

TVScreenedTheatrically contains all of the episodes of a TV show that were
screened theatrically

#### type TVSeason

```go
type TVSeason struct {
	StrID        string      `json:"_id"`
	AirDate      string      `json:"air_date"`
	Episodes     []TVEpisode `json:"episodes"`
	Name         string      `json:"name"`
	Overview     string      `json:"overview"`
	ID           int64       `json:"id"`
	PosterPath   *string     `json:"poster_path"`
	SeasonNumber int64       `json:"season_number"`
}
```

TVSeason contains the details for a season of a TV show

#### type TVShow

```go
type TVShow struct {
	BackdropPath *string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int64   `json:"id"`
		CreditID    string  `json:"credit_id"`
		Name        string  `json:"name"`
		Gender      int64   `json:"gender"`
		ProfilePath *string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime      []int64             `json:"episode_run_time"`
	FirstAirDate        string              `json:"first_air_date"`
	Genres              Genres              `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int64               `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         string              `json:"last_air_date"`
	LastEpisodeToAir    Episode             `json:"last_episode_to_air"`
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
```

TVShow contains all of the information about a TV Show

#### type Timezone

```go
type Timezone string
```

Timezone is used to set the correct timezone for date limiting filters

#### type Timezones

```go
type Timezones []struct {
	Country string   `json:"iso_3166_1"`
	Zones   []string `json:"zones"`
}
```

Timezones represents the list of timezones used by TMDB

#### type Translations

```go
type Translations struct {
	ID           int64 `json:"id"`
	Translations []struct {
		Country     string `json:"iso_3166_1"`
		Language    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title    string `json:"title"`
			Overview string `json:"overview"`
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}
```

Translations represents collection translations data

#### type Videos

```go
type Videos struct {
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
```

Videos contains all of the videos related to a movie

#### type VoteAverageGTE

```go
type VoteAverageGTE float64
```

VoteAverageGTE filters by only allowing results with a vote average greater than
(or equal to) that specified

#### type VoteAverageLTE

```go
type VoteAverageLTE float64
```

VoteAverageLTE filters by only allowing results with a vote average less than
(or equal to) that specified

#### type VoteCountGTE

```go
type VoteCountGTE float64
```

VoteCountGTE filters by only allowing results with a vote count greater than (or
equal to) that specified

#### type VoteCountLTE

```go
type VoteCountLTE float64
```

VoteCountLTE filters by only allowing results with a vote count less than (or
equal to) that specified

#### type WatchProviderData

```go
type WatchProviderData struct {
	DisplayPriority int64  `json:"display_priority"`
	LogoPath        string `json:"logo_path"`
	ProviderID      int64  `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
}
```

WatchProviderData contains information about the provider of a video service

#### type WatchProviders

```go
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
```

WatchProviders contains all of the information about where and how to watch a
movie

#### type WatchRegion

```go
type WatchRegion string
```

WatchRegion filters results by the specified regions

#### type WithCast

```go
type WithCast string
```

WithCast specifies a comma seperated list of person IDs, any of which must be
included in each result

#### type WithCompanies

```go
type WithCompanies string
```

WithCompanies specifies a comma seperated list of company ids to limit the
results to

#### type WithCrew

```go
type WithCrew string
```

WithCrew specifies a comma seperated list of person IDs, any of which must be
included in each result

#### type WithGenres

```go
type WithGenres string
```

WithGenres specifies a comma seperated list of genre ids to limit the results to

#### type WithKeywords

```go
type WithKeywords string
```

WithKeywords filters results by the specified keywords

#### type WithNetworks

```go
type WithNetworks string
```

WithNetworks specifies a comma seperated list of network IDs to limit the
results to

#### type WithOriginalLanguage

```go
type WithOriginalLanguage string
```

WithOriginalLanguage specifies and ISO 639-1 string to filter results by
original language

#### type WithPeople

```go
type WithPeople string
```

WithPeople specifies a comma seperated list of person IDs, any of which must be
included in each result

#### type WithReleaseType

```go
type WithReleaseType int64
```

WithReleaseType filters results with the specified list type

#### type WithRuntimeGTE

```go
type WithRuntimeGTE int64
```

WithRuntimeGTE filters by only allowing results with a runtime greater than (or
equal to) that specified

#### type WithRuntimeLTE

```go
type WithRuntimeLTE int64
```

WithRuntimeLTE filters by only allowing results with a runtime less than (or
equal to) that specified

#### type WithWatchProviders

```go
type WithWatchProviders string
```

WithWatchProviders filters results by the specified watch proviers

#### type WithoutGenres

```go
type WithoutGenres string
```

WithoutGenres specifies a comma seperated list of genre ids to filter from the
results

#### type WithoutKeywords

```go
type WithoutKeywords string
```

WithoutKeywords filters results by the specified keywords

#### type Year

```go
type Year int64
```

Year filters the results by the specified year
