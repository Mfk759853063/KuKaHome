package server

import (
	"KuKaHome/handler"
	"KuKaHome/helper"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = helper.Logger(handler, route.Name)
		router.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func TEST() {
	fmt.Println("in test")
}

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      []string{"GET"},
		Pattern:     "/",
		HandlerFunc: handler.Index,
	},
	Route{
		Name:        "productlist",
		Method:      []string{"GET"},
		Pattern:     "/productlist",
		HandlerFunc: handler.Productlist,
	},
	Route{
		Name:        "login",
		Method:      []string{"GET"},
		Pattern:     "/login",
		HandlerFunc: handler.Login,
	},
	Route{
		Name:        "uploadFile",
		Method:      []string{"POST"},
		Pattern:     "/uploadFile",
		HandlerFunc: handler.UploadFile,
	},
}
