package account

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/altanbgn/arctis/internal/db"
	"github.com/altanbgn/arctis/internal/models"
  "github.com/altanbgn/arctis/internal/utils"
)

func GetByIDService(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	var r models.User
	if err := db.GetCollection("users").FindOne(ctx, bson.M{"_id": id}).Decode(&r); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New("User not found")
		}

		log.Println(err)
	}

	return r, nil
}

func GetByUsernameService(ctx context.Context, username string) (models.User, error) {
	var r models.User
	if err := db.GetCollection("users").FindOne(ctx, bson.M{"username": username}).Decode(&r); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New("User not found")
		}

		log.Println(err)
	}

	return r, nil
}

func CreateService(ctx context.Context, user CreateAccountPayload) error {
	if hashedPassword, err := utils.CreateHash(user.Password, utils.DefaultParams); err != nil {
		log.Println(err)
	} else {
		user.Password = hashedPassword
	}

	if _, err := db.GetCollection("users").InsertOne(ctx, user); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("User already exists")
		}

		return errors.New("Failed to create user")
	}

	return nil
}

func UpdateService(ctx context.Context, id primitive.ObjectID, user UpdateAccountPayload) error {
  if _, err := db.GetCollection("users").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user}); err != nil {
    return errors.New("Failed to update user")
  }

  return nil
}

func DeleteService(ctx context.Context, id primitive.ObjectID) error {
  if _, err := db.GetCollection("users").DeleteOne(ctx, bson.M{"_id": id}); err != nil {
    return errors.New("Failed to delete user")
  }

  return nil
}
