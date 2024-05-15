package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GetGas(w http.ResponseWriter, _ *http.Request) {
	results := make([]string, 0)
	for x := range GASSES {
		results = append(results, x)
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Fatal(err)
	}
}

func AddGas(w http.ResponseWriter, r *http.Request) {
	type AddGasMessage struct {
		GAS string `json:"gas"`
	}
	gaM := new(AddGasMessage)

	err := json.NewDecoder(r.Body).Decode(gaM)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}

	gaM.GAS = strings.ToLower(gaM.GAS)
	GASSES[gaM.GAS] = new(GAS)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("ok"))
	if err != nil {
		log.Panic(err)
	}
}

func AddResult(w http.ResponseWriter, r *http.Request) {
	type AddMessage struct {
		GAS   string `json:"gas"`
		KEY   string `json:"key"`
		VALUE string `json:"value"`
	}
	aM := new(AddMessage)

	err := json.NewDecoder(r.Body).Decode(aM)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}

	aM.KEY = strings.ToLower(aM.KEY)
	aM.GAS = strings.ToLower(aM.GAS)
	aM.VALUE = strings.ToLower(aM.VALUE)

	gas, ok := GASSES[aM.GAS]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}

	gas.AddResult(aM.KEY, aM.VALUE)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("ok"))
	if err != nil {
		log.Panic(err)
	}
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	gasQ := strings.ToLower(r.URL.Query().Get("gas"))
	keyQ := strings.ToLower(r.URL.Query().Get("key"))

	if gasQ == "" || keyQ == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Missing gas or key, GAS: %s, Key: %s", gasQ, keyQ)
		return
	}

	gas, ok := GASSES[gasQ]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusNotFound)
		log.Println("GAS Not Found")
		return
	}

	results := gas.GetResults(keyQ)
	msg, err := json.Marshal(results)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(msg)
	if err != nil {
		log.Panic(err)
	}
}
