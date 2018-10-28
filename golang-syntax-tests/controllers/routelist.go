package controllers

import (
	"net/http"

	"github.com/learning-go/golang-syntax-tests/middleware"
)

type Controller func(w http.ResponseWriter, r *http.Request, data []middleware.MidData)

type Route struct {
	url        string
	controller Controller
}

var RouteList []Route

func AddRoute(url string, controller Controller) {
	RouteList = append(RouteList, Route{url: url, controller: controller})
}

func prepareResponse(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controller(w, r, middleware.ExecuteMiddleware())
	}
}

func LoadRoutes() {
	for _, route := range RouteList {
		http.HandleFunc(route.url, prepareResponse(route.controller))
	}
}
