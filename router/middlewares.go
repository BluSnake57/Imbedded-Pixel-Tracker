package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func default_middleware() middlewareFunc {
	return func(gc *gin.Context) {
		gc.Next()
	}
}

func Client_Middleware() middlewareFunc {
	return func(gc *gin.Context) {
		// println(gc.RemoteIP())
		fmt.Println("request header", gc.Request)
		// println(gc.GetHeader("CF-Connecting-IP"))
		// println(gc.GetHeader("X-Forwarded-For"))
		// println(gc.GetHeader("CF-IPCountry"))
		gc.Next()
	}
}

func Server_Middleware(requiredRole string) middlewareFunc {
	return func(gc *gin.Context) {
		// Assume we get the role from some context or header (example purpose)
		userRole := gc.GetHeader("user")

		if userRole == "" || userRole != requiredRole {
			gc.JSON(http.StatusUnauthorized, gin.H{"error": "Forbidden"})
			gc.Abort()
		} else {
			gc.Next()
		}
	}
}
