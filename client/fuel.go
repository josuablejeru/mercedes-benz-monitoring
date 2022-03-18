package client

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *CarAPI) GetFuelLevel(ctx context.Context) (string, error) {
	url := c.GetBaseURL() + "/containers/fuelstatus"
	bearer := "Bearer " + c.BaererToken

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		log.Fatal(err)
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body), nil
}
