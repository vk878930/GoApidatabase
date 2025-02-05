package repository

import (
	"database/sql"
	"submission-project-enigma-laundry/model"
	"fmt"
)

type UserRepository interface {
	Create(user model.UserCredential) (model.UserCredential, error)
	List() ([]model.UserCredential, error)
	Get(id uint32) (model.UserCredential, error)
	GetByUsernamePassword(username string, password string) (model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(user model.UserCredential) (model.UserCredential, error) {
	err := u.db.QueryRow("INSERT INTO users (username, password, role) VALUES  ($1, $2, $3) RETURNING id").Scan(&user.Id)
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to save user")
	}
	return user, nil
}

func (u *userRepository) List() ([]model.UserCredential, error) {
	var users []model.UserCredential
	rows, err := u.db.Query("SELECT id, username, role FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list user")
	}
	for rows.Next() {
		var user model.UserCredential
		err := rows.Scan(&user.Id, &user.Username, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("failed to scaning data")
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) Get(id uint32) (model.UserCredential, error) {
	var user model.UserCredential
	err := u.db.QueryRow("SELECT id, username, role FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Role)
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to get user by ID")
	}
	return user, nil
}

func (u *userRepository) GetByUsernamePassword(username string, password string) (model.UserCredential, error) {
	var user model.UserCredential
	err := u.db.QueryRow("SELECT id, username, role FROM users WHERE username = $1 and password = $2", username, password).Scan(&user.Id, &user.Username, &user.Role)
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to get user by username and password")
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
