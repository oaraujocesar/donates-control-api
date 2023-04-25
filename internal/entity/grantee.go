package entity

import (
	"errors"

	id "github.com/oaraujocesar/donates-control-api/pkg/entity"
	documents "github.com/paemuri/brdoc"
)

type MaritalStatus string

const (
	Single      MaritalStatus = "solteiro"
	Married     MaritalStatus = "casado"
	Divorced    MaritalStatus = "divorciado"
	Widowed     MaritalStatus = "viúvo"
	StableUnion MaritalStatus = "união estável"
)

type Grantee struct {
	ID            id.ID         `json:"id"`
	Name          string        `json:"name"`
	CPF           string        `json:"cpf"`
	Birthday      string        `json:"birthday,omitempty"`
	MaritalStatus MaritalStatus `json:"marital_status,omitempty"`
	Occupation    string        `json:"occupation,omitempty"`
	FamilyCount   int           `json:"family_count,omitempty"`
	PhoneNumber   string        `json:"phone_number,omitempty"`
	NIS           string        `json:"nis,omitempty"`
	Address       string        `json:"address"`
	Neighborhood  string        `json:"neighborhood,omitempty"`
	IsRuralArea   bool          `json:"is_rural_area,omitempty"`
	Delivered     bool          `json:"delivered,omitempty"`
}

var (
	ErrInvalidMaritalStatus = errors.New("invalid marital status")
	ErrAddressRequired      = errors.New("address is required")
)

func NewGrantee(grantee *Grantee) (*Grantee, error) {
	if grantee.Name == "" {
		return nil, ErrNameRequired
	}

	if grantee.CPF == "" || !documents.IsCPF(grantee.CPF) {
		return nil, ErrInvalidCPF
	}

	if grantee.MaritalStatus != Single && grantee.MaritalStatus != Married && grantee.MaritalStatus != Divorced && grantee.MaritalStatus != Widowed && grantee.MaritalStatus != StableUnion {
		return nil, ErrInvalidMaritalStatus
	}

	if grantee.Address == "" {
		return nil, ErrAddressRequired
	}

	return &Grantee{
		ID:            id.NewID(),
		Name:          grantee.Name,
		CPF:           grantee.CPF,
		Birthday:      grantee.Birthday,
		MaritalStatus: grantee.MaritalStatus,
		Occupation:    grantee.Occupation,
		FamilyCount:   grantee.FamilyCount,
		PhoneNumber:   grantee.PhoneNumber,
		NIS:           grantee.NIS,
		Address:       grantee.Address,
		Neighborhood:  grantee.Neighborhood,
		IsRuralArea:   grantee.IsRuralArea,
		Delivered:     grantee.Delivered,
	}, nil
}
