package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID        string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" valid:"notnull" gorm:"type:varchar(255)"`
	Prices    []Price   `json:"prices" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduct(name string) (*Product, error) {

	product := &Product{
		ID:        uuid.NewV4().String(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}
func (p *Product) Validate() error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	return nil
}
