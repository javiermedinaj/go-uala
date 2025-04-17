package models

import (
	"time"
)

type User struct {
	ID        string
	Username  string
	Following []string
	Followers []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweet struct {
	ID           string
	UserID       string
	Content      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LikeCount    int
	RetweetCount int
	CommentCount int
	Replies      []Tweet
	Retweets     []Tweet
	Likes        []Tweet
	Comments     []Tweet
}
