package services

import (
	"context"
	"github.com/fzambone/LifeMetrics360-UserManagement/models"
	"github.com/fzambone/LifeMetrics360-UserManagement/utils"
	"go.mongodb.org/mongo-driver/bson"
)

const avatarCollection = "avatars"

type AvatarService struct {
	DB *utils.Database
}

func NewAvatarService(db *utils.Database) *AvatarService {
	return &AvatarService{DB: db}
}

func (s *AvatarService) GetAvatarByEmail(ctx context.Context, userEmail string) (string, error) {
	var avatar models.Avatar
	if err := s.DB.Collection(avatarCollection).FindOne(ctx, bson.M{"userEmail": userEmail}).Decode(&avatar); err != nil {
		return "", err
	}
	return avatar.AvatarURL, nil
}
