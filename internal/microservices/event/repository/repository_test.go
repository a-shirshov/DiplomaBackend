package repository

import (
	"Diploma/internal/models"
	"testing"
	"fmt"

	"database/sql/driver"
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