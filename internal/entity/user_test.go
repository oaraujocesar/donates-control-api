package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserAdmin(t *testing.T) {
	user, err := NewUser("César", "602.305.720-90", "test123", "admin")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Role)
	assert.Equal(t, "César", user.Name)
	assert.Equal(t, "602.305.720-90", user.CPF)
	assert.Equal(t, Role("admin"), user.Role)
}

func TestNewUserOperator(t *testing.T) {
	user, err := NewUser("César", "602.305.720-90", "test123", "operator")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Role)
	assert.Equal(t, "César", user.Name)
	assert.Equal(t, "602.305.720-90", user.CPF)
	assert.Equal(t, Role("operator"), user.Role)
}

func TestNewUserWithInvalidCPF(t *testing.T) {
	user, err := NewUser("César", "1111", "test123", "admin")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, ErrInvalidCPF, err)
}

func TestNewUserWithInvalidRole(t *testing.T) {
	user, err := NewUser("César", "602.305.720-90", "test123", "invalid")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, ErrInvalidRole, err)
}

func TestNewUserWithEmptyName(t *testing.T) {
	user, err := NewUser("", "602.305.720-90", "test123", "admin")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, ErrNameRequired, err)
}

func TestNewUserWithEmptyPassword(t *testing.T) {
	user, err := NewUser("César", "602.305.720-90", "", "admin")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, ErrPasswordRequired, err)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("César", "602.305.720-90", "test123", "admin")
	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("test123"))
	assert.False(t, user.ValidatePassword("test"))
	assert.False(t, user.ValidatePassword("test1234"))
	assert.NotEqual(t, "test123", user.Password)
}
