package middlewares

import (
	"dbo-technical-test/helpers"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Token Not Found",
			})
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")

		if !bearer {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Bearer Not FOund",
			})
			return
		}
		tokenStr := strings.Split(token, "Bearer ")[1]

		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Token STR",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)

		if err != nil {
			log.Errorln("ERROR:", err)
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		var data = claims.(jwt.MapClaims)
		userId := data["id"].(float64)
		strUserId := strconv.FormatFloat(userId, 'f', -1, 64)

		ctx.Set("id", strUserId)
		ctx.Set("name", data["name"])
		ctx.Set("email", data["email"])
		ctx.Set("role", data["role"])
		ctx.Set("exp", data["exp"])

		if data["exp_date"] == nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
			return
		}

		timeNow := time.Now()
		expiredTime := data["exp_date"].(string)
		parsed, _ := time.Parse(time.RFC3339, expiredTime)

		if err != nil {
			log.Errorln("ERROR:", err)
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		if timeNow.After(parsed) {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":         "loggedOut",
				"message":       "Token has expired, please login again",
				"is_logged_out": true,
			})
			return
		}

		ctx.Next()
	}
}

func IsSuperAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		if !helpers.IsSuperAdmin(role) {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": helpers.ROLE_NOT_ALLOWED,
			})
			return
		}
	}
}
