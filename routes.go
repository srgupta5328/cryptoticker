package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	APIURL      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (conf *Config) NewRouter() *mux.Router {
	var marketRoutes = Routes{
		Route{
			"GetListing",
			"GET",
			"/api/listing",
			conf.ListingHandler,
		},
		Route{
			"Global",
			"GET",
			"/api/global",
			conf.GlobalHandler,
		},
		Route{
			"Ticker",
			"GET",
			"/api/ticker",
			conf.TickerHandler,
		},

		Route{
			"TargetTicket",
			"GET",
			"/api/ticker/{id:[0-1610]+}",
			conf.TargetTickerHandler,
		},
		Route{
			"Index",
			"GET",
			"/home",
			indexHandler,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, marketRoute := range marketRoutes {
		router.
			Methods(marketRoute.Method).
			Path(marketRoute.APIURL).
			Name(marketRoute.Name).
			Handler(marketRoute.HandlerFunc)
	}
	return router
}
