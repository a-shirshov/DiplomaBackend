package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	elementsPerPage = 2
)

const (
	GetEventsQuery = `select id, name, description, about, category, tags, specialInfo from (
		select ROW_NUMBER() OVER (ORDER BY creationDate) as RowNum, * from "event" where place_id = $1) as eventsPaged 
		where RowNum Between 1 + $2 * ($3-1) and $2 * $3;`
	GetEventQuery = `select id, name, description, about, category, tags, specialInfo from "event" 
		where id = $1;`
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (eR *EventRepository) GetEvents(page int) ([]*models.Event, error) {
	events := []*models.Event{}
	err := eR.db.Select(&events,GetEventsQuery, elementsPerPage, page)
	if err != nil {
		log.Println(err)
		return events, customErrors.ErrPostgres
	}
	return events, nil
}

func (eR *EventRepository) GetEvent(eventId int) (*models.Event, error) {
	event := &models.Event{}
	err := eR.db.Get(&event, GetEventQuery, eventId)
	if err != nil {
		log.Println(err)
		return event, customErrors.ErrPostgres
	}
	return event, nil
}
