package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type AddedNumbersResponse struct {
	AddedNumbers int
}

var routes = Routes{

	Route{
		"GetRequest", // Name
		"GET",        // HTTP method
		"/ping",      // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		},
	},

	Route{
		"Health",  // Name
		"GET",     // HTTP method
		"/health", // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	},

	Route{
		"Readiness",  // Name
		"GET",        // HTTP method
		"/readiness", // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	},

	Route{
		"GetRequest", // Name
		"GET",        // HTTP method
		"/demonstrate-addition/{numberOne}/{numberTwo}", // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")

			params := strings.Split(r.URL.Path, "/")

			numberOne, errNumberOne := strconv.Atoi(params[2])
			numberTwo, errNumberTwo := strconv.Atoi(params[3])
			addedNumbers := AddedNumbersResponse{AddedNumbers: numberOne + numberTwo}

			js, _ := json.Marshal(addedNumbers)

			if errNumberOne != nil || errNumberTwo != nil {
				return
			}

			// Log json object server side
			prettyJs, _ := json.MarshalIndent(addedNumbers, "", "    ")
			fmt.Println(string(prettyJs))

			// Send response
			w.Write(js)

		},
	},
}
