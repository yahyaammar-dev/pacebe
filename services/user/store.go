package user

import (
	"fmt"

	"github.com/yahyaammar-dev/pacebe/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	result := s.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	result := s.db.Preload("Roles").Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (s *Store) UpdateUserRememberToken(user *types.User, token string) error {
	result := s.db.Model(&user).Update("remember_token", token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Store) GetUserByRememberToken(rememberToken string) (*types.User, error) {
	var user types.User
	result := s.db.Where("remember_token = ?", rememberToken).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (s *Store) UpdatePasswordOfUser(user *types.User, password string) error {
	result := s.db.Model(&user).Updates(map[string]interface{}{
		"password":       password,
		"remember_token": "",
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
