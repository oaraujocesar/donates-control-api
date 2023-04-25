package database

import (
	"testing"

	"github.com/oaraujocesar/donates-control-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateGrantee(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Grantee{})
	assert.Nil(t, err)

	g := &entity.Grantee{
		Name:          "Cesar",
		CPF:           "602.305.720-90",
		Birthday:      "13/02/2000",
		MaritalStatus: entity.MaritalStatus("casado"),
		Occupation:    "Desenvolvedor",
		FamilyCount:   6,
		PhoneNumber:   "(11) 99999-9999",
		NIS:           "123456789",
		Address:       "Rua dos bobos",
		Neighborhood:  "Vila do Chaves",
		IsRuralArea:   false,
	}

	grantee, _ := entity.NewGrantee(g)

	granteeDB := NewGrantee(db)
	err = granteeDB.Create(grantee)
	assert.Nil(t, err)

	var granteeFound entity.Grantee
	err = db.First(&granteeFound, "id = ?", grantee.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, grantee.ID, granteeFound.ID)
	assert.Equal(t, grantee.Name, granteeFound.Name)
	assert.Equal(t, grantee.CPF, granteeFound.CPF)
	assert.Equal(t, grantee.Birthday, granteeFound.Birthday)
	assert.Equal(t, grantee.MaritalStatus, granteeFound.MaritalStatus)
	assert.Equal(t, grantee.Occupation, granteeFound.Occupation)
	assert.Equal(t, grantee.FamilyCount, granteeFound.FamilyCount)
	assert.Equal(t, grantee.PhoneNumber, granteeFound.PhoneNumber)
	assert.Equal(t, grantee.NIS, granteeFound.NIS)
	assert.Equal(t, grantee.Address, granteeFound.Address)
	assert.Equal(t, grantee.Neighborhood, granteeFound.Neighborhood)
	assert.Equal(t, grantee.IsRuralArea, granteeFound.IsRuralArea)
	assert.Equal(t, grantee.Delivered, granteeFound.Delivered)
}

func TestGranteeAlreadyExists(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Grantee{})
	assert.Nil(t, err)

	g := &entity.Grantee{
		Name:          "Cesar",
		CPF:           "602.305.720-90",
		Birthday:      "13/02/2000",
		MaritalStatus: entity.MaritalStatus("casado"),
		Occupation:    "Desenvolvedor",
		FamilyCount:   6,
		PhoneNumber:   "(11) 99999-9999",
		NIS:           "123456789",
		Address:       "Rua dos bobos",
		Neighborhood:  "Vila do Chaves",
		IsRuralArea:   false,
	}

	grantee, _ := entity.NewGrantee(g)

	granteeDB := NewGrantee(db)
	err = granteeDB.Create(grantee)
	assert.Nil(t, err)

	err = granteeDB.Create(grantee)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrGranteeExists.Error(), err.Error())
}

func TestUpdateGrantee(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Grantee{})
	assert.Nil(t, err)

	g := &entity.Grantee{
		Name:          "Cesar",
		CPF:           "602.305.720-90",
		Address:       "Rua dos bobos",
		Neighborhood:  "Vila do Chaves",
		MaritalStatus: entity.MaritalStatus("solteiro"),
		IsRuralArea:   false,
	}

	grantee, _ := entity.NewGrantee(g)

	granteeDB := NewGrantee(db)
	err = granteeDB.Create(grantee)
	assert.Nil(t, err)

	grantee.Name = "Cesar Augusto"
	grantee.Address = "Rua dos bobos 2"
	grantee.Neighborhood = "Vila do Chaves 2"
	grantee.IsRuralArea = true
	grantee.MaritalStatus = entity.MaritalStatus("casado")

	err = granteeDB.Update(grantee)
	assert.Nil(t, err)

	var granteeFound entity.Grantee
	err = db.First(&granteeFound, "id = ?", grantee.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, grantee.ID, granteeFound.ID)
	assert.Equal(t, grantee.Name, granteeFound.Name)
	assert.Equal(t, grantee.CPF, granteeFound.CPF)
	assert.Equal(t, grantee.Address, granteeFound.Address)
	assert.Equal(t, grantee.Neighborhood, granteeFound.Neighborhood)
	assert.Equal(t, grantee.IsRuralArea, granteeFound.IsRuralArea)
	assert.Equal(t, grantee.MaritalStatus, granteeFound.MaritalStatus)
}

func TestUpdateWhenGranteeDoesNotExists(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Grantee{})
	assert.Nil(t, err)

	g := &entity.Grantee{
		Name:          "Cesar",
		CPF:           "602.305.720-90",
		Address:       "Rua dos bobos",
		Neighborhood:  "Vila do Chaves",
		MaritalStatus: entity.MaritalStatus("solteiro"),
		IsRuralArea:   false,
	}

	grantee, _ := entity.NewGrantee(g)

	granteeDB := NewGrantee(db)
	err = granteeDB.Update(grantee)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrGranteeNotFound.Error(), err.Error())
}

func TestDeleteGrantee(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Grantee{})
	assert.Nil(t, err)

	g := &entity.Grantee{
		Name:          "Cesar",
		CPF:           "602.305.720-90",
		Address:       "Rua dos bobos",
		Neighborhood:  "Vila do Chaves",
		MaritalStatus: entity.MaritalStatus("solteiro"),
		IsRuralArea:   false,
	}

	grantee, _ := entity.NewGrantee(g)

	granteeDB := NewGrantee(db)
	err = granteeDB.Create(grantee)
	assert.Nil(t, err)

	err = granteeDB.Delete(grantee.CPF)
	assert.Nil(t, err)

	_, err = granteeDB.FindByCPF(grantee.CPF)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrGranteeNotFound.Error(), err.Error())
}

func TestDeleteNotExistantGrantee(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Grantee{})
	assert.Nil(t, err)

	granteeDB := NewGrantee(db)
	err = granteeDB.Delete("602.305.720-90")
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrGranteeNotFound.Error(), err.Error())
}
