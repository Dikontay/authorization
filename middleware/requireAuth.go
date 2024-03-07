package middleware

import (
	"auth/initializers"
	"auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("Auth")
	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error parsing token"})
		}
		return
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if claims, ok := token.Claims.(jwt.MapClaims); ok {

			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is expired"})
			}
			var user *models.User

			initializers.DB.First(&user, claims["sub"])

			if user.ID == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

}
