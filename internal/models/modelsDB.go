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