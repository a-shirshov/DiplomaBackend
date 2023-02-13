package auth

import "github.com/gin-gonic/gin"

type Delivery interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
	SendEmail(c *gin.Context)
	CheckRedeemCode(c *gin.Context)
	UpdatePassword(c *gin.Context)
}
