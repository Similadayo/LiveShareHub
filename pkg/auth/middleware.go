package auth

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/similadayo/internal/user"
	"github.com/similadayo/pkg/logging"
	"github.com/similadayo/pkg/utils"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateRequest(c *gin.Context, req interface{}) error {
	err := c.ShouldBindJSON(req)
	if err != nil {
		return err
	}

	err = validate.Struct(req)
	if err != nil {
		return err
	}

	return nil
}

// logger middleware
func LoggerMiddleWare(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		logger.Info("request", map[string]interface{}{
			"Method":  c.Request.Method,
			"URI":     c.Request.RequestURI,
			"Status":  c.Writer.Status(),
			"Latency": latency,
		})
	}
}

// ContextMiddleware adds a context to the request with timeout and request-scoped values.
func ContextMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		ctx = context.WithValue(ctx, "request_id", c.Writer.Header().Get("X-Request-ID"))

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// RateLimiterMiddleware applies rate limiting to the API.
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit.NewRateLimiter(func(c *gin.Context) string {
			return c.ClientIP()
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			return rate.NewLimiter(rate.Every(time.Second), 1), time.Hour
		}, func(c *gin.Context) {
			c.AbortWithStatusJSON(429, gin.H{
				"error": "too many requests",
			})
		})
	}
}

// SecurityHeadersMiddleware sets security headers in the response.
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("strict-transport-security", "max-age=31536000; includeSubDomains")
		c.Writer.Header().Set("x-content-type-options", "nosniff")
		c.Writer.Header().Set("x-frame-options", "DENY")
		c.Writer.Header().Set("x-xss-protection", "1; mode=block")
		c.Writer.Header().Set("referrer-policy", "no-referrer")
		c.Writer.Header().Set("content-security-policy", "default-src 'none'; img-src 'self'; script-src 'self'; style-src 'self'; object-src 'none'")
		c.Writer.Header().Set("x-permitted-cross-domain-policies", "none")
		c.Next()
	}
}

// AuthMiddleWare checks if the user is authenticated.
func AuthMiddleWare(userService *user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		userID, err := utils.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		user, err := userService.Repository.GetUserByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
