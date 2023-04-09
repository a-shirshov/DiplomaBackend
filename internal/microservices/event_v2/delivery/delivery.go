package delivery

import (
	"Diploma/internal/microservices/event_v2"
	"Diploma/internal/models"
	"Diploma/utils"
	"log"
	"net/http"
	"strconv"

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

	pageStr := utils.GetPageQueryParamFromRequest(c)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.GetExternalEvents(userID, page)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &models.MyEvents{
		Events: *events,
	})
}

func (eD *EventDeliveryV2) GetTodayEvents(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	pageStr := utils.GetPageQueryParamFromRequest(c)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.GetTodayEvents(userID, page)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &models.MyEvents{
		Events: *events,
	})
}

func (eD *EventDeliveryV2) GetCloseEvents(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	pageStr := utils.GetPageQueryParamFromRequest(c)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	lat := getLatitudeQueryParamFromRequest(c)
	lon := getLongitudeQueryParamFromRequest(c)

	events, err := eD.eventUsecaseV2.GetCloseEvents(lat, lon, userID, page)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &models.MyEvents{
		Events: *events,
	})
}

func getLatitudeQueryParamFromRequest(c *gin.Context) string {
	latitudeParam := c.DefaultQuery("lat", "55.7522200")
	return latitudeParam
}

func getLongitudeQueryParamFromRequest(c *gin.Context) string {
	longitudeParam := c.DefaultQuery("lon", "37.6155600")
	return longitudeParam
}

func(eD *EventDeliveryV2) GetExternalEvent(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	resultEvent, err := eD.eventUsecaseV2.GetExternalEvent(userID, eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resultEvent)
}

func(eD *EventDeliveryV2) GetSimilar(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	searchingEvent := c.Query("q")

	eventVector, err := eD.eventUsecaseV2.GetNLPVector(searchingEvent)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.GetRandomEvents(userID)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	resultEvents := []models.MyEvent{}
	count := 0
	for _, event := range *events {
		cosineSimilarity, err := utils.Cosine(eventVector, event.VectorTitle)
		if err != nil {
			log.Println("Cosine: ", err)
			continue
		}
		
		if cosineSimilarity >= 0.7 {
			log.Println(cosineSimilarity)
			resultEvents = append(resultEvents, event)
			count++
		}

		if count == 10 {
			break
		}
	}
	c.JSON(http.StatusOK, &models.MyEvents{
		Events: resultEvents,
	})
}

func(eD *EventDeliveryV2) GetSimilarToEvent(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	eventVector, err := eD.eventUsecaseV2.GetVector(eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.GetRandomEvents(userID)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	resultEvents := []models.MyEvent{}
	count := 0
	for _, event := range *events {
		cosineSimilarity, err := utils.Cosine(*eventVector, event.Vector)
		if err != nil {
			log.Println("Cosine: ", err)
			continue
		}

		if event.KudaGoID == eventID{
			continue
		}
		
		if cosineSimilarity >= 0.7 {
			log.Println(cosineSimilarity)
			resultEvents = append(resultEvents, event)
			count++
		}

		if count == 10 {
			break
		}
	}
	c.JSON(http.StatusOK, &models.MyEvents{
		Events: resultEvents,
	})
}

func(eD *EventDeliveryV2) GetSimilarToEventByTitle(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	eventVector, err := eD.eventUsecaseV2.GetVectorTitle(eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.GetRandomEvents(userID)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	resultEvents := []models.MyEvent{}
	count := 0
	for _, event := range *events {
		cosineSimilarity, err := utils.Cosine(*eventVector, event.VectorTitle)
		if err != nil {
			log.Println("Cosine: ", err)
			continue
		}

		if event.KudaGoID == eventID{
			continue
		}
		
		if cosineSimilarity >= 0.7 {
			log.Println(cosineSimilarity)
			resultEvents = append(resultEvents, event)
			count++
		}

		if count == 10 {
			break
		}
	}
	
	c.JSON(http.StatusOK, &models.MyEvents{
		Events: resultEvents,
	})
}

func(eD *EventDeliveryV2) SwitchLikeEvent(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	event, err := eD.eventUsecaseV2.SwitchLikeEvent(au.UserId, eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, event)
}

func(eD *EventDeliveryV2) GetFavourites(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	checkedUserIDStr := c.Param("user_id")
	checkedUserID, err := strconv.Atoi(checkedUserIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	pageStr := utils.GetPageQueryParamFromRequest(c)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.GetFavourites(userID, checkedUserID, page)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &models.MyEvents{
		Events: *events,
	})
}

func (eD *EventDeliveryV2) SearchEvents(c *gin.Context) {
	userID := 0
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	searchingEvent := c.Query("q")

	pageStr := utils.GetPageQueryParamFromRequest(c)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eD.eventUsecaseV2.SearchEvents(userID, searchingEvent, page)
	if err != nil {
		log.Println(err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &models.MyEvents{
		Events: *events,
	})
}