package middleware

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/logger"
	"github.com/tonyalaribe/monitor-server/messages"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// validate the token
		token, err := jwt.Parse(c.Request.Header.Get("X-AUTH-TOKEN"), func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens, we also only use its public counter part to verify

			publicKey, err := jwt.ParseRSAPublicKeyFromPEM(config.Get().Encryption.Public)

			if err != nil {
				return publicKey, err
			}
			return publicKey, nil
		})

		// branch out into the possible error from signing
		switch err.(type) {

		case nil: // no error

			if !token.Valid { // but may still be invalid
				logger.Error(c, err)
				c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
				c.Abort()
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("User", claims["User"])
				c.Set("UserName", claims["UserName"])
			} else {
				logger.Error(c, err)
				c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
				c.Abort()
			}

			c.Next()

		case *jwt.ValidationError: // something was wrong during the validation
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logger.Error(c, err)
				c.JSON(http.StatusForbidden, messages.ErrBadToken)
				c.Abort()

			default:
				logger.Error(c, err)
				c.JSON(http.StatusForbidden, messages.ErrBadToken)
				c.Abort()
			}

		default: // something else went wrong
			logger.Error(c, err)
			c.JSON(http.StatusForbidden, messages.ErrBadToken)
			c.Abort()
		}
	}
}
