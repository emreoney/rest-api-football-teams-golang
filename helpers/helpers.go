package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken() (string, error) {

	var mySigningKey = []byte("secret_key")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = "123456"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
