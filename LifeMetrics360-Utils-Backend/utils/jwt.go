package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

func GenerateToken(userID primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID.Hex()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
