package models

type DatabaseInterface interface {
	Connect() error
	Insert(table string, columnName string, model interface{}) error
	Query(table string, columnName string, value string, model interface{}) error
	QueryID(table string, ID string, model interface{}) error
	Update(table string, columnName string, model interface{}, ID string) error
	Delete(table string, ID string, model interface{}) error
	QueryAvailable(table string, columnName string, ID string, model interface{}) error
	InsertMany(table string, model interface{}) error
	QueryCount(table string, columnName string, model interface{}, models interface{}) error
}
