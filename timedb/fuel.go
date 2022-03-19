package timedb

import (
	"fmt"
)

func (db *DBclient) WriteFuelLevel(value int64) {
	writeAPI := db.client.WriteAPI("primary", "mercedes-car")

	fmt.Println("Writing fuel level:", value)
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=fuelLevel avg=%d", value))
	writeAPI.Flush()
}
