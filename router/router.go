package router

import (
	"net/http"
	"github.com/ConcesionarioBack/actions"
	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		especificRoute := router.Methods(route.Method, "OPTIONS").Path(route.Pattern).Subrouter()
		especificRoute.Name(route.Name).Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes {
	Route{
		"Healthy",
		"GET",
		"/api/healthy",
		actions.Healthy,
	},
	Route {
		"autos",
		"GET",
		"/api/autos",
		actions.GetAutos,
	},
	Route {
		"getAuto",
		"GET",
		"/api/autos/{id}",
		actions.GetAuto,
	},
	Route{
		"AgregarUsuario",
		"POST",
		"/api/usuario",
		actions.AgregarUsuario,
	},
}