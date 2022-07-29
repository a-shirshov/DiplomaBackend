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
	AccessToken string `json:"access_token,omitempty" example:"22f37ea5-2d12-4309-afbe-17783b44e24f"`
	RefreshToken string `json:"refresh_token" binding:"required" example:"4ffc5f18-99d8-47f6-8141-faf2c2f5a24e"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

type UserWithTokens struct {
	User User `json:"user"`
	Tokens Tokens `json:"tokens"`
}