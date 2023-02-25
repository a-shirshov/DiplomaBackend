package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	log "Diploma/pkg/logger"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

const logMessage = "auth:repository:"

const (
	CreateUserQuery = `insert into "user" (name, surname, email, password, date_of_birth, city, img_url) values ($1, $2, $3, $4, $5, $6, $7) 
		returning id, name, surname, email, date_of_birth, city, img_url;`
	GetUserByEmailQuery = `select id, name, surname, email, password, date_of_birth, city, about, img_url from "user" where email = $1;`
	UpdatePasswordQuery = `update "user" set password = $1 where email = $2;`
	DeleteUserQuery = `delete from "user" where id = $1;`
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (aR *AuthRepository) CreateUser(inputUser *models.User) (*models.User, error) {
	message := logMessage + "CreateUser:"
	log.Debug(message + "started")
	user := models.User{}

	err := aR.db.QueryRowx(CreateUserQuery, 
		&inputUser.Name, 
		&inputUser.Surname, 
		&inputUser.Email, 
		&inputUser.Password, 
		&inputUser.DateOfBirth, 
		&inputUser.City, 
		&inputUser.ImgUrl).StructScan(&user)

	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
			return &models.User{}, customErrors.ErrUserExists
		}
		return &models.User{}, customErrors.ErrPostgres
	}
	user.Password = ""
	return &user, nil
}

func (aR *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	message := logMessage + "GetUserByEmail:"
	log.Debug(message + "started")
	user := models.User{}
	err := aR.db.Get(&user, GetUserByEmailQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, customErrors.ErrWrongEmail
		}
		log.Error(err)
		return &user, customErrors.ErrPostgres
	}
	return &user, nil
}

func (aR *AuthRepository) UpdatePassword(passwordHash string, email string) (error) {
	message := logMessage + "UpdatePassword:"
	log.Debug(message + "started")
	_, err := aR.db.Exec(UpdatePasswordQuery, passwordHash, email)
	return err
}

func (aR *AuthRepository) DeleteUser(userID int) (error) {
	message := logMessage + "DeleteUser:"
	log.Debug(message + "started")

	_, err := aR.db.Exec(DeleteUserQuery, userID)
	if err != nil {
		return customErrors.ErrPostgres
	}

	return nil
}