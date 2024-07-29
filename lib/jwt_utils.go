package lib

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stoewer/go-strcase"
)

func GetJwtTokenForTokenData(clientName string, technologies []string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"exp": jwt.NewNumericDate(time.Now().Add((24 * time.Hour) * 60)),
		"c":   strcase.KebabCase(clientName),
		"n":   clientName,
		"t":   technologies,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
