package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// User is the core domain entity
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	Role     string             `bson:"role" json:"role"`
}

// UserRepository interface defines user data access operations
type UserRepository interface {
	GetUserByEmail(email string) (User, error)
	CountDocuments(email string) (int64, error)
	InsertOne(user User) (User, error)
	UpdateOne(email string, user User) (User, error)
	UpdateRole(email, role string) error
}
