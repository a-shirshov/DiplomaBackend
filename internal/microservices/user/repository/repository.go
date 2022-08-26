package repository

import (
	"Diploma/internal/models"
	"Diploma/internal/errors"
	"log"

	"github.com/jackc/pgx"
)

type UserRepository struct {
	db *pgx.ConnPool
}

func NewUserRepository(db *pgx.ConnPool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (uR *UserRepository) GetUser(Id int) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow("GetUserQuery",Id).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		return nil, errors.ErrPostgres
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}

func (uR *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow("GetUserByEmailQuery",email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
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
		err = uR.db.QueryRow("UpdateUserWithoutImgUrlQuery", &user.Name, &user.Surname, &user.About, &userId).Scan(
			&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl,
		)
	} else {
		err = uR.db.QueryRow("UpdateUserQuery", &user.Name, &user.Surname, &user.About, &user.ImgUrl, &userId).Scan(
			&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl,
		)
	}
	
	if err != nil {
		log.Print(err)
		return nil, errors.ErrPostgres
	}

	return models.ToUserModel(userDB), nil
}