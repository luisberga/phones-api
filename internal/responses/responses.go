package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type DefaultQuery struct {
	ID            uint64     `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	CompanyID     uint64     `json:"company_id,omitempty"`
	Number        string     `json:"phone_number,omitempty"`
	Rows          [][]string `json:"rows,omitempty"`
	PhoneQuantity uint64     `json:"available_phones,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if statusCode != 204 {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Err(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"err"`
	}{
		Error: err.Error(),
	})
}
