package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"Diploma/utils/query"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (uR *AuthRepository) CreateUser(user *models.User) (*models.User, error) {
	err := uR.db.QueryRowx(query.CreateUserQuery, &user.Name, &user.Surname, &user.Email, &user.Password).Scan(&user.ID)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
			return nil, customErrors.ErrUserExists
		}
		return nil, customErrors.ErrPostgres
	}
	user.Password = ""
	return user, nil
}

func (uR *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow(query.GetUserByEmailQuery,email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		log.Print(err.Error())
		return nil, customErrors.ErrWrongEmail
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}