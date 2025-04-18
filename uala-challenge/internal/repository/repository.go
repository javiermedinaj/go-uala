package repository

import (
	"context"

	"github.com/javiermedinaj/uala-challenge/internal/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	FollowerUser(ctx context.Context, userID, targetUserID string) error
	GetFollowing(ctx context.Context, userID string) ([]*models.User, error)
	GetFollowers(ctx context.Context, userID string) ([]*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
type TweetRepository interface {
	CreateTweet(ctx context.Context, tweet *models.Tweet) error
	GetTweet(ctx context.Context, id string) (*models.Tweet, error)
}
