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
	UpdateUserQuery = `update "user" set name = $1, surname = $2, about = $3 where id = $4 returning id, name, surname, email, about, imgUrl`
	GetEventsQuery = `select id, name, description, about, category, tags, specialInfo from (
		select ROW_NUMBER() OVER (ORDER BY creationDate) as RowNum, * from "event") as eventsPaged 
		where RowNum Between 1 + $1 * ($2-1) and $1 * $2`
	GetEventQuery = `select id, name, description, about, category, tags, specialInfo from "event" where id = $1`
)

var queries = []Query {
	{
		Name: "CreateUserQuery",
		Query: CreateUserQuery,
	},
	{
		Name: "GetUserQuery",
		Query: GetUserQuery,
	},
	{
		Name: "GetUserByEmailQuery",
		Query: GetUserByEmailQuery,
	},
	{
		Name: "UpdateUserQuery",
		Query: UpdateUserQuery,
	},
	{
		Name: "GetEventsQuery",
		Query: GetEventsQuery,
	},
	{
		Name: "GetEventQuery",
		Query: GetEventQuery,
	},
}