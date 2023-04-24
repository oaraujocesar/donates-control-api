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
	Birthday      string        `json:"birthday"`
	MaritalStatus MaritalStatus `json:"marital_status"`
	Occupation    string        `json:"occupation"`
	FamilyCount   int           `json:"family_count"`
	PhoneNumber   string        `json:"phone_number"`
	NIS           string        `json:"nis"`
	Address       string        `json:"address"`
	Neighborhood  string        `json:"neighborhood"`
	IsRuralArea   bool          `json:"is_rural_area"`
	Delivered     bool          `json:"delivered"`
}

var (
	ErrBirthdayRequired     = errors.New("birthday is required")
	ErrInvalidMaritalStatus = errors.New("invalid marital status")
	ErrOccupationRequired   = errors.New("occupation is required")
	ErrPhoneNumberRequired  = errors.New("phone number is required")
	ErrNISRequired          = errors.New("NIS is required")
	ErrAddressRequired      = errors.New("address is required")
	ErrNeighborhoodRequired = errors.New("neighborhood is required")
)

func NewGrantee(grantee *Grantee) (*Grantee, error) {
	if grantee.Name == "" {
		return nil, ErrNameRequired
	}

	if grantee.CPF == "" || !documents.IsCPF(grantee.CPF) {
		return nil, ErrInvalidCPF
	}

	if grantee.Birthday == "" {
		return nil, ErrBirthdayRequired
	}

	if grantee.MaritalStatus != Single && grantee.MaritalStatus != Married && grantee.MaritalStatus != Divorced && grantee.MaritalStatus != Widowed && grantee.MaritalStatus != StableUnion {
		return nil, ErrInvalidMaritalStatus
	}

	if grantee.Occupation == "" {
		return nil, ErrOccupationRequired
	}

	if grantee.PhoneNumber == "" {
		return nil, ErrPhoneNumberRequired
	}

	if grantee.NIS == "" {
		return nil, ErrNISRequired
	}

	if grantee.Address == "" {
		return nil, ErrAddressRequired
	}

	if grantee.Neighborhood == "" {
		return nil, ErrNeighborhoodRequired
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
