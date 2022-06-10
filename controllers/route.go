package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	common := negroni.New()

	authMid := common.With(AuthMiddleware())
	router.Methods(http.MethodPost).Path("/signup").HandlerFunc(Signup)
	router.Methods(http.MethodPost).Path("/login").HandlerFunc(Login)
	router.Methods(http.MethodPost).Path("/validate").HandlerFunc(Validate)

	// subrouter for buy
	buySubroute := mux.NewRouter().PathPrefix("/buy").Subrouter().StrictSlash(true)
	buySubroute.Methods(http.MethodGet).HandlerFunc(Buy)
	router.PathPrefix("/buy").Handler(authMid.With(negroni.Wrap(buySubroute)))

	// subrouter for sell coin
	sellSubroute := mux.NewRouter().PathPrefix("/sell").Subrouter().StrictSlash(true)
	sellSubroute.Methods(http.MethodGet).HandlerFunc(Sell)
	router.PathPrefix("/sell").Handler(authMid.With(negroni.Wrap(sellSubroute)))

	return router
}

// type midSvc struct{}

// func NewMiddleware() *midSvc {
// 	return &midSvc{}
// }
// func (mid *midSvc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	// do some stuff before
// 	next(rw, r)
// 	// do some stuff after
// }

// func ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	// do some stuff before
// 	next(rw, r)
// 	// do some stuff after
// }
