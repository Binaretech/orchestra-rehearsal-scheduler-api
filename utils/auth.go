package utils

import (
	"fmt"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, role string) (string, error) {
	var secretKey = config.GetConfig().TokenSecret

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates a JWT token.
func ValidateToken(tokenString string) (*jwt.Token, error) {
	var secretKey = config.GetConfig().TokenSecret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	return token, err

}
