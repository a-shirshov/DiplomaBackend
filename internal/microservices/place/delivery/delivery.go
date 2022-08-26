package delivery

import (
	"Diploma/internal/microservices/place"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaceDelivery struct {
	placeUsecase place.Usecase
}

func NewPlaceDelivery(placeU place.Usecase) *PlaceDelivery {
	return &PlaceDelivery{
		placeUsecase: placeU,
	}
}

// @Summary PlacesList
// @Tags Places
// @Description GetPlaces by selected page
// @Param page query int false "Page of events"
// @Accept json
// @Produce json
// @Success 200 {object} []models.Place
// @Failure 400 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /places [get]
func(pD *PlaceDelivery) GetPlaces(c *gin.Context) {
	pageParam := c.DefaultQuery("page", "1")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	places, err := pD.placeUsecase.GetPlaces(page)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, places)
}

func(pd *PlaceDelivery) GetPlace(c *gin.Context) {
	idStr := c.Param("place_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	place, err := pd.placeUsecase.GetPlace(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, place)
}