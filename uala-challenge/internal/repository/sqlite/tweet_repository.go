package sqlite

import (
	"context"
	"database/sql"

	"github.com/javiermedinaj/uala-challenge/internal/models"
)

type TweetRepository struct {
	db *sql.DB
}

func NewTweetRepository(db *sql.DB) *TweetRepository {
	return &TweetRepository{db: db}
}

func (r *TweetRepository) CreateTweet(ctx context.Context, tweet *models.Tweet) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO tweets (id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		tweet.ID, tweet.UserID, tweet.Content, tweet.CreatedAt, tweet.UpdatedAt)
	return err
}

func (r *TweetRepository) GetTweet(ctx context.Context, id string) (*models.Tweet, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, user_id, content, created_at, updated_at FROM tweets WHERE id = ?", id)
	var tweet models.Tweet
	if err := row.Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tweet not found
		}
		return nil, err
	}
	return &tweet, nil
}
