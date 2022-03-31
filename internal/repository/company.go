package repository

import (
	"fmt"

	"github.com/luisberga/phones-api/internal/models"
)

type Company struct {
	database models.DatabaseInterface
}

func NewCompanyRepo(database models.DatabaseInterface) *Company {
	return &Company{database}
}

func (c *Company) Create(company *models.Company) error {
	err := c.database.Connect()
	if err != nil {
		return err
	}

	err = c.database.Insert("companies", "name", company)
	if err != nil {
		return err
	}

	return nil
}

func (c *Company) List(companies *[]models.Company, name string) error {
	err := c.database.Connect()
	if err != nil {
		return nil
	}
	name = fmt.Sprintf("%%%s%%", name)

	err = c.database.Query("companies", "name", name, companies)
	if err != nil {
		return nil
	}

	return nil
}

func (c *Company) ListID(company *models.Company, ID string) error {
	err := c.database.Connect()
	if err != nil {
		return err
	}
	err = c.database.QueryID("companies", ID, company)
	if err != nil {
		return err
	}

	return nil
}

func (c *Company) Update(ID string, company *models.Company) error {
	err := c.database.Connect()
	if err != nil {
		return err
	}

	err = c.database.Update("companies", "name", company, ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Company) Delete(ID string, company *models.Company) error {
	err := c.database.Connect()
	if err != nil {
		return err
	}

	err = c.database.Delete("companies", ID, company)
	if err != nil {
		return err
	}

	return nil
}
