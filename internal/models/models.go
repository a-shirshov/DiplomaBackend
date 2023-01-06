package models

type LoginUser struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	City string `json:"city" db:"city"`
	About string `json:"about" validate:"max=120"`
	ImgUrl string `json:"imgUrl" db:"img_url"`
}

type Event struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name" validate:"required,max = 100"`
	Description string `json:"description" validate:"required,max = 100"`
	About string `json:"about,omitempty" validate:"required,max = 250"`
	Category string `json:"category,omitempty" validate:"required,max=20"`
	Tags []string `json:"tags,omitempty"`
	SpecialInfo string `json:"specialInfo,omitempty"`
	ImgUrl string `json:"imgUrl,omitempty" db:"img_url"`
}

type Place struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name" validate:"required,max = 100"`
	Description string `json:"description" validate:"required,max = 100"`
	About string `json:"about,omitempty" validate:"required,max = 250"`
	Category string `json:"category,omitempty" validate:"required,max=20"`
	ImgUrl string `json:"imgUrl,omitempty" db:"img_url"`
}

type Tokens struct {
	AccessToken string `json:"access_token,omitempty" example:"22f37ea5-2d12-4309-afbe-17783b44e24f"`
	RefreshToken string `json:"refresh_token" example:"4ffc5f18-99d8-47f6-8141-faf2c2f5a24e"`
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
    UserId   int
}

type KudaGoEvents struct {
	Results []KudaGoResult `json:"results"`
}

type KudaGoResult struct {
	ID int `json:"id"`
	Dates []KudaGoDate `json:"dates"`
	Title   string `json:"title"`
	Images []KudaGoImage `json:"images"`
	Location struct {
		Slug string `json:"slug"`
	} `json:"location"`
	Place struct{
		ID int `json:"id"`
	} `json:"place"`
}

type KudaGoDate struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type KudaGoImage struct {
	Image string `json:"image"`
}

type MyEvents struct {
	Events []MyEvent
}

type MyEvent struct {
	ID int `json:"id"`
	KudaGoID int `json:"kudago_id"`
	Title   string `json:"title"`
	Start int `json:"start"`
	End   int `json:"end"`
	Location string `json:"location"`
	Image string `json:"image"`
	Place int `json:"place"`
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
	Event MyEvent `json:"event"`
	Place KudaGoPlaceResult `json:"place"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

type UserWithTokens struct {
	User User `json:"user"`
	Tokens Tokens `json:"tokens"`
}