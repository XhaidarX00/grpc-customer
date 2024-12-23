package middleware

import (
	"microservice/database"
	"microservice/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Middleware struct {
	log    *zap.Logger
	Cacher database.Cacher
}

func NewMiddleware(log *zap.Logger, cacher database.Cacher) Middleware {
	return Middleware{
		log:    log,
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		const authURL = "http://localhost:8081/authentication"

		client := resty.New().
			SetTimeout(10 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(2 * time.Second).
			SetRetryMaxWaitTime(10 * time.Second)

		// Lakukan request
		resp, err := client.R().Post(authURL)
		if err != nil {
			helper.ResponseError(c, "Failed to communicate with authentication service", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Validasi status HTTP response
		if resp == nil || resp.StatusCode() != http.StatusOK {
			helper.ResponseError(c, "Token is required", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// token := c.GetHeader("Authorization")
		// if token == "" {
		// 	helper.ResponseError(c, "Token is required", "Unauthorized", http.StatusUnauthorized)
		// 	c.Abort()
		// 	return
		// }

		// userID := c.GetHeader("User-ID")
		// if userID == "" {
		// 	helper.ResponseError(c, "User-ID is required", "Unauthorized", http.StatusUnauthorized)
		// 	c.Abort()
		// 	return
		// }

		// m.log.Info("Authenticating user", zap.String("userID", userID), zap.String("token", token))

		// storedToken, err := m.Cacher.Get(userID)
		// if err != nil {
		// 	helper.ResponseError(c, "Failed to retrieve token", "Server error", http.StatusInternalServerError)
		// 	c.Abort()
		// 	return
		// }
		// m.log.Info("Authenticating user", zap.String("storedToken", storedToken), zap.String("token", token))

		// if storedToken == "" || storedToken != token {
		// 	helper.ResponseError(c, "Invalid token", "Unauthorized", http.StatusUnauthorized)
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}

func (m *Middleware) RoleAuthorization(requiredRoles string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("User-Role")
		if role == "" {
			helper.ResponseError(c, "Roles are required", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		m.log.Info("Authorizing user role", zap.String("role", role), zap.String("requiredRoles", requiredRoles))

		if role == requiredRoles {
			c.Next()
			return
		}

		helper.ResponseError(c, "Insufficient permissions", "Unauthorized", http.StatusUnauthorized)
		c.Abort()
	}
}
