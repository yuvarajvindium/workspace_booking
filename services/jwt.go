package services

import (
	"errors"
	"workspace_booking/config"

	"github.com/golang-jwt/jwt"
)

type JWTCustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

var mySigningKey = []byte(config.GetJWTSecret())

func GenerateJWT(userId string) (string, error) {
	// Create the Claims
	claims := JWTCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "indium software",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		token,
		&JWTCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify

			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("Invalid token")
			}

			return mySigningKey, nil
		})
}
