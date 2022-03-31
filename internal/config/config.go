package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseStringConnection = ""
	ORMConnection            = ""
	Port                     = 0
	SecretKey                []byte
	DBImplem                 = ""
	_, b, _, _               = runtime.Caller(0)
	ProjectRootPath          = filepath.Join(filepath.Dir(b), "../..")
)

func Load() {

	var err error
	if err = godotenv.Load(ProjectRootPath + "/.env"); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	DatabaseStringConnection = fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	ORMConnection = fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	DBImplem = os.Getenv("DB_IMPL")
}
