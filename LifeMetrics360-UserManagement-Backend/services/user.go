package services

import (
	"context"
	"github.com/fzambone/LifeMetrics360-UserManagement/models"
	"github.com/fzambone/LifeMetrics360/LifeMetrics360-Utils-Backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	DB utils.DatabaseHelper
}

func NewUserService(db utils.DatabaseHelper) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	// TODO: Implement create user logic
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	// TODO: Implement retrieval logic
}

func (s *UserService) UpdateUser(id string) (*models.User, error) {
	// TODO: Implement update user logic
}

func (s *UserService) DeleteUser(id string) (primitive.ObjectID, error) {
	// TODO: Implement delete user logic
}
