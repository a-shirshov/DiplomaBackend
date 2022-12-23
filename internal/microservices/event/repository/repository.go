package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"Diploma/utils/query"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	elementsPerPage = 2
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (eR *EventRepository) GetEvents(placeId, page int) ([]*models.Event, error) {
	rows, err := eR.db.Query(query.GetEventsQuery, placeId, elementsPerPage, page)
	if err != nil {
		log.Print(err)
		rows.Close()
		return nil, customErrors.ErrPostgres
	}
	var events []*models.Event

	defer rows.Close()

	for rows.Next() {
		event := &models.Event{}
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.About,
			&event.Category,
			&event.Tags,
			&event.SpecialInfo,
		)
		if err != nil {
			return nil, customErrors.ErrPostgres
		}

		events = append(events, event)
	}
	return events, nil
}

func (eR *EventRepository) GetEvent(eventId int) (*models.Event, error) {
	event := &models.Event{}
	err := eR.db.QueryRow(query.GetEventQuery, eventId).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.About,
		&event.Category,
		&event.Tags,
		&event.SpecialInfo,
	)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}

	return event, nil
}
