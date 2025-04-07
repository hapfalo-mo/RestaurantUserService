package models

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Password    string `json:"Password" bson:"password"`
}
