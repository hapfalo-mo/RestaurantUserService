package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginResponse struct {
	UserId      primitive.ObjectID `json:"_id" bson:"_id, omitempty"`
	Id          int                `json:"id"`
	PhoneNumber string             `json:"phone_number"`
	Email       string             `json:"email"`
	FullName    string             `json:"full_name"`
	Role        string             `json:"role"`
	Point       int                `json:"point"`
}
