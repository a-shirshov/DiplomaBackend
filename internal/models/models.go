package models

type User struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	About string `json:"about,omitempty"`
	ImgUrl string `json:"imgUrl,omitempty"`
}

type Event struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	About string `json:"about,omitempty"`
	Category string `json:"category,omitempty"`
	Tags []string `json:"tags,omitempty"`
	SpecialInfo string `json:"specialInfo,omitempty"`
}

type Tokens struct {
	AccessToken string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}