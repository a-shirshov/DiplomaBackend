package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	GetUserQuery = `select id, name, surname, about, img_url from "user" where id = $1;`
	UpdateUserWithoutImgUrlQuery = `update "user" set name = $1, surname = $2, about = $3 where id = $4 returning id, name, surname, email, about, img_url;`
	UpdateUserQuery = `update "user" set name = $1, surname = $2, about = $3, img_url = $4 where id = $5 returning id, name, surname, email, about, img_url;`
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (uR *UserRepository) GetUser(id int) (*models.User, error) {
	user := models.User{}
	err := uR.db.Get(&user, GetUserQuery, id)
	if err != nil {
		log.Println(err)
		return &user, customErrors.ErrPostgres
	}
	return &user, nil
}

func (uR *UserRepository) UpdateUser(inputUser *models.User) (*models.User, error) {
	outputUser := models.User{}
	err := uR.db.QueryRowx(UpdateUserQuery, &inputUser.Name, &inputUser.Surname, &inputUser.About, &inputUser.ImgUrl, &inputUser.ID).StructScan(&outputUser)
	if err != nil {
		log.Print(err)
		return &outputUser, customErrors.ErrPostgres
	}
	return &outputUser, nil
}
