package user

import "github.com/gin-gonic/gin"

type Delivery interface {
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateUserImage(c *gin.Context)
	GetFavourites(c *gin.Context)
}