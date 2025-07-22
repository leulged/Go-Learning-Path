package Repositories

import (
	"context"
	// "errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Concrete implementation
type userRepository struct {
	collection *mongo.Collection
}

// Constructor
func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &userRepository{
		collection: collection,
	}
}

// Method implementations

func (r *userRepository) GetUserByEmail(email string) (domain.User, error) {
	filter := bson.M{"email": email}
	var user domain.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func (r *userRepository) CountDocuments(email string) (int64, error) {
	filter := bson.M{"email": email}
	count, err := r.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) InsertOne(user domain.User) (domain.User, error) {
	result, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *userRepository) UpdateOne(email string, user domain.User) (domain.User, error) {
	filter := bson.M{"email": email}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateRole(email, role string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"role": role}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
