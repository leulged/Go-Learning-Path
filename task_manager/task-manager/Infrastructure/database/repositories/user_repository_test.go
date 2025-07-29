package repositories

import (
	"context"
	"path/filepath"
	"testing"
	"task_manager/Domain/entities"
	"task_manager/Infrastructure/database/models"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) (*mongo.Collection, func()) {
	// Load environment variables
	err := godotenv.Load(filepath.Join("..", "..", "..", ".env"))
	if err != nil {
		// Try current directory as fallback
		err = godotenv.Load()
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	// Connect to test database using MongoDB Atlas
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://leulgedion224:YtxgbwYFwW9snTti@cluster0.gdmxw28.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	require.NoError(t, err)

	// Use a test database
	db := client.Database("test_task_manager")
	collection := db.Collection("users")

	// Clean up before test
	_, err = collection.DeleteMany(context.Background(), bson.M{})
	require.NoError(t, err)

	// Return cleanup function
	cleanup := func() {
		collection.DeleteMany(context.Background(), bson.M{})
		client.Disconnect(context.Background())
	}

	return collection, cleanup
}

func TestUserRepository_InsertOne(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(collection)

	// Test user
	user := entities.NewUser("Test User", "test@example.com", "password123")
	user.SetRole("user")

	// Test InsertOne
	insertedUser, err := repo.InsertOne(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, insertedUser.ID)

	// Verify user was saved
	var doc models.UserDocument
	err = collection.FindOne(context.Background(), bson.M{"email": "test@example.com"}).Decode(&doc)
	assert.NoError(t, err)
	assert.Equal(t, "Test User", doc.Name)
	assert.Equal(t, "test@example.com", doc.Email)
	assert.Equal(t, "user", doc.Role)
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(collection)

	// Create test user
	user := entities.NewUser("Test User", "test@example.com", "password123")
	user.SetRole("user")
	_, err := repo.InsertOne(user)
	require.NoError(t, err)

	// Test GetUserByEmail
	foundUser, err := repo.GetUserByEmail("test@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, "Test User", foundUser.Name)
	assert.Equal(t, "test@example.com", foundUser.Email)
	assert.Equal(t, "user", foundUser.Role)

	// Test GetUserByEmail with non-existent email
	notFoundUser, err := repo.GetUserByEmail("nonexistent@example.com")
	assert.Error(t, err)
	assert.Equal(t, entities.User{}, notFoundUser)
}

func TestUserRepository_UpdateOne(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(collection)

	// Create test user
	user := entities.NewUser("Test User", "test@example.com", "password123")
	user.SetRole("user")
	_, err := repo.InsertOne(user)
	require.NoError(t, err)

	// Update user
	updatedUser := entities.NewUser("Updated User", "test@example.com", "password123")
	updatedUser.SetRole("admin")
	updatedUser.ID = user.ID
	
	resultUser, err := repo.UpdateOne("test@example.com", updatedUser)
	assert.NoError(t, err)

	// Verify update
	foundUser, err := repo.GetUserByEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", foundUser.Name)
	assert.Equal(t, "admin", foundUser.Role)
	assert.Equal(t, resultUser.Name, foundUser.Name)
}

func TestUserRepository_CountDocuments(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(collection)

	// Create test user
	user := entities.NewUser("Test User", "test@example.com", "password123")
	user.SetRole("user")
	_, err := repo.InsertOne(user)
	require.NoError(t, err)

	// Test CountDocuments
	count, err := repo.CountDocuments("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	// Test CountDocuments with non-existent email
	count, err = repo.CountDocuments("nonexistent@example.com")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}

func TestUserRepository_UpdateRole(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(collection)

	// Create test user
	user := entities.NewUser("Test User", "test@example.com", "password123")
	user.SetRole("user")
	_, err := repo.InsertOne(user)
	require.NoError(t, err)

	// Test UpdateRole
	err = repo.UpdateRole("test@example.com", "admin")
	assert.NoError(t, err)

	// Verify role was updated
	foundUser, err := repo.GetUserByEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "admin", foundUser.Role)
}

func TestUserDocument_Mapping(t *testing.T) {
	// Test UserToDomain
	doc := models.UserDocument{
		ID:       primitive.NewObjectID(),
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "hashedpassword",
		Role:     "user",
	}

	user := models.UserToDomain(doc)
	assert.Equal(t, doc.ID.Hex(), user.ID)
	assert.Equal(t, doc.Name, user.Name)
	assert.Equal(t, doc.Email, user.Email)
	assert.Equal(t, doc.Password, user.Password)
	assert.Equal(t, doc.Role, user.Role)

	// Test UserFromDomain
	originalUser := entities.NewUser("Original User", "original@example.com", "password123")
	originalUser.SetRole("admin")
	originalUser.ID = "507f1f77bcf86cd799439011"

	convertedDoc, err := models.UserFromDomain(originalUser)
	assert.NoError(t, err)
	assert.Equal(t, originalUser.ID, convertedDoc.ID.Hex())
	assert.Equal(t, originalUser.Name, convertedDoc.Name)
	assert.Equal(t, originalUser.Email, convertedDoc.Email)
	assert.Equal(t, originalUser.Password, convertedDoc.Password)
	assert.Equal(t, originalUser.Role, convertedDoc.Role)
} 