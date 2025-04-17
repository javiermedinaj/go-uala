package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/javiermedinaj/uala-challenge/internal/models"
	"github.com/javiermedinaj/uala-challenge/internal/repository"
)

type Handler struct {
	userRepo  repository.UserRepository
	tweetRepo repository.TweetRepository
}

func newHandler(userRepo repository.UserRepository, tweetRepo repository.TweetRepository) *Handler {
	return &Handler{
		userRepo:  userRepo,
		tweetRepo: tweetRepo,
	}
}

func (h *Handler) CreateTweet(c *gin.Context) {
	userID := c.GetHeader("X-User-ID") // Según el requisito, el ID viene en el header
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Content) > 280 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tweet exceeds 280 character limit"})
		return
	}

	user, err := h.userRepo.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}

	tweet := &models.Tweet{
		ID:        uuid.New().String(),
		UserID:    userID,
		Username:  user.Username,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}

	if err := h.tweetRepo.CreateTweet(c.Request.Context(), tweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create tweet"})
		return
	}
	c.JSON(http.StatusCreated, tweet)
}
