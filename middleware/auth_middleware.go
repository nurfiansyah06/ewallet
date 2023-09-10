package middleware

import (
	"errors"
	"ewalletgolang/db"
	"ewalletgolang/repository"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWT string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = now.Add(ttl).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWT))
	if err != nil {
		return "", fmt.Errorf("failed token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func AuthMiddleware(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func Authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
       authHeader := c.GetHeader("Authorization")

	   if !strings.Contains(authHeader, "Bearer") {
		   c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
		   return
	   }

	   tokenString := ""
	   resultToken := strings.Split(authHeader, " ")
	   if len(resultToken) == 2 {
			tokenString	= resultToken[1]
	   }

	   token, err := ValidateToken(tokenString)
	   if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
	   }

	   claim, ok := token.Claims.(jwt.MapClaims)
	   if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
	   }

	   userId := int(claim["user_id"].(float64))

	   db := db.ConnectDB()

	   userRepo := repository.NewRepository(db)
	   userUsecase := usecase.NewUsecase(userRepo)
	   user, err := usecase.UserUsecase.FindUserById(userUsecase, userId)

	   if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
	   }

	   c.Set("currentUser", user)
	   c.Set("claims", claim)

	   c.Next()
	}
}