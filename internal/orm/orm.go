package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/luisberga/phones-api/internal/config"
	"github.com/luisberga/phones-api/internal/models"
)

type ORM struct {
	DB *gorm.DB
}

func TypeAssertion(model interface{}) interface{} {
	value, ok := model.(*models.Company)
	if !ok {
		value, ok := model.(*models.Phone)
		if !ok {
			return ""
		}
		return value.CompanyID
	}
	return value.Name
}

func (d *ORM) Connect() error {
	db, err := gorm.Open(mysql.Open(config.ORMConnection), &gorm.Config{})
	if err != nil {
		return err
	}

	d.DB = db
	return err
}

func (d *ORM) Insert(table string, columnName string, model interface{}) error {

	result := d.DB.Select(columnName).Create(model)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *ORM) QueryID(table string, ID string, model interface{}) error {
	result := d.DB.First(model, ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *ORM) Update(table string, columnName string, model interface{}, ID string) error {
	value := TypeAssertion(model)

	result := d.DB.Model(model).Where("id = ?", ID).Update(columnName, value)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *ORM) InsertMany(table string, model interface{}) error {
	result := d.DB.Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *ORM) Delete(table string, ID string, model interface{}) error {
	result := d.DB.Delete(model, ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *ORM) Query(table string, columnName string, values string, model interface{}) error {
	queryString := fmt.Sprintf("%s LIKE ?", columnName)
	result := d.DB.Where(queryString, values).Find(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *ORM) QueryAvailable(table string, columnName string, ID string, model interface{}) error {
	queryString := fmt.Sprintf("%s = ?", columnName)
	result := d.DB.Where(queryString, ID).First(model)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (d *ORM) QueryCount(table string, columnName string, model interface{}, modelGroup interface{}) error {
	phone, _ := model.(*models.Phone)
	phones, _ := modelGroup.(*[]models.PhoneGroup)

	d.DB.Model(phone).Select("company_id, count(company_id) as available_phones").Group("company_id").Scan(&phones)

	return nil
}
