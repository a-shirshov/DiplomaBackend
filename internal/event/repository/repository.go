package repository

import (
	"Diploma/internal/models"

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

func (eR *EventRepository) GetEvents(page int) ([]*models.Event, error) {
	rows, err := eR.db.Query("GetEventsQuery", elementsPerPage, page)
	if err != nil {
		rows.Close()
		return nil, err
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
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}