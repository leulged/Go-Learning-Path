package models

import (
	"task_manager/Domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDocument represents the MongoDB document structure
type UserDocument struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

// UserFromDomain converts domain User to MongoDB UserDocument
func UserFromDomain(user entities.User) (UserDocument, error) {
	var objectID primitive.ObjectID
	var err error

	if user.ID != "" {
		objectID, err = primitive.ObjectIDFromHex(user.ID)
		if err != nil {
			return UserDocument{}, err
		}
	}

	return UserDocument{
		ID:       objectID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

// UserToDomain converts MongoDB UserDocument to domain User
func UserToDomain(doc UserDocument) entities.User {
	return entities.User{
		ID:       doc.ID.Hex(),
		Name:     doc.Name,
		Email:    doc.Email,
		Password: doc.Password,
		Role:     doc.Role,
	}
} 