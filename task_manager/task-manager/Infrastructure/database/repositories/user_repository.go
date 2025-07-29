package repositories

import (
	"context"
	"strings"
	"task_manager/Domain/entities"
	"task_manager/Domain/errors"
	"task_manager/Domain/interfaces"
	"task_manager/Infrastructure/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) interfaces.UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) GetUserByEmail(email string) (entities.User, error) {
	filter := bson.M{"email": strings.ToLower(email)}
	var doc models.UserDocument
	
	err := r.collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.User{}, errors.UserNotFoundError{}
		}
		return entities.User{}, err
	}
	
	return models.UserToDomain(doc), nil
}

func (r *userRepository) CountDocuments(email string) (int64, error) {
	filter := bson.M{"email": strings.ToLower(email)}
	count, err := r.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) InsertOne(user entities.User) (entities.User, error) {
	doc, err := models.UserFromDomain(user)
	if err != nil {
		return entities.User{}, err
	}

	result, err := r.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return entities.User{}, err
	}

	doc.ID = result.InsertedID.(primitive.ObjectID)
	return models.UserToDomain(doc), nil
}

func (r *userRepository) UpdateOne(email string, user entities.User) (entities.User, error) {
	doc, err := models.UserFromDomain(user)
	if err != nil {
		return entities.User{}, err
	}

	filter := bson.M{"email": strings.ToLower(email)}
	update := bson.M{"$set": doc}
	
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return entities.User{}, err
	}
	
	return models.UserToDomain(doc), nil
}

func (r *userRepository) UpdateRole(email, role string) error {
	filter := bson.M{"email": strings.ToLower(email)}
	update := bson.M{"$set": bson.M{"role": role}}
	
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.UserPromotionError{Message: "failed to update user role"}
	}
	
	return nil
} 