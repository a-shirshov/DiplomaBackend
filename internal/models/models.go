package models

type LoginUser struct {
	Email    string `json:"email" valid:"required,email" san:"xss"`
	Password string `json:"password" valid:"optional,minstringlength(8)" san:"xss"`
}

type User struct {
	ID          int    `json:"id,omitempty" san:"xss"`
	Name        string `json:"name" valid:"required" san:"xss"`
	Surname     string `json:"surname" valid:"required" san:"xss"`
	Email       string `json:"email,omitempty" valid:"required,email" san:"xss"`
	Password    string `json:"password,omitempty" valid:"optional,minstringlength(8)" san:"xss"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth" san:"xss"`
	City        string `json:"city" db:"city" san:"xss"`
	About       string `json:"about" valid:"max=120" san:"xss"`
	ImgUrl      string `json:"imgUrl" db:"img_url" san:"xss"`
}

type Event struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name" valid:"required,max = 100" san:"xss"`
	Description string   `json:"description" valid:"required,max = 100" san:"xss"`
	About       string   `json:"about,omitempty" valid:"required,max = 250" san:"xss"`
	Category    string   `json:"category,omitempty" valid:"required,max=20" san:"xss"`
	Tags        []string `json:"tags,omitempty" san:"xss"`
	SpecialInfo string   `json:"specialInfo,omitempty" san:"xss"`
	ImgUrl      string   `json:"imgUrl,omitempty" db:"img_url" san:"xss"`
}

type Place struct {
	ID          int    `json:"id,omitempty" san:"xss"`
	Name        string `json:"name" valid:"required,max = 100" san:"xss"`
	Description string `json:"description" valid:"required,max = 100" san:"xss"`
	About       string `json:"about,omitempty" valid:"required,max = 250" san:"xss"`
	Category    string `json:"category,omitempty" valid:"required,max=20" san:"xss"`
	ImgUrl      string `json:"imgUrl,omitempty" db:"img_url" san:"xss"`
}

type Tokens struct {
	AccessToken  string `json:"access_token,omitempty" example:"22f37ea5-2d12-4309-afbe-17783b44e24f" san:"xss"`
	RefreshToken string `json:"refresh_token" example:"4ffc5f18-99d8-47f6-8141-faf2c2f5a24e" san:"xss"`
}

type RedeemCodeStruct struct {
	Email      string `json:"email" valid:"email" san:"xss"`
	RedeemCode int    `json:"redeem_code,omitempty" san:"xss"`
	Password   string `json:"password,omitempty" valid:"optional,minstringlength(8)" san:"xss"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     int
}

//////
type KudaGoEvents struct {
	Results []KudaGoResult `json:"results"`
}

type KudaGoResult struct {
	ID       int           `json:"id"`
	Dates    []KudaGoDate  `json:"dates"`
	Title    string        `json:"title"`
	Images   []KudaGoImage `json:"images"`
	Location struct {
		Slug string `json:"slug"`
	} `json:"location"`
	Place struct {
		ID int `json:"id"`
		IsStub bool `json:"is_stub"`
	} `json:"place"`
	Description string `json:"description,omitempty"`
	Price       string `json:"price,omitempty"`
}

type KudaGoDate struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type KudaGoImage struct {
	Image string `json:"image"`
}
////

type MyEvents struct {
	Events []MyEvent `json:"events"`
}

type MyEvent struct {
	ID          int    `json:"id" db:"id"`
	KudaGoID    int    `json:"kudago_id" db:"kudago_id"`
	Title       string `json:"title" db:"title"`
	Start       int    `json:"start" db:"start_time"`
	End         int    `json:"end" db:"end_time"`
	Location    string `json:"location" db:"location"`
	Image       string `json:"image" db:"image"`
	Place       int    `json:"place" db:"place_id"`
	Description string `json:"description" db:"description"`
	Price       string `json:"price" db:"price"`
	Vector	    []float64 `json:"-" db:"vector"`
	VectorTitle []float64 `json:"-" db:"vector_title"`
	IsLiked		bool   `json:"is_liked" db:"is_liked"`
}

type MyPlace struct {
	ID          int   `json:"id" db:"id"`
	KudaGoID    int   `json:"kudago_id" db:"kudago_id"`
	Title      string `json:"title" db:"title"`
	Address    string `json:"address" db:"address"`
	SiteURL    string `json:"site_url" db:"site_url"`
	ForeignURL string `json:"foreign_url" db:"foreign_url"`
	Phone	   string `json:"phone" db:"phone"`
	Timetable  string `json:"timetable" db:"timetable"`
	Coords     struct {
		Lat float64 `json:"lat" db:"lat"`
		Lon float64 `json:"lon" db:"lon"`
	} `json:"coords"`
}

type MyFullEvent struct {
	Event MyEvent `json:"event"`
	Place MyPlace `json:"place"`
}

type KudaGoPlaceResult struct {
	Title      string `json:"title"`
	Address    string `json:"address"`
	SiteURL    string `json:"site_url"`
	ForeignURL string `json:"foreign_url"`
	Coords     struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coords"`
}

type KudaGoPlaceAndEvent struct {
	Event       MyEvent           `json:"event"`
	Place       KudaGoPlaceResult `json:"place"`
	PeopleCount int               `json:"peopleCount"`
	IsGoing     bool              `json:"is_going"`
	IsFavourite bool              `json:"is_favourite"`
}
////

type AutoGenerated struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		ID              int    `json:"id"`
		Title           string `json:"title"`
		FavoritesCount  int    `json:"favorites_count"`
		CommentsCount   int    `json:"comments_count"`
		Description     string `json:"description"`
		ItemURL         string `json:"item_url"`
		DisableComments bool   `json:"disable_comments"`
		Ctype           string `json:"ctype"`
		Place           struct {
			ID int `json:"id"`
		} `json:"place"`
		Daterange struct {
			StartDate        int           `json:"start_date"`
			StartTime        int           `json:"start_time"`
			Start            int           `json:"start"`
			EndDate          interface{}   `json:"end_date"`
			EndTime          int           `json:"end_time"`
			End              int           `json:"end"`
			IsContinuous     bool          `json:"is_continuous"`
			IsEndless        bool          `json:"is_endless"`
			IsStartless      bool          `json:"is_startless"`
			Schedules        []interface{} `json:"schedules"`
			UsePlaceSchedule bool          `json:"use_place_schedule"`
		} `json:"daterange"`
		FirstImage struct {
			Image      string `json:"image"`
			Thumbnails struct {
				Six40X384 string `json:"640x384"`
				One44X96  string `json:"144x96"`
			} `json:"thumbnails"`
			Source struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"source"`
		} `json:"first_image"`
		AgeRestriction int `json:"age_restriction"`
	} `json:"results"`
}

type KudaGoSearchResults struct {
	Results  []KudaGoSearchResult `json:"results"`
}

type KudaGoSearchResult struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Place           struct {
		ID int `json:"id"`
		IsStub bool `json:"is_stub"`
	} `json:"place"`
	Daterange struct {
		Start            int           `json:"start"`
		End              int           `json:"end"`
	} `json:"daterange"`
	FirstImage struct {
		Image      string `json:"image"`
	} `json:"first_image"`
	AgeRestriction int `json:"age_restriction"`
}

type Message struct {
	Message string `json:"message,omitempty"`
}

type UserWithTokens struct {
	User   User   `json:"user"`
	Tokens Tokens `json:"tokens"`
}


