package middleware

import (
	"e-commerce-api/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthenEmployee() gin.HandlerFunc {
	// custom middleware
	return func(c *gin.Context) {
		// check token
		hmacSampleSecret := []byte(utils.ViperGetString("jwt.secretKeyEmployee"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Please authorized before.",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"])
			c.Set("role_id", claims["role_id"])
			c.Set("shop_id", claims["shop_id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": err.Error(),
			})
		}

		// before request
		c.Next()

	}
}

func JWTAuthenCustomer() gin.HandlerFunc {
	// custom middleware
	return func(c *gin.Context) {
		// check token
		hmacSampleSecret := []byte(utils.ViperGetString("jwt.secretKeyCustomer"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Please authorized before.",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": err.Error(),
			})
		}

		// before request
		c.Next()

	}
}
