package delivery

import (
	"Diploma/internal/microservices/event"
	"Diploma/internal/models"
	"Diploma/utils"
	"fmt"
	"net/http"
	"strconv"
	"sync"
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
	sendKudagoRequestAndParseToStruct(httpClient, kudaGoURL.url, kudaGoEvents)

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		events.Events = append(events.Events, toMyEvent(result))
	}
	c.JSON(http.StatusOK, events)
}

func sendKudagoRequestAndParseToStruct(httpClient *http.Client, url string, jsonUnmarshalStruct interface{}) () {
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(jsonUnmarshalStruct)
	if err != nil {
		fmt.Println(err)
		return 
	}
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
	sendKudagoRequestAndParseToStruct(httpClient, kudaGoURL.url, kudaGoEvents)

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
	sendKudagoRequestAndParseToStruct(httpClient, kudaGoURL.url, kudaGoEvents)

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		events.Events = append(events.Events, toMyEvent(result))
	}
	c.JSON(http.StatusOK, events)
}

func (eD *EventDelivery) GetExternalEvent(c *gin.Context) {
	KudaGoEventUrl := NewKudaGoUrl(KudaGoEventURL)
	KudaGoPlaceUrl := NewKudaGoUrl(KudaGoPlaceUrl)

	eventID := c.Param("event_id")
	KudaGoEventUrl.AddEventId(eventID)
	KudaGoEventUrl.AddEventFields()

	placeID := c.Param("place_id")
	KudaGoPlaceUrl.AddPlaceId(placeID)
	KudaGoPlaceUrl.AddPlaceFields()

	var wg sync.WaitGroup
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	
	KudaGoEvent := &models.KudaGoResult{}
	KudaGoPlace := &models.KudaGoPlaceResult{}

	wg.Add(2)
	go sendKudagoRequestAndParseToStruct(httpClient, KudaGoEventUrl.url, KudaGoEvent)
	sendKudagoRequestAndParseToStruct(httpClient,  KudaGoPlaceUrl.url, KudaGoPlace)
	wg.Done()
	wg.Wait()

	event := toMyEvent(*KudaGoEvent)
	eventAndPlace := &models.KudaGoPlaceAndEvent{
		Event: event,
		Place: *KudaGoPlace,
	}
	c.JSON(http.StatusOK, eventAndPlace)
}
