package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver

	"github.com/luisberga/phones-api/internal/config"
	"github.com/luisberga/phones-api/internal/models"
)

type Mysql struct {
	DB *sql.DB
}

func TypeAssertion(model interface{}) interface{} {
	value, ok := model.(*models.Company)
	if !ok {
		value, ok := model.(*models.Phone)
		if !ok {
			return ""
		}
		return value.ID
	}
	return value.Name
}

func UpdateID(model interface{}, ID uint64) {
	value, ok := model.(*models.Company)
	if !ok {
		value, ok := model.(*models.Phone)
		if !ok {
			return
		}
		value.ID = ID
		return
	}
	value.ID = ID
}

func (d *Mysql) Connect() error {
	db, err := sql.Open("mysql", config.DatabaseStringConnection)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}
	d.DB = db
	return err
}

func (d *Mysql) Insert(table string, columnName string, model interface{}) error {
	value := TypeAssertion(model)

	x := fmt.Sprintf("insert into %s (%s) values(?)", table, columnName)
	statement, err := d.DB.Prepare(x)
	if err != nil {
		return err
	}
	defer d.DB.Close()

	result, err := statement.Exec(value)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	UpdateID(model, uint64(lastID))

	return nil

}

func (d *Mysql) InsertMany(table string, model interface{}) error {
	phones, _ := model.(*[]models.Phone)
	x := fmt.Sprintf("insert into %s (company_id, number) values(?, ?)", table)
	statement, err := d.DB.Prepare(x)
	if err != nil {
		return err
	}
	defer d.DB.Close()
	for _, row := range *phones {
		_, err := statement.Exec(row.CompanyID, row.Number)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Mysql) QueryID(table string, value string, model interface{}) error {
	x := fmt.Sprintf("select * from %s where ID = ?", table)
	query, err := d.DB.Query(x, value)
	if err != nil {
		return nil
	}
	defer d.DB.Close()

	company, ok := model.(*models.Company)
	if !ok {
		phone, _ := model.(*models.Phone)

		for query.Next() {

			if err = query.Scan(
				&phone.ID,
				&phone.Number,
			); err != nil {
				return err
			}
		}
		return nil
	}

	for query.Next() {
		if err = query.Scan(
			&company.ID,
			&company.Name,
		); err != nil {
			return err
		}
	}

	return nil

}

func (d *Mysql) Query(table string, columnName string, values string, model interface{}) error {
	x := fmt.Sprintf("select * from %s where %s LIKE ?", table, columnName)
	query, err := d.DB.Query(x, values)
	if err != nil {
		return err
	}
	defer d.DB.Close()

	companies, ok := model.(*[]models.Company)
	if !ok {
		phones, _ := model.(*[]models.Phone)

		for query.Next() {
			phone := models.Phone{}

			if err = query.Scan(
				&phone.ID,
				&phone.Number,
			); err != nil {
				return err
			}

			*phones = append(*phones, phone)
		}

		return nil
	}

	for query.Next() {
		company := models.Company{}

		if err = query.Scan(
			&company.ID,
			&company.Name,
		); err != nil {
			return err
		}

		*companies = append(*companies, company)
	}

	return nil
}

func (d *Mysql) Update(table string, columnName string, model interface{}, ID string) error {
	value := TypeAssertion(model)

	x := fmt.Sprintf("update %s set %s = ? where ID = ?", table, columnName)
	statement, err := d.DB.Prepare(x)
	if err != nil {
		return err
	}
	defer d.DB.Close()

	_, err = statement.Exec(value, ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mysql) Delete(table string, ID string, model interface{}) error {
	x := fmt.Sprintf("delete from %s where ID = ?", table)
	statement, err := d.DB.Prepare(x)

	if err != nil {
		return err
	}
	defer d.DB.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mysql) QueryAvailable(table string, columnName string, ID string, model interface{}) error {
	x := fmt.Sprintf("select * from %s where %s = ? order by rand() limit 1", table, columnName)

	query, err := d.DB.Query(x, ID)
	if err != nil {
		return nil
	}

	company, ok := model.(*models.Company)
	if !ok {
		phone, _ := model.(*models.Phone)

		for query.Next() {

			if err = query.Scan(
				&phone.ID,
				&phone.CompanyID,
				&phone.Number,
			); err != nil {
				return err
			}
		}
		return nil
	}

	for query.Next() {
		if err = query.Scan(
			&company.ID,
			&company.Name,
		); err != nil {
			return err
		}
	}

	return nil

}

func (d *Mysql) QueryCount(table string, columnName string, model interface{}, modelGroup interface{}) error {
	x := fmt.Sprintf("select %s, count(*) AS `available_phones` from %s group by %s", columnName, table, columnName)
	query, err := d.DB.Query(x)
	if err != nil {
		return nil
	}
	defer d.DB.Close()

	phones, _ := modelGroup.(*[]models.PhoneGroup)

	for query.Next() {
		var phone models.PhoneGroup

		if err = query.Scan(
			&phone.CompanyID,
			&phone.AvailablePhones,
		); err != nil {
			return nil
		}

		*phones = append(*phones, phone)
	}

	return nil
}
