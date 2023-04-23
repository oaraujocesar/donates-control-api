package database

import (
	"github.com/oaraujocesar/donates-control-api/internal/entity"
	documents "github.com/paemuri/brdoc"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entity.User) error {
	userFound, _ := u.FindByCPF(user.CPF)
	if userFound != nil {
		return entity.ErrUserExists
	}

	return u.DB.Create(user).Error
}

func (u *User) FindByCPF(cpf string) (*entity.User, error) {
	if cpf == "" || !documents.IsCPF(cpf) {
		return nil, entity.ErrInvalidCPF
	}

	var user entity.User
	err := u.DB.Where("cpf = ?", cpf).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
