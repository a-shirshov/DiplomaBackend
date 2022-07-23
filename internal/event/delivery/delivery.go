package delivery

import (
	"Diploma/internal/event"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventDelivery struct {
	eventUsecase event.Usecase
}

func NewEventDelivery(eventU event.Usecase) (*EventDelivery) {
	return &EventDelivery{
		eventUsecase: eventU,
	}
}

// swagger:route GET /api/events/{page} Events ListEvents
// Returns a list of events of selected page
// responses:
// 200: eventsResponse

// GetEvents returns the events from database by the amount of elementsPerPage
// with selected page
func (eD *EventDelivery) GetEvents(c *gin.Context) {
	pageParam := c.Param("page")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
	}

	resultEvents, err := eD.eventUsecase.GetEvents(page)
	if err != nil {
		c.String(http.StatusBadRequest, "no events for you")
	}
	c.JSON(http.StatusOK,resultEvents)
}