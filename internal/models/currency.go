package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Currency struct {
	ID        string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" valid:"notnull" gorm:"type:varchar(255)"`
	Acronym   string    `json:"acronym" valid:"notnull" gorm:"type:varchar(255)"`
	Prices    []Price   `json:"-" valid:"-"`
	Country   []Country `json:"-" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCurrency(name, acronym string) (*Currency, error) {
	currency := &Currency{
		ID:        uuid.NewV4().String(),
		Name:      name,
		Acronym:   acronym,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := currency.Validate()
	if err != nil {
		return nil, err
	}

	return currency, nil

}

func (c *Currency) Validate() error {
	_, err := govalidator.ValidateStruct(c)

	if err != nil {
		return err
	}

	return nil
}
