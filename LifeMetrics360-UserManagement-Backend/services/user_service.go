package services

import (
	"context"
	"github.com/fzambone/LifeMetrics360-UserManagement/models"
	"github.com/fzambone/LifeMetrics360/LifeMetrics360-Utils-Backend/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	DB utils.DatabaseHelper
}

func NewUserService(db utils.DatabaseHelper) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	result, err := s.DB.InsertOne(ctx, "users", user)
	if err != nil {
		logrus.WithError(err).Error("Failed to create user")
		return primitive.NilObjectID, err
	}

	return result, nil
}
