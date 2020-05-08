package main

import (
	"fmt"
	"log"

	"github.com/morpheusnephew/goflightservice/internal/flights"
	"github.com/morpheusnephew/goflightservice/internal/schedule"
)

func main() {

	schedule.ConfigureSchedule(flightRunner)

	for {

	}

}

func flightRunner() {
	flightInformation, err := flights.GetAvailableFlights()

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(flightInformation)
	}
}
