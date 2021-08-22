package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Price struct {
	ID         string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Value      float64   `json:"value" valid:"notnull"`
	ProductID  string    `json:"product_id" valid:"-" gorm:"column:product_id;type:uuid;notnull"`
	CurrencyID string    `json:"-" valid:"-" gorm:"column:currency_id;type:uuid;notnull"`
	Currency   Currency  `json:"currency" valid:"-"`
	CreatedAt  time.Time `json:"created_at" valid:"-"`
	UpdatedAt  time.Time `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewPrice(value float64, currency Currency) (*Price, error) {

	price := &Price{
		ID:        uuid.NewV4().String(),
		Value:     value,
		Currency:  currency,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := price.Validate()
	if err != nil {
		return nil, err
	}

	return price, nil

}

func (pe *Price) Validate() error {
	_, err := govalidator.ValidateStruct(pe)

	if err != nil {
		return err
	}

	return nil
}
