# Grafanacon Custom Data Source Plugin 

Your one stop shop for all resources you'll need in our hands on lab on how to build a custom data source plugin! This repo contains a local version of the api we plan to hit today. It will run on `localhost:8080`. It will also serve all of documentation we'll need as well as references to git commits to [the fully built version of the plugin](https://github.com/grafana/grafanacon-custom-data-source-plugin-example).  

## Table of Contents
- [Prerequisites](#prerequisites)
- [Setup (DO THIS NOW)](#setup-do-this-now)
- [Our Target API: Seattle Open Data -- Road Weather Station](#our-target-api-seattle-open-data----road-weather-station)
  - [Running the Target API Locally](#running-the-target-api-locally)
  - [Target API Features](#target-api-features)
- [Live Data (after workshop)](#live-data-after-workshop)


## Prerequisites
- Docker
- Go 1.22
- Mage
- LTS version of Node.js
- Git

Not sure if you have one of these installed? Ask for help!  

## Setup (DO THIS NOW)

To Scaffold your plugin:
1. Find a place on your file system where you'd like to create a new folder for your plugin
2. `npx @grafana/create-plugin@latest` (creates it's own folder, downloads template, and starts wizard):
   - Select a plugin type: data source
   - Add a backend to support server-side functionality?: y
   - Enter a name for your plugin: test
   - Enter your organization name: my-test-org
3. `cd ./mytestorg-test-datasource`
4. `npm install`
5. `npm run dev` (keep this running in a separate tab throughout the workshop, it will watch frontend files and compile them)
7. `mage -v && docker compose up` (or specify the build for your machine to speed it up ex: `mage -v build:linuxARM64`)
9. open http://localhost:3000

Note: Hotel wifi can be spotty and `npm install` can take a while to run. If npm install is taking forever consider using your phone as a hotspot temporarily! Or wait a few minutes and try again! Or if your neighbor's is working, consider pairing together. 

## Our Target API: Seattle Open Data -- Road Weather Station

A grafana data source plugin queries data from an external api and then visualizes it in grafana. Our project for today will be to build a data source plugin that connects to [Seattle Open Data's Road Weather Station Data](https://data.seattle.gov/Transportation/Road-Weather-Information-Stations/egc4-d24i/about_data). As internet bandwidth can sometimes be in limited supply at conferences, we've duplicated their api and data to run locally in a very small no dependency go server. For the purposes of this workshop, we encourage you to run it on your machine. However after this workshop, it should be easy to edit your configuration settings to connect to a [live endpoint](#live-data).

### Running the Target API Locally 

1. Clone this repository
   ```
   git clone https://github.com/grafana/grafanacon-custom-data-source-plugin.git
   ```
2. `cd grafanacon-custom-data-source-plugin`
3. Run the server:
   ```bash
   go run main.go
   ```
4. In your browser go to `http://localhost:8080`

### Target API Features

- Serves a static JSON file (`data.json`) at `/resource/egc4-d24i.json` with valid token: example: http://localhost:8080/resource/egc4-d24i.json?$$app_token=good
- Requires `$$app_token=good` query parameter, can test failing auth at http://localhost:8080/resource/egc4-d24i.json?$$app_token=bad
- Optional `stationname` parameter to filter data by station: for example:  http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=JoseRizalBridgeNorth
- Returns 403 Forbidden for invalid or missing tokens
- Returns 404 for all other routes

## Live Data (after workshop)
Once you have better internet access, you may wish to update your plugin to use a real API endpoint rather than this local dupe. To do so: 

1. [Sign up for an access token](https://data.seattle.gov/signup)
2. In your datasource configuration change your domain to with https://data.seattle.gov/ and swap out "good" for your app token.

So that requests go from:

`http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=SpokaneSwingBridge`

to this:

`https://data.seattle.gov/resource/egc4-d24i.json?$$app_token={a real app token you get from data.seattle.gov}&stationname=SpokaneSwingBridge`

Explore the api docs here: https://dev.socrata.com/foundry/data.seattle.gov/egc4-d24i

There are lots of ways to extend your plugin after this talk! Some things to try after this workshop:
- adding support for geolocation coordinates
- supporting [SoQL queries](https://dev.socrata.com/docs/queries/)
- connecting to other datasets (resources beyond egc4-d24i)
- building out query vs builder mode
- practice writing e2e tests withour testing framework
