package service

import (
	"context"
	// "errors"
	"log"

	"user_management/config"
	"user_management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	GetUsers() []models.User
	GetUserById(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id int, updateUser models.User) (models.User, error)
	DeleteUser(id int) error
}

type userServiceImpl struct {
	collection *mongo.Collection
}

// Constructor function (exported)
func NewUserService() UserService {
	return &userServiceImpl{
		collection: config.UserCollection,

	}
}

func (usi *userServiceImpl) GetUsers() []models.User {
	var result []models.User

	cursor, err := usi.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error fetching tasks:", err)
		return result
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding User:", err)
			continue
		}
		result = append(result, user)
	}

	return result
}

func (usi *userServiceImpl) CreateUser(user models.User) (models.User , error) {
	_ , err := usi.collection.InsertOne(context.TODO() , user)
	if err != nil {
		return models.User{} , err
	}
	return user , nil
}

func (usi *userServiceImpl) GetUserById (id int) (models.User , error) {
	var user models.User
	err := usi.collection.FindOne(context.TODO() , bson.M{"id" : id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{} , err
		}
		return models.User{} , err
	}
	
	return user , nil
}

func (usi *userServiceImpl) UpdateUser(id int, updateUser models.User) (models.User, error) {
	filter := bson.M{"id" : id}
	update := bson.M{
		"$set" : bson.M{
			"name" : updateUser.Name,
			"age" : updateUser.Age,
			"email" : updateUser.Email,
			"id" : updateUser.ID,
		},
	

	}
	result , err := usi.collection.UpdateOne(context.TODO() , filter, update)
	if err != nil {
		return models.User{} , err
	}
	if result.MatchedCount == 0{
		return models.User{} , err
	}
	updateUser.ID = id
	return updateUser , nil
}

func (usi *userServiceImpl) DeleteUser(id int) error {
	result , err := usi.collection.DeleteOne(context.TODO() , bson.M{"id" : id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return err
	}
	return nil

}