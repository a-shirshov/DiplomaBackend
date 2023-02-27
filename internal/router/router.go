package router

import (
	auth "Diploma/internal/microservices/auth"
	event "Diploma/internal/microservices/event/delivery"
	user "Diploma/internal/microservices/user"

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
	r.DELETE("/remove", mws.TokenAuthMiddleware(), aD.DeleteUser)
}

func UserEndpoints(r *gin.RouterGroup, mws middleware.Middleware, uD user.Delivery) {
	r.POST("/:user_id", mws.TokenAuthMiddleware(), mws.MiddlewareValidateUser(), uD.UpdateUser)
	r.GET("/:user_id", uD.GetUser)
	r.POST("/:user_id/image", uD.UpdateUserImage)
}

func EventEndpoints(r *gin.RouterGroup, mws middleware.Middleware, eD *event.EventDelivery) {
	r.GET("/external", mws.TokenAuthMiddleware(), eD.GetExternalEvents)
	r.GET("/external/close", mws.TokenAuthMiddleware(), eD.GetCloseExternalEvents)
	r.GET("/external/today", mws.TokenAuthMiddleware(), eD.GetTodayEvents)
	r.GET("/external/:place_id/:event_id", mws.TokenAuthMiddleware(), eD.GetExternalEvent)
	r.POST("/external/:event_id/go", mws.TokenAuthMiddleware(), eD.SwitchEventMeeting)
	r.POST("/external/:event_id/favourite", mws.TokenAuthMiddleware(), eD.SwitchEventFavourite)
	r.GET("/external/likes/:user_id", mws.TokenAuthMiddleware(), eD.GetFavourites)
	//r.POST("", mws.TokenAuthMiddleware(), mws.MiddlewareValidateUserEvent(), eD.CreateUserEvent)
}
