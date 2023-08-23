package middleware

import (
	"ewalletgolang/repository"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

func ValidateToken(token string, signedJWT string) (interface{}, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return []byte(signedJWT), nil
	})

	fmt.Println(err.Error())

	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}

	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok || !tokenString.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}

func DeserializeUser(userRepository repository.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		tokenSecret := os.Getenv("TOKEN_SECRET")
		sub, err := ValidateToken(token, tokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err := strconv.Atoi(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		
		result, err := userRepository.FindUserById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Name)
		ctx.Next()

	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// In a real-world scenario, replace with a secure secret
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})

		fmt.Println("Token String:", tokenString) 

		if err != nil || !token.Valid {
			fmt.Println("Error verifying token:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Error extracting claims from token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			fmt.Println("Error extracting email from claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		fmt.Println("Claims:", claims)
		c.Set("email", email)
		c.Next()
	}
}


