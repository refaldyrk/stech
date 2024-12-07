package middleware

import (
	"aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
	"kreditplus-test/config"
	"kreditplus-test/helper"
	"kreditplus-test/repository"
	"net/http"
	"strings"
	"time"
)

func PasetoMiddleware(config *config.Config, custRepsoitory *repository.CustomerRepository, key paseto.V4SymmetricKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		authorizationHeader := c.Request.Header.Get("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnauthorized, "unauthorized", gin.H{}))
			c.Abort()
			return
		}

		trackingReq := make(map[string]any)
		trackingReq["timestamp"] = startTime
		trackingReq["method"] = c.Request.Method
		trackingReq["path"] = c.Request.URL.Path
		trackingReq["user-agent"] = c.Request.Header.Get("User-Agent")
		trackingReq["client-ip"] = c.ClientIP()
		trackingReq["query-params"] = c.Request.URL.RawQuery
		trackingReq["response-time"] = time.Since(startTime).String()

		tokenStringHeader := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		claims, err := helper.ValidatePaseto(key, tokenStringHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnauthorized, err.Error(), gin.H{}))
			c.Abort()
			return
		}

		userID := claims["identity"].(string)
		user, err := custRepsoitory.Find("id", userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnauthorized, "user not found", gin.H{}))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
