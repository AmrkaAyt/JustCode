package repository

import (
	"Lecture9/internal/user/entity"
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetById(ctx context.Context, userID int) (*entity.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, userID)

	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, userID int, user entity.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, userID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *UserRepository) Register(ctx context.Context, user entity.User, hashedPassword []byte) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var userID int
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *UserRepository) Login(ctx context.Context, user entity.User) (*entity.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	row := r.db.QueryRowContext(ctx, query, user.Email)

	var storedUser entity.User
	if err := row.Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return &storedUser, nil
}
