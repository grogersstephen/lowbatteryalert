package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Type battery will hold all the information about our battery status
type battery struct {
	Charge   int  // This is the percentage of the battery
	Charging bool // Is the battery connected to power?
	// Level will tell us how low the battery is compared to the constants declared above:
	//     SAFE, LOW, VERY_LOW, and CRITICAL
	Level int
}

func (b *battery) getValues() error {
	// This method will call other methods to get the status and capacity values for the battery
	errmsg := "Cannot get Battery Info from file: "

	err := b.getStatus()
	if err != nil {
		return fmt.Errorf(errmsg + err.Error())
	}

	err = b.getCharge()
	if err != nil {
		return fmt.Errorf(errmsg + err.Error())
	}
	return nil
}

func (b *battery) getCharge() error {
	// This will get the contents of the capacity file, strip the white space, convert to int
	//    And assign the value to b.Charge
	dat, err := os.ReadFile(PATH + "capacity")
	if err != nil {
		return err
	}
	b.Charge, err = strconv.Atoi(strings.TrimSpace(string(dat)))
	return err
}

func (b *battery) getStatus() error {
	// This will get the contents of the status file, strip the white space, and assign the value to b.Status
	dat, err := os.ReadFile(PATH + "status")
	if err != nil {
		return err
	}
	if strings.TrimSpace(string(dat)) == "Charging" {
		b.Charging = true
		return nil
	}
	b.Charging = false
	return nil
}

func (b *battery) updateLevel() int {
	// Will return:
	//     +1 if we ascended a level
	//     -1 if we descended a level
	//     0 if the level is unchanged
	if b.levelDown() {
		return -1
	}
	if b.levelUp() {
		return 1
	}
	return 0
}

func (b *battery) levelUp() bool {
	if b.Charge >= THRESHOLD[LOW] && b.Level < SAFE {
		b.Level = SAFE
		return true
	}
	if b.Charge >= THRESHOLD[VERY_LOW] && b.Level < LOW {
		b.Level = LOW
		return true
	}
	if b.Charge >= THRESHOLD[CRITICAL] && b.Level < VERY_LOW {
		b.Level = VERY_LOW
		return true
	}
	return false
}

func (b *battery) levelDown() bool {
	// If we dropped a level,
	//    Sets b.Level to new level
	//    Returns true
	// Otherwise returns false
	if b.Charge < THRESHOLD[CRITICAL] && b.Level > CRITICAL {
		b.Level = CRITICAL
		return true
	}
	if b.Charge < THRESHOLD[VERY_LOW] && b.Level > VERY_LOW {
		b.Level = VERY_LOW
		return true
	}
	if b.Charge < THRESHOLD[LOW] && b.Level > LOW {
		b.Level = LOW
		return true
	}
	return false
}
