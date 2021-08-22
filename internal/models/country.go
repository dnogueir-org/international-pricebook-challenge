package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Country struct {
	ID           string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name         string    `json:"name" valid:"notnull" gorm:"type:varchar(255)"`
	Abbreviation string    `json:"abbreviation" valid:"notnull" gorm:"type:varchar(255)"`
	CurrencyID   string    `json:"-" valid:"-" gorm:"column:currency_id;type:uuid;notnull"`
	Currency     Currency  `json:"currency" valid:"-"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewCountry(name, abbreviation string, currency Currency) (*Country, error) {
	country := &Country{
		ID:           uuid.NewV4().String(),
		Name:         name,
		Abbreviation: abbreviation,
		Currency:     currency,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := country.Validate()
	if err != nil {
		return nil, err
	}

	return country, nil

}

func (c *Country) Validate() error {
	_, err := govalidator.ValidateStruct(c)

	if err != nil {
		return err
	}

	return nil
}
