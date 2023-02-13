package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	TokenAuthMiddleware() gin.HandlerFunc
	CORSMiddleware() gin.HandlerFunc
	MiddlewareValidateUser() gin.HandlerFunc
	MiddlewareValidateLoginUser() gin.HandlerFunc
	MiddlewareValidateUserFormData() gin.HandlerFunc
	MiddlewareValidateRedeemCode() gin.HandlerFunc
}