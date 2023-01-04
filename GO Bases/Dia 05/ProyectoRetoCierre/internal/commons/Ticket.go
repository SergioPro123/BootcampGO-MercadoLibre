package commons

import "time"

type Ticket struct {
	Id                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         time.Time
	Price              float64
}
