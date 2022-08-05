package models

type LoginUser struct {
	Email string `json:"email" validate:"reqiured,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty"`
	About string `json:"about,omitempty" validate:"max=120"`
	ImgUrl string `json:"imgUrl,omitempty"`
}

type Event struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name" validate:"required,max = 100"`
	Description string `json:"description" validate:"required,max = 100"`
	About string `json:"about,omitempty" validate:"required,max = 250"`
	Category string `json:"category,omitempty" validate:"required,max=20"`
	Tags []string `json:"tags,omitempty"`
	SpecialInfo string `json:"specialInfo,omitempty"`
}

type Place struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name" validate:"required,max = 100"`
	Description string `json:"description" validate:"required,max = 100"`
	About string `json:"about,omitempty" validate:"required,max = 250"`
	Category string `json:"category,omitempty" validate:"required,max=20"`
	ImgUrl string `json:"img_url"`
}

type Tokens struct {
	AccessToken string `json:"access_token,omitempty" example:"22f37ea5-2d12-4309-afbe-17783b44e24f"`
	RefreshToken string `json:"refresh_token" example:"4ffc5f18-99d8-47f6-8141-faf2c2f5a24e"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

type UserWithTokens struct {
	User User `json:"user"`
	Tokens Tokens `json:"tokens"`
}