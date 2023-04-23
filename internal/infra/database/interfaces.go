package database

import "github.com/oaraujocesar/donates-control-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByCPF(cpf string) (*entity.User, error)
}
