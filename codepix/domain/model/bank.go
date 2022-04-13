package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := validator.ValidateStruct(Bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.created_at = time.Now()

	err := Bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
