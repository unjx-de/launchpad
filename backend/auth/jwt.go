package auth

import (
	"dashboard/message"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func createJWT(expiryDuration time.Duration) string {
	t := time.Now()
	claims := &jwt.RegisteredClaims{
		IssuedAt:  &jwt.NumericDate{Time: t},
		ExpiresAt: &jwt.NumericDate{Time: t.Add(expiryDuration)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(Config.Auth.Secret))
	return tokenString
}

func parseJWT() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v\n", token.Header["alg"])
		}
		return []byte(Config.Auth.Secret), nil
	}
}

func validateJWT(tokenString string) error {
	_, err := jwt.Parse(tokenString, parseJWT())
	if err != nil {
		logrus.WithField("error", err).Error(message.CannotProcess.String())
		return err
	} else {
		return nil
	}
}
