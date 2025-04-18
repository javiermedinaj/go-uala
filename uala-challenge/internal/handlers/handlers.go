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

func NewHandler(userRepo repository.UserRepository, tweetRepo repository.TweetRepository) *Handler {
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

func (h *Handler) CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		ID:        uuid.New().String(),
		Username:  input.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.userRepo.CreateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userRepo.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetTweetByID(c *gin.Context) {
	id := c.Param("id")
	tweet, err := h.tweetRepo.GetTweet(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get tweet"})
		return
	}
	if tweet == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
		return
	}
	c.JSON(http.StatusOK, tweet)
}
