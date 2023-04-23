package entity

import (
	"errors"

	"github.com/oaraujocesar/donates-control-api/pkg/entity"
	documents "github.com/paemuri/brdoc"
	"golang.org/x/crypto/bcrypt"
)

// Role is either admin or operator
// and it can be used as Role("admin") or Role("operator")
type Role string

const (
	Admin    Role = "admin"
	Operator Role = "operator"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	CPF      string    `json:"cpf"`
	Role     Role      `json:"role"`
	Password string    `json:"-"`
}

var (
	ErrPasswordRequired = errors.New("password is required")
	ErrInvalidRole      = errors.New("role must be either admin or operator")
	ErrNameRequired     = errors.New("name is required")
	ErrInvalidCPF       = errors.New("invalid CPF")
	ErrUserExists       = errors.New("user already exists")
)

func NewUser(name, cpf, password string, role Role) (*User, error) {
	if name == "" {
		return nil, ErrNameRequired
	}

	if cpf == "" || !documents.IsCPF(cpf) {
		return nil, ErrInvalidCPF
	}

	if password == "" {
		return nil, ErrPasswordRequired
	}

	if role != Admin && role != Operator {
		return nil, ErrInvalidRole
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		CPF:      cpf,
		Role:     role,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
