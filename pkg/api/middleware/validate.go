package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func ValidateToken(cookie string) (int, error) {
	//parses, validates, verifies the signature and returns the parsed token
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		//secret was used for signing the string
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	var parsedID interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedID = claims["id"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return 0, fmt.Errorf("token expired")
		}
	}

	//type assertion
	value, ok := parsedID.(float64)
	if !ok {
		return 0, fmt.Errorf("expected an int value, but got %T", parsedID)
	}

	id := int(value)

	return id, nil
}
