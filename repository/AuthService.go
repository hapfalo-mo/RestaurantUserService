package repository

import (
	"fmt"
	models "restaurantuserservice/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var secretKey = []byte("a-string-secret-at-least-256-bits-long")

type Claims struct {
	UserID    primitive.ObjectID `json:"userID"`
	Id        int                `json:"id"`
	Role      string             `json:"role"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	FullName  string             `json:"fullName"`
	Point     int                `json:"point"`
	Password  string             `json:"password"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	DeletedAt *string            `json:"deleted_at"`
	jwt.RegisteredClaims
}

func CreateToken(data *models.User) (result string, err error) {
	claimss := &Claims{
		UserID:    data.UserId,
		Id:        data.Id,
		Role:      data.Role,
		Email:     data.Email,
		Phone:     data.PhoneNumber,
		FullName:  data.FullName,
		Point:     data.Point,
		Password:  "",
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "restaurant-user-service",
			Subject:   "Authentication",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claimss).SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(tokenString string) (claims *Claims, err error) {
	claims = &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Token is not valid!")
	}
	return claims, nil
}
