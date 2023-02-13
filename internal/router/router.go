package router

import (
	event "Diploma/internal/microservices/event/delivery"
	user "Diploma/internal/microservices/user"
	auth "Diploma/internal/microservices/auth"

	"Diploma/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AuthEndpoints(r *gin.RouterGroup, mws middleware.Middleware, aD auth.Delivery) {
	r.POST("/signup", mws.MiddlewareValidateUserFormData(), aD.SignUp)
	r.POST("/login", mws.MiddlewareValidateLoginUser(), aD.SignIn)
	r.GET("/logout", mws.TokenAuthMiddleware(), aD.Logout)
	r.POST("/refresh", aD.Refresh)
	r.POST("/redeem", mws.MiddlewareValidateRedeemCode(), aD.SendEmail)
	r.POST("/codecheck", mws.MiddlewareValidateRedeemCode(), aD.CheckRedeemCode)
	r.POST("/credentials", mws.MiddlewareValidateRedeemCode(), aD.UpdatePassword)
}

func UserEndpoints(r *gin.RouterGroup, mws middleware.Middleware, uD user.Delivery) {
	r.POST("/:user_id", mws.TokenAuthMiddleware(), mws.MiddlewareValidateUser(), uD.UpdateUser)
	r.GET("/:user_id", uD.GetUser)
	r.GET("/:user_id/favourites", uD.GetFavourites)
	r.POST("/:user_id/image", uD.UpdateUserImage)
}

func EventEndpoints(r *gin.RouterGroup, mws middleware.Middleware, eD *event.EventDelivery) {
	r.GET("/external", eD.GetExternalEvents)
	r.GET("/external/close", eD.GetCloseExternalEvents)
	r.GET("/external/today", eD.GetTodayEvents)
	r.GET("/external/:place_id/:event_id", mws.TokenAuthMiddleware(), eD.GetExternalEvent)
	r.POST("external/:event_id/go", mws.TokenAuthMiddleware(), eD.SwitchEventMeeting)
	r.POST("external/:event_id/favourite", mws.TokenAuthMiddleware(), eD.SwitchEventFavourite)
}