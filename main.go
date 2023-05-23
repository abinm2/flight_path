package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type response struct {
	path []string
	code int
}

// CalculateFlightPath calculates the flight path given a list of flights
func CalculateFlightPath(flights [][]string) (response, error) {

	if len(flights) == 0 {
		return response{nil, http.StatusBadRequest}, errors.New("flight details not found")
	}

	if len(flights) == 1 {
		return response{flights[0], http.StatusOK}, nil
	}

	// Create a map to store the source and destination airports
	dmap := make(map[string]string)
	smap := make(map[string]string)

	// Iterate through the flights and update the airports map
	for _, flight := range flights {
		dmap[flight[1]] = flight[0]
		smap[flight[0]] = flight[1]
	}

	// Find the starting airport
	var startAirport string
	for _, dest := range dmap {
		if _, ok := dmap[dest]; !ok {
			startAirport = dest
			break
		}
	}
	if startAirport == "" {
		return response{nil, http.StatusBadRequest}, errors.New("invalid input")
	}
	// Find the starting airport
	var destairport string
	for _, dst := range smap {
		if _, ok := smap[dst]; !ok {
			destairport = dst
			break
		}
	}
	if destairport == "" {
		return response{nil, http.StatusBadRequest}, errors.New("invalid input")
	}

	return response{[]string{startAirport, destairport}, http.StatusOK}, nil
}

func calculateFlightPathHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Parse the request body
	var flights [][]string
	err := json.NewDecoder(r.Body).Decode(&flights)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Calculate the flight path
	flightPath, err := CalculateFlightPath(flights)
	if err != nil {
		w.WriteHeader(flightPath.code)
		w.Write([]byte(err.Error()))
	}
	// Convert the flight path to JSON
	response, err := json.Marshal(flightPath.path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func main() {
	http.HandleFunc("/calculate", calculateFlightPathHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
