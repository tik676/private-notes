package authorization

import (
	"private-notes/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

type Maker interface {
	CreateToken(userID int, durarion time.Duration) (string, error)
	VerifyToken(token string) (*models.User, error)
}

func NewJWTMaker(secret string) *JWTMaker {
	return &JWTMaker{secretKey: secret}
}

func (maker *JWTMaker) CreateToken(userID int, durarion time.Duration) (string, error) {
	payload := &models.Payload{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(durarion)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(tokenStr string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.Payload{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, models.ErrTypeOfSignature
		}
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Payload)
	if !ok || !token.Valid {
		return nil, models.ErrInvalidToken
	}

	return &models.User{ID: claims.UserID}, nil
}

/*func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
*/
