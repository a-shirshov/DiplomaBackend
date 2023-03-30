package delivery

import (
	"Diploma/internal/microservices/event_v2"
	"Diploma/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventDeliveryV2 struct {
	eventUsecaseV2 eventV2.Usecase
}

func NewEventDeliveryV2(eventV2U eventV2.Usecase) *EventDeliveryV2 {
	return &EventDeliveryV2{
		eventUsecaseV2: eventV2U,
	}
}

func (eD *EventDeliveryV2) GetExternalEvents(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	events, err := eD.eventUsecaseV2.GetExternalEvents(userID)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, events)
}