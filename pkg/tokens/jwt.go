package tokens

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type Token struct {
	Value      string
	Expiration time.Time
}

func GenerateTokens(ip string, guid string) (Token, Token, string) {
	acExpirationTime := time.Now().Add(5 * time.Minute)
	rfExpirationTime := time.Now().Add(10 * 24 * time.Hour)

	jti := uuid.New().String()

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"ip":   ip,
		"iat":  time.Now(),
		"exp":  acExpirationTime,
		"jti":  jti,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("Error signing JWT: %v", err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"ip":   ip,
		"iat":  time.Now(),
		"exp":  rfExpirationTime,
		"jti":  jti,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("Error signing JWT: %v", err)
	}

	return Token{accessToken, acExpirationTime}, Token{refreshToken, rfExpirationTime}, jti
}
