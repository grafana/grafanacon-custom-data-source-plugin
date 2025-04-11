# Grafanacon Custom Data Source Plugin 

Your one stop shop for all resources you'll need in our hands on lab on how to build a custom data source plugin!

## Prerequisites
- Docker
- Go 1.22
- Mage
- LTS version of Node.js
- Git

Not sure if you have one of these installed? Ask for help!  

## Getting setup (DO THIS NOW)

Scaffold your plugin:
1. Find a place on your file system where you'd like to create a new folder for your plugin
2. Run create plugin (creates it's own folder): `npx @grafana/create-plugin@latest`
3. Follow the wizard with the following answers:
   - Select a plugin type: data source
   - Add a backend to support server-side functionality?: y
   - Enter a name for your plugin: test
   - Enter your organization name: my-test-org
4. `cd ./mytestorg-test-datasource`
5. `npm install`
6. `npm run dev` (keep this running throughout the workshop)
7. in a separate tab: `mage -v` (or specify the build for your machine to speed it up ex: `mage -v build:linuxARM64`)
8. `docker compose up`
9. open http://localhost:3000

Note: Hotel wifi can be spotty and `npm install` can take a while to run. If npm install is taking forever consider using your phone as a hotspot temporarily! Or wait a few minutes and try again! Or if your neighbor's is working, consider pairing together. 

## Our Target API

A grafana data source plugin brings data from different apis and into grafana. Our project for today will be to build a data source plugin that connects to [Seattle Open Data's Road Weather Station Data](https://data.seattle.gov/Transportation/Road-Weather-Information-Stations/egc4-d24i/about_data). As hotel internet bandwidth can sometimes be in limited supply, we've duplicated their api and data to run locally in a very small no dependency go server. For the purposes of this workshop, we encourage you to run it on your machine. However after the conference, it should be easy to edit your configuration settings to connect to a [live endpoint](#live-data).

## Running the Target API Locally 

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
2. In your datasource configuration replace http://localhost:8080 with https://data.seattle.gov/ and add your token.

So that requests go from:
http://localhost:8080/resource/egc4-d24i.json?$$app_token=good&stationname=SpokaneSwingBridge

to this:
https://data.seattle.gov/resource/egc4-d24i.json?$$app_token={a real app token you get from data.seattle.gov}&stationname=SpokaneSwingBridge

Explore the api docs here: https://dev.socrata.com/foundry/data.seattle.gov/egc4-d24i

There's lots of ways to extend your plugin! Some things to try after this workshop:
- adding support for geolocation coordinates
- supporting [SoQL queries](https://dev.socrata.com/docs/queries/)
- connecting to other datasets (resources beyond egc4-d24i)
