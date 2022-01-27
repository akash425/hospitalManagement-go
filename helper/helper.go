package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string
	jwt.StandardClaims
}

var jwtKey = []byte("akash")

func WriteResponse(w http.ResponseWriter, r interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(r)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Loginvalid(email string) string {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claims.Email)
	return tokenString

}

func Validation(r string) bool {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(r, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil && token.Valid {
		fmt.Println("valid token")
		return true
	} else {
		fmt.Println("invalid token")
		return false
	}
}
