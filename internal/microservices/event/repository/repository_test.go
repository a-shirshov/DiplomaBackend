package repository

import (
	"Diploma/internal/models"
	"database/sql/driver"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type customConverter struct{}

func (s customConverter) ConvertValue(v interface{}) (driver.Value, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case []string:
		return v, nil
	case int:
		return v, nil
	default:
		return nil, fmt.Errorf("cannot convert %T with value %v", v, v)
	}
}

type getEventsTest struct {
	page 			int
	outputEvents 	[]*models.Event
	outputErr    	error
}

var getEventsByPlaceTests = []getEventsTest{
	{
		1, []*models.Event{
			{
				ID: 1,
				Name: "Event_1",
				Description: "Description_1",
				About: "About_1",
				Category: "Category_1",
				Tags: []string{"tag_1","tag_2"},
				SpecialInfo: "SpecialInfo_1",
			},
			{
				ID: 2,
				Name: "Event_2",
				Description: "Description_2",
				About: "About_2",
				Category: "Category_2",
				Tags: []string{"tag_1","tag_2"},
				SpecialInfo: "SpecialInfo_2",
			},
		}, nil,
	},
}

type getPeopleCountTest struct {
	name string
	inputEventID int
	beforeTest func(mockSQL sqlmock.Sqlmock)
	expextedPeopleCount int
	expectedErr error
}

var getPeopleCountTests = []getPeopleCountTest {
	{
		"Successfully get people count",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"people_count"}
			rows := mockSQL.NewRows(columns).
				AddRow(10)
			
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetPeopleCount)).
				WithArgs(1).
				WillReturnRows(rows)
		},
		10,
		nil,
	},
}

type createKudaGoEventTest struct {
	name string
	inputEventID int
	beforeTest func(mockSQL sqlmock.Sqlmock)
	expectedErr error
}

var createKudaGoEventTests = []createKudaGoEventTest {
	{
		"Successfully create KudaGo Event",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectExec(regexp.QuoteMeta(CreateKudaGoEvent)).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1,1)).
				WillReturnError(nil)
		},
		nil,
	},
}

func TestGetEventsByPlace(t *testing.T) {
	db, mock, err := sqlmock.New(
		sqlmock.ValueConverterOption(customConverter{}), 
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db,"sqlmock")
	repositoryTest := NewEventRepository(sqlxDB)

	columns := []string{"id", "name", "description", "about", "category", "tags", "specialinfo"}

	for _, test := range getEventsByPlaceTests {
		rows := mock.NewRows(columns).
			AddRow(
				test.outputEvents[0].ID,
				test.outputEvents[0].Name,
				test.outputEvents[0].Description,
				test.outputEvents[0].About,
				test.outputEvents[0].Category,
				test.outputEvents[0].Tags,
				test.outputEvents[0].SpecialInfo).
			AddRow(
				test.outputEvents[1].ID,
				test.outputEvents[1].Name,
				test.outputEvents[1].Description,
				test.outputEvents[1].About,
				test.outputEvents[1].Category,
				test.outputEvents[1].Tags,
				test.outputEvents[1].SpecialInfo)
		
		mock.ExpectQuery(GetEventsQuery).
			WithArgs(elementsPerPage, test.page).
			RowsWillBeClosed().WillReturnRows(rows)

		out, dbErr := repositoryTest.GetEvents(test.page)
		assert.Equal(t, test.outputEvents, out)
		assert.Nil(t, dbErr)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}
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
	for _, test := range getPeopleCountTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualPeopleCount, actualErr := repositoryTest.GetPeopleCount(test.inputEventID)
			assert.Equal(t, test.expextedPeopleCount, actualPeopleCount)
			assert.Equal(t, test.expectedErr, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}

func TestCreateKudaGoEvent(t *testing.T) {
	for _, test := range createKudaGoEventTests {
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
			assert.Equal(t, test.expectedErr, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}
