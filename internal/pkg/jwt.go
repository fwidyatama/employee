package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtClaim struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp int64  `json:"time"`
	jwt.RegisteredClaims
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func GenerateJWT(payload map[string]string, jwtKey string) (string, error) {
	jwtKeyByte := []byte(jwtKey)

	claims := jwtClaim{
		payload["name"],
		payload["phone"],
		payload["role"],
		time.Now().Unix(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeyByte)
	return tokenString, err
}
