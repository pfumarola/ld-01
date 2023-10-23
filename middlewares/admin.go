package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := []byte(os.Getenv("JWTSECRET"))
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if isAdmin := claims["user"].(map[string]interface{})["isAdmin"].(bool); isAdmin {
			c.Next()
			return
		}
	}
	c.JSON(http.StatusForbidden, gin.H{"msg": "You are not allowed to do this"})
	c.Abort()
	return
}
