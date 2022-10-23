package utils

import (

)

type Query struct {
	Name string
	Query string
}

const (
	CreateUserQuery = `insert into "user" (name, surname, email, password) values ($1, $2, $3, $4)`
	GetUserQuery = `select id, name, surname, email, about, imgUrl from "user" where id = $1`
	GetUserByEmailQuery = `select * from "user" where email = $1`
	UpdateUserWithoutImgUrlQuery = `update "user" set name = $1, surname = $2, about = $3 where id = $4 returning id, name, surname, email, about, imgUrl`
	UpdateUserQuery = `update "user" set name = $1, surname = $2, about = $3, imgUrl = $4 where id = $5 returning id, name, surname, email, about, imgUrl`

	GetPlacesQuery = `select id, name, description, about, category, imgUrl from (
		select ROW_NUMBER() OVER() as RowNum, * from "place") as placesPaged 
		where RowNum Between 1 + $1 * ($2-1) and $1 * $2`
	GetPlaceQuery = `select * from place where id = $1`

	GetEventsQuery = `select id, name, description, about, category, tags, specialInfo from (
		select ROW_NUMBER() OVER (ORDER BY creationDate) as RowNum, * from "event" where place_id = $1) as eventsPaged 
		where RowNum Between 1 + $2 * ($3-1) and $2 * $3`
	GetEventQuery = `select id, name, description, about, category, tags, specialInfo from "event" 
		where id = $1`
)

// var queries = []Query {
// 	{
// 		Name: "CreateUserQuery",
// 		Query: CreateUserQuery,
// 	},
// 	{
// 		Name: "GetUserQuery",
// 		Query: GetUserQuery,
// 	},
// 	{
// 		Name: "GetUserByEmailQuery",
// 		Query: GetUserByEmailQuery,
// 	},
// 	{
// 		Name: "UpdateUserWithoutImgUrlQuery",
// 		Query: UpdateUserWithoutImgUrlQuery,
// 	},
// 	{
// 		Name: "UpdateUserQuery",
// 		Query: UpdateUserQuery,
// 	},
// 	{
// 		Name: "GetEventsQuery",
// 		Query: GetEventsQuery,
// 	},
// 	{
// 		Name: "GetEventQuery",
// 		Query: GetEventQuery,
// 	},
// 	{
// 		Name: "GetPlacesQuery",
// 		Query: GetPlacesQuery,
// 	},
// 	{
// 		Name: "GetPlaceQuery",
// 		Query: GetPlaceQuery,
// 	},
// }