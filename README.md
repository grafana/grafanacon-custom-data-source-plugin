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
3. In your browser go to http://localhost:8080

## Features

- Serves a static JSON file (`data.json`) at `/resource/egc4-d24i.json` with valid token: example: http://localhost:8080/resource/egc4-d24i.json?$$app_token=good
- Requires `$$app_token=good` query parameter, can test failing auth at http://localhost:8080/resource/egc4-d24i.json?$$app_token=bad
- Optional `stationname` parameter to filter data by station: for example:  http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=JoseRizalBridgeNorth
- Returns 403 Forbidden for invalid or missing tokens
- Returns 404 for all other routes

## Live Data (after workshop)
Once you have better internet access, you may wish to update your plugin to use a real API endpoint rather than this local dupe. To do so: 

1. [Sign up for an access token](https://data.seattle.gov/signup)
2. In your plugin code replace http://localhost:8080 with https://data.seattle.gov/
3. In your datasource replace your app token for a real one. 

So that requests go from:
http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=SpokaneSwingBridge

to this:
https://data.seattle.gov/resource/egc4-d24i.json?$$app_token={a real app token you get from data.seattle.gov}&stationname=SpokaneSwingBridge

Explore the api docs here: https://dev.socrata.com/foundry/data.seattle.gov/egc4-d24i
