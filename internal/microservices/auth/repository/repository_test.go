package repository

import (
	"Diploma/internal/models"
	"Diploma/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type createUserTest struct {
	inputUser *models.User
	outputUser *models.User
	outputErr error
}

var createUserTests = []createUserTest{
	{
		&models.User{
			Name: "User_1",
			Surname: "Surname_1",
			Email: "Email_1",
			Password: "password_1",
		},
		&models.User{
			Name: "User_1",
			Surname: "Surname_1",
			Email: "Email_1",
		}, nil,
	},
}

type getUserByEmailTest struct {
	email string
	outputUser *models.User
	outputErr error
}

var getUserByEmailTests = []getUserByEmailTest{
	{
		"Email_1", &models.User{
			ID: 1,
			Name: "User_1",
			Surname: "Surname_1",
			Email:  "Email_1",
			Password: "Password_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		}, nil,
	},
}

type updateUserTest struct {
	
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repositoryTest := NewAuthRepository(sqlxDB)

	for _, test := range createUserTests {
		mock.ExpectExec(utils.CreateUserQuery).
			WithArgs(
				test.inputUser.Name,
				test.inputUser.Surname,
				test.inputUser.Email,
				test.inputUser.Password).
				WillReturnResult(sqlmock.NewResult(1,1))
		
		out, dbErr := repositoryTest.CreateUser(test.inputUser)
		assert.Equal(t, test.outputUser, out)
		assert.Nil(t, dbErr)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repositoryTest := NewAuthRepository(sqlxDB)
	
	columns := []string{"id", "name", "surname", "email", "password", "about", "imgurl"}

	for _, test := range getUserByEmailTests {
		rows := sqlmock.NewRows(columns).
			AddRow(
				test.outputUser.ID,
				test.outputUser.Name,
				test.outputUser.Surname,
				test.outputUser.Email,
				test.outputUser.Password,
				test.outputUser.About,
				test.outputUser.ImgUrl)

		mock.ExpectQuery(utils.GetUserByEmailQuery).
			WithArgs(test.email).
			RowsWillBeClosed().
			WillReturnRows(rows)

		out, dbErr := repositoryTest.GetUserByEmail(test.email)
		assert.Equal(t, test.outputUser, out)
		assert.Nil(t, dbErr)
	}
}