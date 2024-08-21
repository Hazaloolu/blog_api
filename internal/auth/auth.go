package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func generateSecureKey() string {
	secretBytes := make([]byte, 32) // 32 bytes for a 256-bit key
	_, err := rand.Read(secretBytes)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(secretBytes)
}

var jwtKey = []byte(generateSecureKey())

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
	jwt.RegisteredClaims
}

// hash password

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// check password hash

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

// Generate jwt with specified username

func GenerateJwt(username string, userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("Generated Token:", token)
	return token.SignedString(jwtKey)
}

func ValidateJwt(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil

	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil

}
