package database

import "github.com/oaraujocesar/donates-control-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByCPF(cpf string) (*entity.User, error)
}

type GranteeInterface interface {
	Create(grantee *entity.Grantee) error
	FindByID(id string) (*entity.Grantee, error)
	FindAll() ([]entity.Grantee, error)
	Update(grantee *entity.Grantee) error
	Delete(id string) error
}
