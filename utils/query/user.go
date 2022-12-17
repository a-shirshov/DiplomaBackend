package query

const (
	CreateUserQuery = `insert into "user" (name, surname, email, password) values ($1, $2, $3, $4) returning id;`
	GetUserQuery = `select id, name, surname, email, about, imgUrl from "user" where id = $1;`
	GetUserByEmailQuery = `select * from "user" where email = $1;`
	UpdateUserWithoutImgUrlQuery = `update "user" set name = $1, surname = $2, about = $3 where id = $4 returning id, name, surname, email, about, imgUrl;`
	UpdateUserQuery = `update "user" set name = $1, surname = $2, about = $3, imgUrl = $4 where id = $5 returning id, name, surname, email, about, imgUrl;`
)