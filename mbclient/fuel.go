package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetFuelLevelResp []struct {
	Tanklevelpercent struct {
		Value     string `json:"value"`
		Timestamp int64  `json:"timestamp"`
	} `json:"tanklevelpercent,omitempty"`
	Rangeliquid struct {
		Value     string `json:"value"`
		Timestamp int64  `json:"timestamp"`
	} `json:"rangeliquid,omitempty"`
}

func (c *CarAPI) GetFuelLevel(ctx context.Context) (GetFuelLevelResp, error) {
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

	d, _ := ioutil.ReadAll(resp.Body)
	returnData := GetFuelLevelResp{}
	json.Unmarshal(d, &returnData)

	return returnData, nil
}
