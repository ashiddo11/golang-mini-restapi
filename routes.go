package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetPeople",
		"GET",
		"/people",
		GetPeopleEndpoint,
	},
	Route{
		"GetPerson",
		"GET",
		"/people/{id}",
		GetPersonEndpoint,
	},
	Route{
		"CreatePerson",
		"POST",
		"/todos/{id}",
		CreatePersonEndpoint,
	},
	Route{
		"TodoShow",
		"DELETE",
		"/people/{id}",
		DeletePersonEndpoint,
	},
}