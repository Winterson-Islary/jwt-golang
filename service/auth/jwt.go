package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Winterson-Islary/jwt-golang.git/config"
	"github.com/Winterson-Islary/jwt-golang.git/types"
	"github.com/Winterson-Islary/jwt-golang.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {

	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		tokenString := GetTokenFromRequest(req)
		token, err := ValidateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(res)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(res)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, _ := strconv.Atoi(str)

		user, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(res)
			return
		}

		// TODO: IMPLEMENT CONTEXT

	}
}

func GetTokenFromRequest(req *http.Request) string {
	tokenAuth := req.Header.Get("Authorization")
	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(res http.ResponseWriter) {
	utils.WriteError(res, http.StatusForbidden, fmt.Errorf("permission denied"))
}
