package main

import (
	"context"
	"fmt"
	"log"

	"github.com/josuablejeru/mercedes-benz-monitoring/client"
)

const (
	vehicleId = "WDB111111ZZZ22222"                                             // example ID
	serverURL = "https://api.mercedes-benz.com/vehicledata_tryout/v2/vehicles/" // Playground URL
)

func main() {
	ctx := context.Background()

	// Create a new client
	carApi, err := client.NewCarAPI(vehicleId, serverURL, "7c7c777c-f123-4123-s123-7c7c777c7c77")
	if err != nil {
		log.Fatal(err)
	}

	result, err := carApi.GetFuelLevel(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result[1].Rangeliquid.Value)
}
