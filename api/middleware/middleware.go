package middleware

import (
	"errors"
	"net/http"
	"strings"

	auth "gateway/api/token"
	"gateway/redis"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Redis *redis.RedisRepo
}

func NewAuthMiddleware(r *redis.RedisRepo) *AuthMiddleware {
	return &AuthMiddleware{Redis: r}
}

// 🔹 JWT + Redis blacklist check
func (m *AuthMiddleware) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenStr := strings.TrimPrefix(accessToken, "Bearer ")

		// Redis blacklist check
		isBlacklisted, err := m.Redis.IsBlacklisted(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
			return
		}

		if isBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
			return
		}

		// JWT validation
		claims, err := auth.ValidateAccessToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userID, ok1 := claims["user_id"].(string)
		role, ok2 := claims["role"].(string)
		if !ok1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id in token"})
			return
		}
		if !ok2 {
			role = "user" // default role
		}

		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}
}

// 🔹 Casbin permission check
type CasbinMiddleware struct {
	Enforcer *casbin.Enforcer
}

func NewCasbinMiddleware(enf *casbin.Enforcer) *CasbinMiddleware {
	return &CasbinMiddleware{Enforcer: enf}
}

func (c *CasbinMiddleware) GetRole(ctx *gin.Context) (string, int) {
	roleValue, exists := ctx.Get("role")
	if !exists {
		return "unauthorized", http.StatusUnauthorized
	}
	role, ok := roleValue.(string)
	if !ok {
		return "invalid role", http.StatusInternalServerError
	}
	return role, 0
}

func (c *CasbinMiddleware) CheckPermission(ctx *gin.Context) (bool, error) {
	sub, status := c.GetRole(ctx)
	if status != 0 {
		return false, errors.New("failed to get role")
	}

	ok, err := c.Enforcer.Enforce(sub, ctx.FullPath(), ctx.Request.Method)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (c *CasbinMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok, err := c.CheckPermission(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			return
		}
		ctx.Next()
	}
}