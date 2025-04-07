package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId      primitive.ObjectID `json:"_id" bson:"_id, omitempty"`
	Id          int                `json:"id" bson:"id, omitempty"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Password    string             `json:"password" bsson:"password"`
	Email       string             `json:"email" bson :"email"`
	FullName    string             `json:"full_name" bson:"full_name"`
	Role        string             `json:"role" bson:"role"`
	Point       int                `json:"point" bson:"point"`
	CreatedAt   string             `json:"created_at" bson: "created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeletedAt   *string            `json:"deleted_at" bson:"deleted_at"`
}
