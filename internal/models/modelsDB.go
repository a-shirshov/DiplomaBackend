package models

import "database/sql"


type UserDB struct {
	ID int 
	Name string 
	Surname string 
	Email string 
	Password string 
	About sql.NullString 
	ImgUrl sql.NullString 
}

type PlaceDB struct {
	ID int 
	Name string 
	Description string 
	Category string
	About sql.NullString 
	ImgUrl sql.NullString 
}