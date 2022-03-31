package controllers

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luisberga/phones-api/internal/middleware"
	"github.com/luisberga/phones-api/internal/models"
	"github.com/luisberga/phones-api/internal/repository"
	"github.com/luisberga/phones-api/internal/responses"
)

func AvailablePhones(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]

	var phone models.Phone

	database := middleware.DatabaseChoose()
	phoneRepository := repository.NewPhoneRepo(database)

	err := phoneRepository.ListAvailable(&phone, ID)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, phone)
}

func CreatePhones(w http.ResponseWriter, r *http.Request) {
	var phones []models.Phone

	ID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer file.Close()

	reads := csv.NewReader(file)
	rows, err := reads.ReadAll()
	if err != nil || len(rows) > 5000000 {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	for _, row := range rows {
		phone := models.Phone{CompanyID: ID, Number: row[0]}
		phones = append(phones, phone)
	}

	database := middleware.DatabaseChoose()
	phoneRepository := repository.NewPhoneRepo(database)

	err = phoneRepository.Create(&phones)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, "file uploaded")
}

func ListPhones(w http.ResponseWriter, r *http.Request) {

	database := middleware.DatabaseChoose()
	phoneRepository := repository.NewPhoneRepo(database)
	var phone models.Phone
	var phonesGroup []models.PhoneGroup

	err := phoneRepository.List(&phone, &phonesGroup)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, phonesGroup)
}
