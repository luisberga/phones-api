package repository

import (
	"strconv"

	"github.com/luisberga/phones-api/internal/models"
)

type Phone struct {
	database models.DatabaseInterface
}

func NewPhoneRepo(database models.DatabaseInterface) *Phone {
	return &Phone{database}
}

func (p *Phone) ListAvailable(phone *models.Phone, ID string) error {
	err := p.database.Connect()
	if err != nil {
		return err
	}
	err = p.database.QueryAvailable("phones", "company_id", ID, phone)
	if err != nil {
		return err
	}

	err = p.database.Delete("phones", strconv.FormatUint(phone.ID, 10), phone)
	if err != nil {
		return err
	}

	return nil

}

func (p *Phone) Create(phones *[]models.Phone) error {
	err := p.database.Connect()
	if err != nil {
		return err
	}

	err = p.database.InsertMany("phones", phones)
	if err != nil {
		return err
	}

	return nil
}

func (p *Phone) List(phone *models.Phone, phones *[]models.PhoneGroup) error {
	err := p.database.Connect()
	if err != nil {
		return err
	}

	err = p.database.QueryCount("phones", "company_id", phone, phones)
	if err != nil {
		return err
	}

	return nil
}
