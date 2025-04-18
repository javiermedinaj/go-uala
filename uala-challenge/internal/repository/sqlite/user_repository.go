package sqlite

import (
	"context"
	"database/sql"

	"github.com/javiermedinaj/uala-challenge/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, username, created_at, updated_at FROM users WHERE id = ?", id)
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, username, created_at, updated_at FROM users WHERE username = ?", username)
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, username, created_at, updated_at) VALUES (?, ?, ?, ?)", user.ID, user.Username, user.CreatedAt, user.UpdatedAt)
	return err
}

//implementar otras funciones para el repositorio de usuarios luego igual : p

func (r *UserRepository) FollowerUser(ctx context.Context, userID, targetUserID string) error {
	return nil // Obvio no son la logica real
}

func (r *UserRepository) GetFollowing(ctx context.Context, userID string) ([]*models.User, error) {
	return nil, nil
}

func (r *UserRepository) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	return nil, nil
}
