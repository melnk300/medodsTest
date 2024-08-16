package tokens

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type JwtToken struct {
	Value      string
	Expiration time.Time
}

type Claims struct {
	GUID string `json:"guid"`
	IP   string `json:"ip"`
	JTI  string `json:"jti"`
	Exp  int64  `json:"exp"`
	Iat  int64  `json:"iat"`
	jwt.StandardClaims
}

func GenerateTokens(ip string, guid string) (JwtToken, JwtToken, string) {
	acExpirationTime := time.Now().Add(5 * time.Minute)
	rfExpirationTime := time.Now().Add(10 * 24 * time.Hour)

	jti := uuid.New().String()

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"ip":   ip,
		"iat":  time.Now().Unix(),
		"exp":  acExpirationTime.Unix(),
		"jti":  jti,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("Error signing JWT: %v", err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"ip":   ip,
		"iat":  time.Now().Unix(),
		"exp":  rfExpirationTime.Unix(),
		"jti":  jti,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatalf("Error signing JWT: %v", err)
	}

	return JwtToken{accessToken, acExpirationTime}, JwtToken{refreshToken, rfExpirationTime}, jti
}

func ParseToken(clientToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(clientToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Fatalf("Parsing error: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("failed to cast claims to struct")
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
