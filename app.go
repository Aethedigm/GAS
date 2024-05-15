package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"main/gas"
	"net/http"
	"strconv"
)

const PORT int = 6605

var GASSES map[string]*gas.GAS

func main() {
	GASSES = make(map[string]*gas.GAS)

	Serve()
}

func Serve() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Get("/", GetResults)
	router.Post("/", AddResult)
	router.Get("/gas", GetGas)
	router.Post("/gas", AddGas)

	log.Println("Starting server...")
	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), router); err != nil {
		log.Fatal(err)
	}
}
