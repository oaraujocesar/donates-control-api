package database

import (
	"testing"

	"github.com/oaraujocesar/donates-control-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createInMemoryDatabase(t *testing.T, table interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(table)

	return db, nil
}

func TestCreateAdmin(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.User{})

	user, _ := entity.NewUser("Cesar", "602.305.720-90", "123456", entity.Role("admin"))
	assert.Nil(t, err)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.NotEmpty(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.CPF, userFound.CPF)
	assert.Equal(t, user.Role, userFound.Role)
	assert.NotNil(t, userFound.Password)
}

func TestCreateOperator(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.User{})

	user, _ := entity.NewUser("Cesar", "602.305.720-90", "123456", entity.Role("operator"))
	assert.Nil(t, err)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.NotEmpty(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.CPF, userFound.CPF)
	assert.Equal(t, user.Role, userFound.Role)
	assert.NotNil(t, userFound.Password)
}

func TestCreateUserWhenCPFIsAlreadyRegistered(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.User{})
	assert.Nil(t, err)

	userDB := NewUser(db)

	user, err := entity.NewUser("Cesar", "602.305.720-90", "123123", entity.Role("admin"))
	assert.Nil(t, err)

	err = userDB.Create(user)
	assert.Nil(t, err)

	user2, err := entity.NewUser("Cesar 2", "602.305.720-90", "123123", entity.Role("admin"))
	assert.Nil(t, err)

	err = userDB.Create(user2)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "user already exists")
}

func TestFindUserByCPF(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.User{})

	user, _ := entity.NewUser("Cesar", "602.305.720-90", "123123", entity.Role("admin"))
	assert.Nil(t, err)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	userFound, err := userDB.FindByCPF(user.CPF)

	assert.Nil(t, err)
	assert.NotEmpty(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.CPF, userFound.CPF)
	assert.NotNil(t, userFound.Password)
}

func TestFindUserByCPFWhenCPFIsInvalid(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.User{})
	assert.Nil(t, err)

	userDB := NewUser(db)

	_, err = userDB.FindByCPF("123")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid CPF")
}

func TestFindByCPFWhenUserDoesNotExists(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.User{})
	assert.Nil(t, err)

	userDB := NewUser(db)

	_, err = userDB.FindByCPF("602.305.720-90")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "record not found")
}
