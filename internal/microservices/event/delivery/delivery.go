package delivery

import (
	"Diploma/internal/microservices/event"
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

// @Summary EventsList
// @Tags Events
// @Description GetEvents by selected page
// @Param page query int false "Page of events"
// @Param id path int true "Place id"
// @Accept json
// @Produce json
// @Success 200 {object} []models.Event
// @Failure 400 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /places/{id}/events [get]
func (eD *EventDelivery) GetEvents(c *gin.Context) {
	idParam := c.Param("place_id")
	placeId, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
	}
	pageParam := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
	}

	resultEvents, err := eD.eventUsecase.GetEvents(placeId, page)
	if err != nil {
		c.String(http.StatusBadRequest, "no events for you")
	}
	c.JSON(http.StatusOK,resultEvents)
}

// @Summary One Event
// @Tags Events
// @Description Get Event by id
// @Param id path int true "Event id"
// @Accept json
// @Produce json
// @Success 200 {object} models.Event
// @Failure 400 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /events/{id} [get]
func (eD *EventDelivery) GetEvent(c *gin.Context) {
	eventIdString := c.Param("event_id")
	eventId, err := strconv.Atoi(eventIdString)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resultEvent, err := eD.eventUsecase.GetEvent(eventId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultEvent)
}