package repository

import (
	"Diploma/internal/errors"
	"Diploma/internal/models"
	"log"

	"github.com/jackc/pgx"
)

const (
	elementsPerPage = 2
)

type EventRepository struct {
	db *pgx.ConnPool
}

func NewEventRepository(db *pgx.ConnPool) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (eR *EventRepository) GetEvents(placeId, page int) ([]*models.Event, error) {
	rows, err := eR.db.Query("GetEventsQuery", placeId, elementsPerPage, page)
	if err != nil {
		rows.Close()
		return nil, errors.ErrPostgres
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
			return nil, errors.ErrPostgres
		}

		events = append(events, event)
	}
	return events, nil
}

func (eR *EventRepository) GetEvent(eventId int) (*models.Event, error) {
	event := &models.Event{}
	err := eR.db.QueryRow("GetEventQuery", eventId).Scan(
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
		return nil, errors.ErrPostgres
	}

	return event, nil
}