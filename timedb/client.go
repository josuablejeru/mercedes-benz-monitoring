package timedb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type DBclient struct {
	client    influxdb2.Client
	serverURL string
}

func New(serverURL, token string) (*DBclient, error) {
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	return &DBclient{client: influxdb2.NewClient(serverURL, token), serverURL: serverURL}, nil
}
