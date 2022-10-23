package repository

import (
	"Diploma/internal/errors"
	"Diploma/internal/models"
	"Diploma/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (uR *UserRepository) GetUser(Id int) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow(utils.GetUserQuery, Id).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		return nil, errors.ErrPostgres
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}

func (uR *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow(utils.GetUserByEmailQuery, email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.ErrPostgres
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}

func (uR *UserRepository) UpdateUser(userId int, user *models.User) (*models.User, error) {
	userDB := &models.UserDB{}
	var err error
	if user.ImgUrl == "" {
		err = uR.db.QueryRow(utils.UpdateUserWithoutImgUrlQuery, &user.Name, &user.Surname, &user.About, &userId).Scan(
			&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl,
		)
	} else {
		err = uR.db.QueryRow(utils.UpdateUserQuery, &user.Name, &user.Surname, &user.About, &user.ImgUrl, &userId).Scan(
			&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl,
		)
	}
	
	if err != nil {
		log.Print(err)
		return nil, errors.ErrPostgres
	}

	return models.ToUserModel(userDB), nil
}