package client

import "fmt"

type CarAPI struct {
	VehicleId   string
	ServerURL   string
	BaererToken string
}

func NewCarAPI(vehicleId, serverUrl, baererToken string) (*CarAPI, error) {
	if vehicleId == "" || serverUrl == "" || baererToken == "" {
		return nil, fmt.Errorf("invalid parameters")
	}
	return &CarAPI{VehicleId: vehicleId, ServerURL: serverUrl, BaererToken: baererToken}, nil
}

func (c *CarAPI) GetBaseURL() string {
	return c.ServerURL + c.VehicleId
}
