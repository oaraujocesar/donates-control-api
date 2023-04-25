package database

import (
	"github.com/oaraujocesar/donates-control-api/internal/entity"
	"gorm.io/gorm"
)

type Grantee struct {
	DB *gorm.DB
}

func NewGrantee(db *gorm.DB) *Grantee {
	return &Grantee{
		DB: db,
	}
}

func (g *Grantee) Create(grantee *entity.Grantee) error {
	granteeFound, _ := g.FindByCPF(grantee.CPF)
	if granteeFound != nil {
		return entity.ErrGranteeExists
	}

	return g.DB.Create(grantee).Error
}

func (g *Grantee) FindByCPF(cpf string) (*entity.Grantee, error) {
	var grantee entity.Grantee
	err := g.DB.Where("cpf = ?", cpf).First(&grantee).Error
	if err != nil {
		return nil, entity.ErrGranteeNotFound
	}

	return &grantee, nil
}

func (g *Grantee) Update(grantee *entity.Grantee) error {
	_, err := g.FindByCPF(grantee.CPF)
	if err != nil {
		return entity.ErrGranteeNotFound
	}

	return g.DB.Save(grantee).Error
}

func (g *Grantee) Delete(cpf string) error {
	grantee, err := g.FindByCPF(cpf)
	if err != nil {
		return entity.ErrGranteeNotFound
	}

	return g.DB.Delete(grantee).Error
}

func (g *Grantee) Deliver(cpf string) error {
	grantee, err := g.FindByCPF(cpf)
	if err != nil {
		return entity.ErrGranteeNotFound
	}

	grantee.Delivered = true

	return g.DB.Save(grantee).Error
}
