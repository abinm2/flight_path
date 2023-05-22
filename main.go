package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// CalculateFlightPath calculates the flight path given a list of flights
func CalculateFlightPath(flights [][]string) []string {

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

	// Find the starting airport
	var destairport string
	for _, dst := range smap {
		if _, ok := smap[dst]; !ok {
			destairport = dst
			break
		}
	}

	return []string{startAirport, destairport}
}

func calculateFlightPathHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var flights [][]string
	err := json.NewDecoder(r.Body).Decode(&flights)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Calculate the flight path
	flightPath := CalculateFlightPath(flights)

	// Convert the flight path to JSON
	response, err := json.Marshal(flightPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/calculate", calculateFlightPathHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
