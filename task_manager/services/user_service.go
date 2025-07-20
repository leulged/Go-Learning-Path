package services

import (
	"context"
	"errors"

	// "log"
	"strings"
	// "time"

	// "task_manager/data"
	"task_manager/data"
	"task_manager/models"
	"task_manager/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user models.User) (models.User, error)
	Login(email, password string) (string, error)
	PromoteToAdmin(email string) error
	GetUserByEmail(email string) (models.User, error)
}

type userServiceImpl struct {
	collection *mongo.Collection
}

func NewUserService() UserService {
	return &userServiceImpl{collection: data.UserCollection}
}

func (us *userServiceImpl) Register(user models.User) (models.User, error) {
	user.Email = strings.ToLower(user.Email)

	count, _ := us.collection.CountDocuments(context.TODO(), bson.M{})
	role := "user"
	if count == 0 {
		role = "admin" // First user becomes admin
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)
	user.Role = role

	_, err = us.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (us *userServiceImpl) Login(email, password string) (string, error) {
	var user models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.Email, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (us *userServiceImpl) PromoteToAdmin(email string) error {
	update := bson.M{"$set": bson.M{"role": "admin"}}
	res, err := us.collection.UpdateOne(context.TODO(), bson.M{"email": email}, update)
	if err != nil || res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (us *userServiceImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}
