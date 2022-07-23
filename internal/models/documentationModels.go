package models

// A list of events returns in the response
// swagger:response eventsResponse
type eventsResponseWrapper struct {
	// All Events of the selected page
	// in: body
	Body []Event
}

//swagger:parameters ListEvents
type pageParameterWrapper struct {
	// The page of events from the database
	// in: path
	// required: true
	Page int `json:"page"`
}

//swagger:parameters id
type idParameterWrapper struct {
	//
}

// Empty Response
// swagger:response noContent
type noContent struct {
}

// Error message about Bad Request
// swagger:response badRequest
type badRequestMessage struct {
	// Bad Request
	// in: body
	// example: wrong model
	ErrorMessage string
}

// Error message about Server Error
// swagger:response serverError
type serverErrorMessage struct {
	// Internal Server Error
	// in: body
	// example: database is not working
	ErrorMessage string
}

type userSignUp struct {
	// Name of the user
	//
	// example: Artyom
	//
	// max length: 50
	Name string
	// Surname of the user
	//
	// example: Shirshov
	//
	// max length: 50
	Surname string
	// Email of the user
	//
	// example: artyom@mail.ru
	//
	// max length: 50
	Email string
	// Password of the user
	//
	// example: 12345678
	//
	// max length: 30
	Password string
}

// Data for registration
// swagger:parameters SignUpUser
type userRequestSignUpWrapper struct {
	// in: body
	Body userSignUp
}

// swagger:model
type userProfileResponse struct {
	// Name of the user
	//
	// example: Artyom
	//
	// max length: 50
	Name string `json:"name,omitempty"`
	// Surname of the user
	//
	// example: Shirshov
	//
	// max length: 50
	Surname string `json:"surname,omitempty"`
	// Email of the user
	//
	// example: artyom@mail.ru
	//
	// max length: 50
	Email string `json:"email,omitempty"`
	// About user
	//
	// example: about user
	//
	// max length: 100
	About string `json:"about,omitempty"`
	// ImgUrl for avatar
	//
	// example: about user
	//
	// max length: 100
	ImgUrl string `json:"imgurl,omitempty"`
}

// User profile from the database
// swagger:response userResponse
type userResponseWrapper struct {
	// User from the database
	// in: body
	Body userProfileResponse
}
// swagger:response tokensResponse
type tokensResponseWrapper struct {
	// in: body
	Body Tokens
}

// Get new access token
// swagger:parameters RefreshToken
type refreshTokenWrapper struct {
	// in: body
	Body refreshToken
}

// Refresh token
// swagger:model
type refreshToken struct {
	// example: d0a93fb4-26fe-4980-836a-b0d481f0aa68 
	Refresh_token string `json:"refresh_token,omitempty"`
}

// Updated info about user
// swagger:parameters UpdateUser
type userProfileRequest struct {
	// Updated user
	// in: body
	Body userProfileResponse
}