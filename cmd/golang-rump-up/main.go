package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisberga/phones-api/internal/config"
	"github.com/luisberga/phones-api/internal/router"
)

func main() {
	config.Load()
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
