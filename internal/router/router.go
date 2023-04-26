package router

import (
	auth "Diploma/internal/microservices/auth"
	eventV2 "Diploma/internal/microservices/event_v2"
	user "Diploma/internal/microservices/user"

	"Diploma/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthEndpoints(r *gin.RouterGroup, mws middleware.Middleware, aD auth.Delivery) {
	r.POST("/signup", mws.MiddlewareValidateUserFormData(), aD.SignUp)
	r.POST("/login", mws.MiddlewareValidateLoginUser(), aD.SignIn)
	r.POST("/logout", mws.TokenAuthMiddleware(), aD.Logout)
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

func EventV2Endpoints(r *gin.RouterGroup, mws middleware.Middleware, eD eventV2.Delivery) {
	r.GET("/external", mws.TokenAuthMiddleware(), eD.GetExternalEvents)
	r.GET("/external/today", mws.TokenAuthMiddleware(), eD.GetTodayEvents)
	r.GET("/external/close", mws.TokenAuthMiddleware(), eD.GetCloseEvents)
	r.GET("/external/:place_id/:event_id", mws.TokenAuthMiddleware(), eD.GetExternalEvent)
	r.GET("/external/similar", mws.TokenAuthMiddleware(), eD.GetSimilar)
	r.GET("/external/alike/:event_id", mws.TokenAuthMiddleware(), eD.GetSimilarToEvent)
	r.GET("/external/alike_title/:event_id", mws.TokenAuthMiddleware(), eD.GetSimilarToEventByTitle)
	r.POST("/external/:event_id/like", mws.TokenAuthMiddleware(), eD.SwitchLikeEvent)
	r.POST("/external/:event_id/dislike", mws.TokenAuthMiddleware(), eD.SwitchLikeEvent)
	r.GET("/external/likes/:user_id", mws.TokenAuthMiddleware(), eD.GetFavourites)
	r.GET("/external/search", mws.TokenAuthMiddleware(), eD.SearchEvents)
}
