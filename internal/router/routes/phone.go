package routes

import (
	"net/http"

	"github.com/luisberga/phones-api/internal/controllers"
)

var PhonesRoute = []Route{
	{
		URI:    "/phones",
		Method: http.MethodGet,
		Func:   controllers.ListPhones,
		Auth:   true,
	},
}
