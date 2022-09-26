package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ConcesionarioBack/database"
	"github.com/ConcesionarioBack/model"
	"github.com/gorilla/mux"
)

func Healthy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}

func GetAutos(w http.ResponseWriter, r *http.Request) {
	autos := database.GetAutos()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(autos)
}

func GetAuto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strVar := params["id"]
	intVal, err := strconv.Atoi(strVar)
	if err != nil {
		log.Panic(err)
	}
	auto := database.GetAuto(intVal)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(auto)
}

func AgregarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := database.AddUsuario(usuario)
	if err != nil {
		log.Panic(err)
		w.WriteHeader(400)
	} else {
		fmt.Println("Se agrego",usuario)
		w.WriteHeader(200)
	}
	json.NewEncoder(w).Encode(usuario)
}
