package database

import "github.com/gustavorafaeldev/curso-go/APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}