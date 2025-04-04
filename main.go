package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type StationData []struct {
	Stationname     string `json:"stationname"`
	Stationlocation struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"stationlocation"`
	Datetime               string `json:"datetime"`
	Recordid               string `json:"recordid"`
	Roadsurfacetemperature string `json:"roadsurfacetemperature"`
	Airtemperature         string `json:"airtemperature"`
}

func main() {
	// Serve the static JSON file only at the specific path with correct token
	http.HandleFunc("/resource/egc4-d24i.json", func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct query parameter
		token := r.URL.Query().Get("$$app_token")
		if token != "good" {
			http.Error(w, "Forbidden: Invalid or missing token", http.StatusForbidden)
			return
		}

		// Read the JSON file
		data, err := os.ReadFile("data.json")
		if err != nil {
			log.Printf("Error reading data file: %v", err)
			http.Error(w, "Error reading data file", http.StatusInternalServerError)
			return
		}

		// Parse the JSON
		var stationData StationData
		if err := json.Unmarshal(data, &stationData); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			log.Printf("Raw JSON data: %s", string(data))
			http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
			return
		}

		// Check for stationname filter
		stationName := r.URL.Query().Get("stationname")
		if stationName != "" {
			// Filter stations by name
			var filteredStations StationData
			for _, station := range stationData {
				if station.Stationname == stationName {
					filteredStations = append(filteredStations, station)
				}
			}
			stationData = filteredStations
		}

		// Set content type and send response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(stationData); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	// Handle root path with a friendly message
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("the server for your target api is up and running"))
	})

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	log.Println("JSON available at: http://localhost:8080/resource/egc4-d24i.json?$$app_token=good")
	log.Println("Filter by station: http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=JoseRizalBridgeNorth")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
