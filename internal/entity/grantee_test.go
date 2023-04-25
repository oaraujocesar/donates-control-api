package entity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrantee(t *testing.T) {
	grantee, err := NewGrantee(&Grantee{
		Name:          "Caio",
		CPF:           "929.500.850-24",
		Birthday:      "07/08/97",
		MaritalStatus: MaritalStatus("solteiro"),
		Occupation:    "Desempregado",
		FamilyCount:   0,
		PhoneNumber:   "(81)123456789",
		NIS:           "1234567891011",
		Address:       "Rua dos Bobos, n 0",
		Neighborhood:  "País das Maravilhas",
		IsRuralArea:   false,
		Delivered:     false,
	})

	assert.Nil(t, err)
	assert.NotNil(t, grantee)
	assert.NotEmpty(t, grantee.ID)
	assert.NotEmpty(t, grantee.CPF)
	assert.NotEmpty(t, grantee.Birthday)
	assert.NotEmpty(t, grantee.MaritalStatus)
	assert.NotEmpty(t, grantee.Occupation)
	assert.Equal(t, reflect.TypeOf(1), reflect.TypeOf(grantee.FamilyCount))
	assert.NotEmpty(t, grantee.PhoneNumber)
	assert.NotEmpty(t, grantee.NIS)
	assert.NotEmpty(t, grantee.Address)
	assert.NotEmpty(t, grantee.Neighborhood)
	assert.Equal(t, false, grantee.IsRuralArea)
	assert.Equal(t, false, grantee.Delivered)
	assert.Equal(t, "Caio", grantee.Name)
	assert.Equal(t, "929.500.850-24", grantee.CPF)
	assert.Equal(t, MaritalStatus("solteiro"), grantee.MaritalStatus)
}

func TestNewGranteeWithInvalidCPF(t *testing.T) {
	grantee, err := NewGrantee(&Grantee{
		Name:          "Caio",
		CPF:           "",
		Birthday:      "07/08/97",
		MaritalStatus: MaritalStatus("solteiro"),
		Occupation:    "Desempregado",
		FamilyCount:   0,
		PhoneNumber:   "(81)123456789",
		NIS:           "1234567891011",
		Address:       "Rua dos Bobos, n 0",
		Neighborhood:  "País das Maravilhas",
		IsRuralArea:   false,
		Delivered:     false,
	})

	assert.NotNil(t, err)
	assert.Nil(t, grantee)
	assert.Equal(t, ErrInvalidCPF, err)

	grantee, err = NewGrantee(&Grantee{
		Name:          "Caio",
		CPF:           "111",
		Birthday:      "07/08/97",
		MaritalStatus: MaritalStatus("solteiro"),
		Occupation:    "Desempregado",
		FamilyCount:   0,
		PhoneNumber:   "(81)123456789",
		NIS:           "1234567891011",
		Address:       "Rua dos Bobos, n 0",
		Neighborhood:  "País das Maravilhas",
		IsRuralArea:   false,
		Delivered:     false,
	})
	assert.NotNil(t, err)
	assert.Nil(t, grantee)
	assert.Equal(t, ErrInvalidCPF, err)
}

func TestNewGranteeWithInvalidMaritalStatus(t *testing.T) {
	grantee, err := NewGrantee(&Grantee{
		Name:          "Caio",
		CPF:           "929.500.850-24",
		Birthday:      "07/08/97",
		MaritalStatus: "wrong",
		Occupation:    "Desempregado",
		FamilyCount:   0,
		PhoneNumber:   "(81)123456789",
		NIS:           "1234567891011",
		Address:       "Rua dos Bobos, n 0",
		Neighborhood:  "País das Maravilhas",
		IsRuralArea:   false,
		Delivered:     false,
	})

	assert.NotNil(t, err)
	assert.Nil(t, grantee)
	assert.Equal(t, ErrInvalidMaritalStatus, err)
}

func TestNewGranteeWithoutName(t *testing.T) {
	grantee, err := NewGrantee(&Grantee{
		Name:          "",
		CPF:           "929.500.850-24",
		Birthday:      "07/08/97",
		MaritalStatus: MaritalStatus("solteiro"),
		Occupation:    "Desempregado",
		FamilyCount:   0,
		PhoneNumber:   "(81)123456789",
		NIS:           "1234567891011",
		Address:       "Rua dos Bobos, n 0",
		Neighborhood:  "País das Maravilhas",
		IsRuralArea:   false,
		Delivered:     false,
	})

	assert.NotNil(t, err)
	assert.Nil(t, grantee)
	assert.Equal(t, ErrNameRequired, err)
}

func TestNewGranteeWithoutAddress(t *testing.T) {
	grantee, err := NewGrantee(&Grantee{
		Name:          "Caio",
		CPF:           "929.500.850-24",
		Birthday:      "07/08/97",
		MaritalStatus: MaritalStatus("solteiro"),
		Occupation:    "Desempregado",
		FamilyCount:   0,
		PhoneNumber:   "(81)123456789",
		NIS:           "1234567891011",
		Address:       "",
		Neighborhood:  "País das Maravilhas",
		IsRuralArea:   false,
		Delivered:     false,
	})

	assert.NotNil(t, err)
	assert.Nil(t, grantee)
	assert.Equal(t, ErrAddressRequired, err)
}
