package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/event"
	"Diploma/internal/models"
	"Diploma/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
	pageParam := c.DefaultQuery("page", "1")
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

	event, err := eD.eventUsecase.GetEvent(eventId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, event)
}

// func (eD *EventDelivery) GetTodayEvents(c *gin.Context) {
// 	page, err := utils.GetPageQueryParamFromRequest(c)
// 	if err != nil {
// 		log.Println(err.Error())
// 		utils.SendErrorMessage(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	events, err := eD.eventUsecase.GetTodayEvents(page)
// 	if err != nil {
// 		log.Println(err.Error())
// 		utils.SendErrorMessage(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, events)
// }

func (eD *EventDelivery) GetExternalEvents(c *gin.Context) {
	kudaGoURL := NewKudaGoUrl(mainKudaGoEventURL)
	page := utils.GetPageQueryParamFromRequest(c)
	kudaGoURL.AddPage(page)
	
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	kudaGoEvents := &models.KudaGoEvents{}
	eventErr := make(chan error, 1)
	sendKudagoRequestAndParseToStruct(httpClient, kudaGoURL.url, kudaGoEvents, eventErr)
	if <-eventErr != nil {
		utils.SendErrorMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}
	
	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		events.Events = append(events.Events, toMyEvent(result))
	}
	c.JSON(http.StatusOK, events)
}

func sendKudagoRequestAndParseToStruct(httpClient *http.Client, url string, jsonUnmarshalStruct interface{}, errChan chan<- error) () {
	resp, err := httpClient.Get(url)
	defer close(errChan)
	if err != nil {
		fmt.Println(err)
		errChan <- err
		return 
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(jsonUnmarshalStruct)
	if err != nil {
		fmt.Println(err)
		errChan <- err
		return 
	}
	errChan <- nil
}

func toMyEvent(result models.KudaGoResult) (models.MyEvent) {
	event := models.MyEvent{}
	event.KudaGoID = result.ID
	event.Title = result.Title
	event.Start = result.Dates[0].Start
	event.End = result.Dates[0].End
	event.Image = result.Images[0].Image
	event.Place = result.Place.ID
	event.Location = result.Location.Slug
	return event
}

func (eD *EventDelivery) GetCloseExternalEvents(c *gin.Context) {
	kudaGoURL := NewKudaGoUrl(mainKudaGoEventURL)
	page := utils.GetPageQueryParamFromRequest(c)
	kudaGoURL.AddPage(page)

	longitude := getLongitudeQueryParamFromRequest(c)
	kudaGoURL.AddLongitude(longitude)

	latitude := getLatitudeQueryParamFromRequest(c)
	kudaGoURL.AddLatitude(latitude)

	kudaGoURL.AddRadius()
	
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	kudaGoEvents := &models.KudaGoEvents{}
	eventErr := make(chan error, 1)
	sendKudagoRequestAndParseToStruct(httpClient, kudaGoURL.url, kudaGoEvents, eventErr)
	if <-eventErr != nil {
		utils.SendErrorMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		events.Events = append(events.Events, toMyEvent(result))
	}
	c.JSON(http.StatusOK, events)
}

func getLongitudeQueryParamFromRequest(c *gin.Context) (string) {
	longitudeParam := c.DefaultQuery("lon", "37.6155600")
	return longitudeParam
}

func getLatitudeQueryParamFromRequest(c *gin.Context) (string) {
	latitudeParam := c.DefaultQuery("lat", "55.7522200")
	return latitudeParam
}

func (eD *EventDelivery) GetTodayEvents(c *gin.Context) {
	kudaGoURL := NewKudaGoUrl(mainKudaGoEventURL)
	page := utils.GetPageQueryParamFromRequest(c)
	kudaGoURL.AddPage(page)
	kudaGoURL.AddActualSince()
	kudaGoURL.AddActualUntil()

	var httpClient = &http.Client{Timeout: 10 * time.Second}
	kudaGoEvents := &models.KudaGoEvents{}
	eventErr := make(chan error, 1)
	sendKudagoRequestAndParseToStruct(httpClient, kudaGoURL.url, kudaGoEvents, eventErr)
	if <-eventErr != nil {
		utils.SendErrorMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		events.Events = append(events.Events, toMyEvent(result))
	}
	c.JSON(http.StatusOK, events)
}

func (eD *EventDelivery) GetExternalEvent(c *gin.Context) {
	var userID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}
	KudaGoEventUrl := NewKudaGoUrl(KudaGoEventURL)
	KudaGoPlaceUrl := NewKudaGoUrl(KudaGoPlaceUrl)

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	KudaGoEventUrl.AddEventId(eventIDStr)
	KudaGoEventUrl.AddEventFields()

	placeID := c.Param("place_id")
	KudaGoPlaceUrl.AddPlaceId(placeID)
	KudaGoPlaceUrl.AddPlaceFields()

	var httpClient = &http.Client{Timeout: 10 * time.Second}
	
	KudaGoEvent := &models.KudaGoResult{}
	KudaGoPlace := &models.KudaGoPlaceResult{}

	eventErr := make(chan error, 1)
	placeErr := make(chan error, 1)
	go sendKudagoRequestAndParseToStruct(httpClient, KudaGoEventUrl.url, KudaGoEvent, eventErr)
	go sendKudagoRequestAndParseToStruct(httpClient, KudaGoPlaceUrl.url, KudaGoPlace, placeErr)
	peopleCount, isGoing, err := eD.eventUsecase.GetPeopleCountAndCheckMeeting(userID, eventID)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusInternalServerError, customErrors.ErrPostgres.Error())
		return
	}
	
	if <-eventErr != nil {
		utils.SendErrorMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	if <-placeErr != nil {
		utils.SendErrorMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	event := toMyEvent(*KudaGoEvent)
	eventAndPlace := &models.KudaGoPlaceAndEvent{
		Event: event,
		Place: *KudaGoPlace,
		PeopleCount: peopleCount,
		IsGoing: isGoing,
	}
	c.JSON(http.StatusOK, eventAndPlace)
}

func (eD *EventDelivery) SwitchEventMeeting(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err = eD.eventUsecase.SwitchEventMeeting(au.UserId, eventID)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}