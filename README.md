# Grafanacon Seattle Open Data 

A basic Go web server that serves a static JSON file with filtering capabilities.

It is based on real data from [Seattle Open Data's Road Weather Station Data](https://data.seattle.gov/Transportation/Road-Weather-Information-Stations/egc4-d24i/about_data). It was created to use in demo purposes at GrafanaCon during our workshops, where hotel internet bandwidth may be in limited supply.

However after the conference if you'd like to access fresher/larger datasets you should be able to easily swap out the localhost endpoint with a [live one](#live-data).

## Prerequisites

- Go 1.16 or later

## Running Locally 

1. Clone this repository
2. Run the server:
   ```bash
   go run main.go
   ```
3. Access the JSON file at: http://localhost:8080/resource/egc4-d24i.json?$$app_token=good
4. Filter by station: http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=JoseRizalBridgeNorth

## Features

- Serves a static JSON file (`data.json`) at `/resource/egc4-d24i.json` with valid token
- Requires `$$app_token=good` query parameter
- Optional `stationname` parameter to filter data by station
- Returns 403 Forbidden for invalid or missing tokens
- Returns 404 for all other routes
- Runs on port 8080
- Simple and lightweight implementation

## Live Data
Once you are in a place with live data, you may wish to update your project to use real data. 

1. [Sign up for an access token](https://data.seattle.gov/signup)
2. In your plugin code replace http://localhost:8080 with https://data.seattle.gov/
3. In your datasource replace your app token for a real one. 

So that requests go from:
http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=SpokaneSwingBridge

to this:
https://data.seattle.gov/resource/egc4-d24i.json?$$app_token={a real app token you get from data.seattle.gov}&stationname=SpokaneSwingBridge

Explore the api docs here: https://dev.socrata.com/foundry/data.seattle.gov/egc4-d24i
