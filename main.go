package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	mbclient "github.com/josuablejeru/mercedes-benz-monitoring/mbclient"
	"github.com/josuablejeru/mercedes-benz-monitoring/timedb"
)

var (
	influxToken     string
	vehicleId       string
	mbServerURL     string
	timeDBServerURL string
)

func main() {
	loadEnv()
	ctx := context.Background()

	influxToken = os.Getenv("INFLUXDB_BUCKET_TOKEN")
	vehicleId = os.Getenv("VEHICLE_ID")
	mbServerURL = os.Getenv("SERVER_URL")
	timeDBServerURL = os.Getenv("INFLUXDB_SERVER_URL")

	// Create a new client
	carApi, err := mbclient.NewCarAPI(vehicleId, mbServerURL, "7c7c777c-f123-4123-s123-7c7c777c7c77")
	if err != nil {
		log.Fatal(err)
	}

	result, err := carApi.GetFuelLevel(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Write to TimeDB
	timeclient, _ := timedb.New(timeDBServerURL, influxToken)
	intValue, _ := strconv.ParseInt(result[1].Rangeliquid.Value, 10, 64)
	timeclient.WriteFuelLevel(intValue)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
