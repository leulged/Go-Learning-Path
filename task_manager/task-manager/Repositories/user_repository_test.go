package Repositories

import (
	"testing"

	"task_manager/domain"
	// infrastructure "task_manager/Infrastructure"


	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddUser(t *testing.T) {
	setupTestDB()
	repo := NewUserRepository(userCollection)

	dummyUser := domain.User{
		Email: "test@test.com",
		Password: "password",
		Role: "user",
	}
	user , err := repo.InsertOne(dummyUser)
	assert.NoError(t, err)
	assert.NotEqual(t, primitive.NilObjectID, user.ID)

}

func TestGetUserByEmail(t *testing.T) {
	setupTestDB()
	repo := NewUserRepository(userCollection)

	dummyUser := domain.User{
		Email: "test@test.com",
		Password: "password",
		Role: "user",
	}
	inserted, err := repo.InsertOne(dummyUser)
	assert.NoError(t, err)

	user, err := repo.GetUserByEmail(inserted.Email)
	assert.NoError(t, err)
	assert.Equal(t, inserted.Email, user.Email)
}

func TestCountDocuments(t *testing.T) {
	setupTestDB()
	repo := NewUserRepository(userCollection)

	dummyUser := domain.User{
		Email:    "test@test.com",
		Password: "password",
		Role:     "user",
	}
	inserted, err := repo.InsertOne(dummyUser)
	assert.NoError(t, err)

	count, err := repo.CountDocuments(inserted.Email)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestUpdateOne(t *testing.T) {
	setupTestDB()
	repo := NewUserRepository(userCollection)

	dummyUser := domain.User{
		Email: "test@test.com",
		Password: "password",
		Role: "user",
	}
	inserted, err := repo.InsertOne(dummyUser)
	assert.NoError(t, err)

	updateUser := domain.User{
		Email: "test@test.com",
		Password: "password",
		Role: "admin",
	}
	updated , err := repo.UpdateOne(inserted.Email, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, updateUser.Role, updated.Role)
	assert.Equal(t, updateUser.Email, updated.Email)

}
func TestUpdateRole(t *testing.T) {
	setupTestDB()
	repo := NewUserRepository(userCollection)

	dummyUser := domain.User{
		Email: "test@test.com",
		Password: "password",
		Role: "user",
	}
	inserted, err := repo.InsertOne(dummyUser)
	assert.NoError(t, err)

	err = repo.UpdateRole(inserted.Email, "admin")
	assert.NoError(t, err)

	updated, err := repo.GetUserByEmail(inserted.Email)
	assert.NoError(t, err)
	assert.Equal(t, "admin", updated.Role)

}