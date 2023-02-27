package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

const (
	GetUserQuery = `select id, name, surname, about, img_url from "user" where id = $1;`
	UpdateUserQuery = `update "user" set name = $1, surname = $2, about = $3 where id = $4 returning id, name, surname, about, img_url;`
	UpdateUserImageQuery = `update "user" set img_url = $1 where id = $2 returning id, name, surname, about, img_url;`
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
	err := uR.db.Get(&user, GetUserQuery, &id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, customErrors.ErrUserNotFound
		}
		return &user, customErrors.ErrPostgres
	}
	return &user, nil
}

func (uR *UserRepository) UpdateUser(inputUser *models.User) (*models.User, error) {
	outputUser := models.User{}
	err := uR.db.QueryRowx(UpdateUserQuery, &inputUser.Name, &inputUser.Surname, &inputUser.About, &inputUser.ID).StructScan(&outputUser)
	if err != nil {
		return &outputUser, customErrors.ErrPostgres
	}
	return &outputUser, nil
}

func (uR *UserRepository) UpdateUserImage(userID int, imgUUID string) (*models.User, error) {
	outputUser := models.User{}
	err := uR.db.QueryRowx(UpdateUserQuery, &imgUUID, &userID).StructScan(&outputUser)
	if err != nil {
		return &outputUser, customErrors.ErrPostgres
	}
	return &outputUser, nil
}