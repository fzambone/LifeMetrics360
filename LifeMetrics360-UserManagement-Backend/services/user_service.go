package services

import (
	"context"
	"errors"
	"github.com/fzambone/LifeMetrics360-UserManagement/models"
	"github.com/fzambone/LifeMetrics360-UserManagement/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const userCollection = "users"

type UserService struct {
	DB *utils.Database
}

func NewUserService(db *utils.Database) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	result, err := s.DB.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("error casting inserted ID to ObjectID")
	}

	return insertedID, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var user models.User

	if err := s.DB.Collection(userCollection).FindOne(ctx, bson.M{"_id": objID}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	if err := s.DB.Collection(userCollection).FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userID string, updateData *models.User) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": updateData}
	result, err := s.DB.Collection(userCollection).UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	result, err := s.DB.Collection(userCollection).DeleteOne(ctx, bson.M{"_id": objID})

	if result.DeletedCount == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

// ValidateUserCredentials validates a user's credentials
func (s *UserService) ValidateUserCredentials(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalida credentials")
	}

	return user, nil
}
