package types

import (
	"time"
)

type User struct {
	ID            int     `json:"id"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	Email         string  `json:"email"`
	Password      string  `json:"-"`
	RememberToken *string `json:"rememberToken"`
	Roles         []Role  `json:"-" gorm:"many2many:user_roles;"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"default:0"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
	UpdateUserRememberToken(*User, string) error
	GetUserByRememberToken(rememberToken string) (*User, error)
	UpdatePasswordOfUser(*User, string) error
}

type ProductStore interface {
	GetProducts(page, size int, search, category string) ([]Product, int64, error)
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResetPasswordPayload struct {
	Email string `json:"email" validate:"required,email"`
}

type CreatePasswordPayload struct {
	Password      string `json:"password" validate:"required"`
	RememberToken string `json:"rememberToken" validate:"required"`
}

type Event struct {
	Name    string
	Payload interface{}
}

type Listener func(Event)
