package main

import (
	"fmt"
	"log"
	"time"
)

const (
	PATH     = "/sys/class/power_supply/BAT0/" // Battery information path
	CRITICAL = iota                            // Levels for the battery
	VERY_LOW
	LOW
	SAFE
)

var THRESHOLD = map[int]int{ // Battery Percentage thresholds that define the levels
	CRITICAL: 5,
	VERY_LOW: 10,
	LOW:      20,
}

func main() {
	// Declare our battery struct
	b := battery{
		Level: SAFE,
	}

	// Main loop
	for {

		// Get battery status from file
		err := b.getValues()
		if err != nil {
			n := notification{
				"-t":   "ERROR",
				"-c":   err.Error(),
				"--fs": "12",
			}
			e := n.notify()
			if e != nil {
				log.Fatal(err)
			}
		}

		// updateLevel checks to see if the battery level drops below or rises above a defined threshold
		if b.updateLevel() == -1 && b.Charging == false { // Battery level dropped below a threshold
			n := notification{
				"-t":    "LOW BATTERY",
				"-c":    fmt.Sprintf("Your battery is currently at %d%%. Please provide power.", b.Charge),
				"-s":    "50",
				"-d":    "5000",
				"--fs":  "30",
				"--fg":  "red",
				"--pos": "bottom_right",
			}
			err := n.notify()
			if err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(time.Minute)
	}
}
