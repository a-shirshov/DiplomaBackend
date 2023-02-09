package repository

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type GetPeopleCountTest struct {
	name	    		string
	inputEventID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedPeopleCount int
	expectedError 		error
}

var GetPeopleCountTests = []GetPeopleCountTest{
	{
		"Succesfully get people count from kudago event",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"people_count"}
			rows := mockSQL.NewRows(columns).
				AddRow(10)

			mockSQL.ExpectQuery(regexp.QuoteMeta(GetPeopleCount)).
				WithArgs(1).
				RowsWillBeClosed().
				WillReturnRows(rows)
		},
		10,
		nil,
	},
	{
		"Event not found",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetPeopleCount)).
				WithArgs(1).
				WillReturnError(sql.ErrNoRows)
		},
		0,
		sql.ErrNoRows,
	},
}

type CreateKudaGoEventTest struct {
	name	    		string
	inputEventID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedError 		error
}

var CreateKudaGoEventTests = []CreateKudaGoEventTest {
	{
		"Successfully created event",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectExec(regexp.QuoteMeta(CreateKudaGoEvent)).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 1))
		},
		nil,
	},
}

type SwitchEventMeetingTest struct {
	name	    		string
	inputUserID			int
	inputEventID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedError 		error
}

var SwitchEventMeetingTests = []SwitchEventMeetingTest {
	{
		"Successfully create KudaGo Meeting",
		1,
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(CheckKudaGoMeeting)).
				WithArgs(1, 1).
				WillReturnError(sql.ErrNoRows)

			mockSQL.ExpectExec(regexp.QuoteMeta(CreateKudaGoMeeting)).
				WithArgs(1, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))
		},
		nil,
	},
}

type checkKudaGoMeetingTest struct {
	name	    		string
	inputUserID			int
	inputEventID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedIsGoing		bool
	expectedError 		error
}

var checkKudaGoMeetingTests = []checkKudaGoMeetingTest {
	{
		"Successfully created event",
		1,
		1,
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id"}
			rows := mockSQL.NewRows(columns).
				AddRow(1)

			mockSQL.ExpectQuery(regexp.QuoteMeta(CheckKudaGoMeeting)).
				WithArgs(1, 1).
				WillReturnRows(rows)
		},
		true,
		nil,
	},
}

type checkKudaGoFavouriteTest struct {
	name	    		string
	inputUserID			int
	inputEventID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedIsFavourite	bool
	expectedError 		error
}

var checkKudaGoFavouriteTests = []checkKudaGoFavouriteTest {
	{
		"Successfully created event",
		1,
		1,
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id"}
			rows := mockSQL.NewRows(columns).
				AddRow(1)

			mockSQL.ExpectQuery(regexp.QuoteMeta(CheckEventFavourite)).
				WithArgs(1, 1).
				WillReturnRows(rows)
		},
		true,
		nil,
	},
}

type SwitchEventFavouriteTest struct {
	name	    		string
	inputUserID			int
	inputEventID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedError 		error
}

var SwitchEventFavouriteTests = []SwitchEventFavouriteTest {
	{
		"Successfully create KudaGo Meeting",
		1,
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(CheckEventFavourite)).
				WithArgs(1, 1).
				WillReturnError(sql.ErrNoRows)

			mockSQL.ExpectExec(regexp.QuoteMeta(AddEventToFavourite)).
				WithArgs(1, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))
		},
		nil,
	},
}

func prepareTestEnvironment() (*EventRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repositoryTest := NewEventRepository(sqlxDB)

	return repositoryTest, mock, nil
}

func TestGetPeopleCount(t *testing.T) {
	for _, test := range GetPeopleCountTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualUser, actualErr := repositoryTest.GetPeopleCount(test.inputEventID)
			assert.Equal(t, test.expectedPeopleCount, actualUser)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}

func TestCreateKudaGoEvent(t *testing.T) {
	for _, test := range CreateKudaGoEventTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualErr := repositoryTest.CreateKudaGoEvent(test.inputEventID)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})
	}
}

func TestSwitchEventMeeting(t *testing.T) {
	for _, test := range SwitchEventMeetingTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualErr := repositoryTest.SwitchEventMeeting(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})
	}
}

func TestCheckKudaGoMeeting(t *testing.T) {
	for _, test := range checkKudaGoMeetingTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualIsGoing, actualErr := repositoryTest.CheckKudaGoMeeting(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.expectedIsGoing, actualIsGoing)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})
	}
}

func TestCheckKudaGoFavourite(t *testing.T) {
	for _, test := range checkKudaGoFavouriteTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualIsFavourite, actualErr := repositoryTest.CheckKudaGoFavourite(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.expectedIsFavourite, actualIsFavourite)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})
	}
}

func TestSwitchEventFavourite(t *testing.T) {
	for _, test := range SwitchEventFavouriteTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualErr := repositoryTest.SwitchEventFavourite(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})
	}
}