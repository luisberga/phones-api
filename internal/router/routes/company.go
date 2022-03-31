package routes

import (
	"net/http"

	"github.com/luisberga/phones-api/internal/controllers"
)

var CompaniesRoute = []Route{
	{
		URI:    "/companies",
		Method: http.MethodPost,
		Func:   controllers.CreateCompany,
		Auth:   true,
	},
	{
		URI:    "/companies",
		Method: http.MethodGet,
		Func:   controllers.ListCompanies,
		Auth:   true,
	},

	{
		URI:    "/companies/{id}",
		Method: http.MethodGet,
		Func:   controllers.ListCompany,
		Auth:   true,
	},
	{
		URI:    "/companies/{id}",
		Method: http.MethodPut,
		Func:   controllers.UpdateCompany,
		Auth:   true,
	},
	{
		URI:    "/companies/{id}",
		Method: http.MethodDelete,
		Func:   controllers.DeleteCompany,
		Auth:   true,
	},

	{
		URI:    "/companies/{id}/phone",
		Method: http.MethodGet,
		Func:   controllers.AvailablePhones,
		Auth:   true,
	},
	{
		URI:    "/companies/{id}/phones",
		Method: http.MethodPost,
		Func:   controllers.CreatePhones,
		Auth:   true,
	},
}
