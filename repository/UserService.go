package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/hapfalo-mo/RestaurantUserService/custom"
	"github.com/hapfalo-mo/RestaurantUserService/db"
	errorList "github.com/hapfalo-mo/RestaurantUserService/error"
	"github.com/hapfalo-mo/RestaurantUserService/interfaces"
	models "github.com/hapfalo-mo/RestaurantUserService/models"
	dto "github.com/hapfalo-mo/RestaurantUserService/models/dto"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var _ interfaces.UserInterface = &UserService{}

type UserService struct{}

func (u *UserService) GetAllUser() (result []models.User, err error) {
	collection := db.GetCollectionUser("user")
	log.Println(collection.Name())
	ctx, cancle := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancle()
	cursor, err := collection.Find(ctx, bson.M{})
	log.Println(cursor)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user models.User
		err = cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func (u *UserService) HashPassword(phone, password string) string {
	newPassword := password + phone
	hashPassword := sha256.Sum256([]byte(newPassword))
	hasStrPassword := hex.EncodeToString(hashPassword[:])
	return hasStrPassword
}

//	func (u *UserService) Login(request *dto.LoginRequest) (response dto.LoginResponse, err error) {
//		collection := db.GetCollectionUser("user")
//		ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
//		defer cancle()
//		var user models.User
//		newPassowrd := u.HashPassword(request.PhoneNumber, request.Password)
//		err = collection.FindOne(ctx, bson.M{"phone_number": request.PhoneNumber, "password": newPassowrd}).Decode(&user)
//		if err != nil {
//			return response, err
//		}
//		response = dto.LoginResponse{
//			UserId:      user.UserId,
//			Id:          user.Id,
//			PhoneNumber: user.PhoneNumber,
//			Email:       user.Email,
//			FullName:    user.FullName,
//			Role:        user.Role,
//			Point:       user.Point,
//		}
//		return response, nil
//	}
func (u *UserService) LoginToken(request *dto.LoginRequest) (response custom.Data[dto.LoginResponse], error custom.Error) {
	collection := db.GetCollectionUser("user")
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()
	var user models.User
	newPassowrd := u.HashPassword(request.PhoneNumber, request.Password)
	userdb := collection.FindOne(ctx, bson.M{"phone_number": request.PhoneNumber, "password": newPassowrd}).Decode(&user)
	if userdb != nil {
		return custom.Data[dto.LoginResponse]{}, custom.Error{
			Message:    errorList.InvalidPhoneNumber.Error(),
			ErrorField: userdb.Error(),
			Field:      "credentials",
		}
	}
	userInfo := models.User{
		UserId:      user.UserId,
		Id:          user.Id,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		FullName:    user.FullName,
		Role:        user.Role,
		Point:       user.Point,
		Password:    "",
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}
	token, err := CreateToken(&userInfo)
	if err != nil {
		return custom.Data[dto.LoginResponse]{}, custom.Error{
			Message:    errorList.TokenGenerateError.Error(),
			ErrorField: err.Error(),
			Field:      "token",
		}
	}
	loginResponse := dto.LoginResponse{
		Data:        userInfo,
		TokenString: token,
	}
	response = custom.Data[dto.LoginResponse]{
		Data: loginResponse,
	}
	return response, custom.Error{}
}

func (u *UserService) IsAcceptUserAccess(tokenString string) (response bool, err error) {
	_, err = ParseToken(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserService) GetUserByUserId(id int) (response custom.Data[dto.UserResponse], err custom.Error) {
	userCollection := db.GetCollectionUser("user")
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()
	var user dto.UserResponse
	filter := bson.M{
		"id":         id,
		"deleted_at": "",
	}
	userdb := userCollection.FindOne(ctx, filter).Decode(&user)
	if userdb != nil {
		return custom.Data[dto.UserResponse]{}, custom.Error{
			Message:    errorList.ErrGetUserById.Error(),
			ErrorField: userdb.Error(),
			Field:      "UserService - GetUserByUserId",
		}
	}
	userInfo := dto.UserResponse{
		UserId:      user.UserId,
		Id:          user.Id,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		FullName:    user.FullName,
		Role:        user.Role,
		Point:       user.Point,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}
	return custom.Data[dto.UserResponse]{Data: userInfo}, custom.Error{}
}

// Internal Function
