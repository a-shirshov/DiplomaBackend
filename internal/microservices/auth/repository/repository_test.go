package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type createUserTest struct {
	name	    string
	inputUser   *models.User
	beforeTest  func(sqlmock.Sqlmock)
	expectedUser  *models.User
	expectedError error
}

var createUserTests = []createUserTest{
	{
		"Successfully create user",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email: "ash@mail.ru",
			Password: "password",
			DateOfBirth: "2001-08-06",
			City: "msk",
			ImgUrl: "user_face.png",
		},
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id", "name", "surname", "email", "password", "date_of_birth", "city", "about", "img_url"}
			rows := mockSQL.NewRows(columns).AddRow(1, "Artyom", "Shirshov", "ash@mail.ru", "password", "2001-08-06", "msk", "", "user_face.png")
			mockSQL.ExpectQuery(regexp.QuoteMeta(CreateUserQuery)).
				WithArgs("Artyom", "Shirshov", "ash@mail.ru", "password", "2001-08-06", "msk", "user_face.png").
				RowsWillBeClosed().
				WillReturnRows(rows)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email: "ash@mail.ru",
			DateOfBirth: "2001-08-06",
			City: "msk",
			ImgUrl: "user_face.png",
		}, 
		nil,
	},
	{
		"User with this email exists",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email: "ash@mail.ru",
			Password: "password",
			DateOfBirth: "2001-08-06",
			City: "msk",
			ImgUrl: "user_face.png",
		},
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(CreateUserQuery)).
				WithArgs("Artyom", "Shirshov", "ash@mail.ru", "password", "2001-08-06", "msk", "user_face.png").
		        WillReturnError(errors.New(("(SQLSTATE 23505)")))
		},
		&models.User{}, 
		customErrors.ErrUserExists,
	},
	{
		"Database internal problem",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email: "ash@mail.ru",
			Password: "password",
			DateOfBirth: "2001-08-06",
			City: "msk",
			ImgUrl: "user_face.png",
		},
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(CreateUserQuery)).
				WithArgs("Artyom", "Shirshov", "ash@mail.ru", "password", "2001-08-06", "msk", "user_face.png").
				WillReturnError(errors.New("sql error"))
		},
		&models.User{},
		customErrors.ErrPostgres,
	},
}

type getUserByEmailTest struct {
	name string
	inputEmail string
	beforeTest func(sqlmock.Sqlmock)
	expectedUser *models.User
	expectedError error
}

var getUserByEmailTests = []getUserByEmailTest{
	{
		"Successfully found user by email",
		"ash@mail.ru",
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id", "name", "surname", "email", "password", "date_of_birth", "city", "about", "img_url"}
			rows := mockSQL.NewRows(columns).
				AddRow(1, "Artyom", "Shirshov", "ash@mail.ru", "password", "2001-08-06", "msk", "about", "img_url")

			mockSQL.ExpectQuery(regexp.QuoteMeta(GetUserByEmailQuery)).
				WithArgs("ash@mail.ru").
		        RowsWillBeClosed().
				WillReturnRows(rows)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			Password: "password",
			DateOfBirth: "2001-08-06",
			City: "msk",
			About: "about",
			ImgUrl: "img_url",
		}, 
		nil,
	},
	{
		"User not found by email",
		"ash@mail.ru",
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetUserByEmailQuery)).
				WithArgs("ash@mail.ru").
				WillReturnError(sql.ErrNoRows)
		},
		&models.User{},
		customErrors.ErrWrongEmail,
	},
	{
		"Database internal problem",
		"ash@mail.ru",
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetUserByEmailQuery)).
				WithArgs("ash@mail.ru").
				WillReturnError(errors.New("sql error"))
		},
		&models.User{},
		customErrors.ErrPostgres,
	},
}

func prepareTestEnvironment() (*AuthRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repositoryTest := NewAuthRepository(sqlxDB)

	return repositoryTest, mock, nil
}

func TestCreateUser(t *testing.T) {
	for _, test := range createUserTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualUser, actualErr := repositoryTest.CreateUser(test.inputUser)
			assert.Equal(t, test.expectedUser, actualUser)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}

func TestGetUserByEmail(t *testing.T) {
	for _, test := range getUserByEmailTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualUser, actualErr := repositoryTest.GetUserByEmail(test.inputEmail)
			assert.Equal(t, test.expectedUser, actualUser)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})
	}
}