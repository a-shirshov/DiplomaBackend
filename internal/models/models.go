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

// swagger:model
type Event struct {
	// the id of event
	//
	// min: 1
	ID int `json:"id,omitempty"`
	// the name of event
	//
	// required: true
	//
	// max length: 80
	//
	// Fairytale in theatre
	Name string `json:"name,omitempty"`
	// short info about event
	//
	// required: true
	//
	// max length: 80
	//
	// A famous fairytale in local theatre
	Description string `json:"description,omitempty"`
	// long info about event
	//
	// required: true
	//
	// max length: 500
	//
	// example: very interesting ...
	About string `json:"about,omitempty"`
	// category of event (still in dev)
	//
	// required: true
	//
	// max length: 50
	//
	// example: theatre
	Category string `json:"category,omitempty"`
	// tags of event
	//
	// max items: 5
	//
	// items max length : 30
	Tags []string `json:"tags,omitempty"`
	// important info about event
	//
	// max length: 150
	// example: For people older than 18
	SpecialInfo string `json:"specialInfo,omitempty"`
}

// swagger:model
type Tokens struct {
	// access token for user
	// example: b53bcc38-efce-46d9-b7b2-fbeb175b91ab
	// max length: 50
	AccessToken string `json:"access_token,omitempty"`
	// refresh token for user
	// example: d0a93fb4-26fe-4980-836a-b0d481f0aa68
	// max length: 50
	RefreshToken string `json:"refresh_token,omitempty"`
}