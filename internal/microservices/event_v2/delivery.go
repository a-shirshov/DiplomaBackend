package eventV2

import "github.com/gin-gonic/gin"

type Delivery interface {
	GetExternalEvents(c *gin.Context)
	GetTodayEvents(c *gin.Context)
	GetCloseEvents(c *gin.Context)
	GetExternalEvent(c *gin.Context)
	GetSimilar(c *gin.Context)
	GetSimilarToEvent(c *gin.Context)
	GetSimilarToEventByTitle(c *gin.Context)
	SwitchLikeEvent(c *gin.Context)
	GetFavourites(c *gin.Context)
	SearchEvents(c *gin.Context)
}