package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Type battery will hold all the information about our battery status
type battery struct {
	// Charge   int  // This is the percentage of the battery
	// Charging bool // Is the battery connected to power?
	// Level will tell us how low the battery is compared to the constants declared above:
	//     SAFE, LOW, VERY_LOW, and CRITICAL
	// FIXES: now that you have it as a type alias it is easier to jump to the definition of the enums
	Level threshold
}

func (b *battery) getValues() (bool, int, error) {
	// This method will call other methods to get the status and capacity values for the battery
	// FIXES: don't put capitals in go errors, but also don't do errors like this. it can be hard to track down.
	errmsg := "cannot get battery info from file: %w"
	// NOTE: i'm doing this to beable to leave the variable and not use it. just for note purposes
	_ = errmsg

	isCharging, err := b.isCharging()
	if err != nil {
		// FIXES: in go you want to beable to unwrap errors, to dont concat the strings together.
		return false, 0, fmt.Errorf("is charging: %w", err)
	}

	charge, err := b.getCharge()
	if err != nil {
		return false, 0, fmt.Errorf("get charge: %w", err)
	}
	return isCharging, charge, nil
}

// NOTE: i'm not a big fan of this style of OOP stuff. i.e., the method that sets a value on the class/struct
// if it is just a functional read. it should just be returned. random functions that set values on the parent struct and lead to "side effects" or "symtopms" which are just things that are hard to track down at scale
// it is not wrong here. just my opinion.
func (b *battery) getCharge() (int, error) {
	// This will get the contents of the capacity file, strip the white space, convert to int
	//    And assign the value to b.Charge
	// FIXES: I like the habit of using filepath
	dat, err := os.ReadFile(filepath.Join(PATH, "capacity"))
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(dat)))
}

func (b *battery) isCharging() (bool, error) {
	// This will get the contents of the status file, strip the white space, and assign the value to b.Status
	dat, err := os.ReadFile(filepath.Join(PATH, "status"))
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(string(dat)) == "Charging" {
		// b.Charging = true
		return true, nil
	}
	// b.Charging = false
	return false, nil
}

func (b *battery) updateLevel(charge int) int {
	// Will return:
	//     +1 if we ascended a level
	//     -1 if we ascended a level
	//     0 if the level is unchanged
	if b.levelDown(charge) {
		return -1
	}
	if b.levelUp(charge) {
		return 1
	}
	return 0
}

<<<<<<< HEAD
func (b *battery) levelUp(charge int) bool {
	if charge >= THRESHOLD[CRITICAL] && b.Level < VERY_LOW {
		b.Level = VERY_LOW
=======
func (b *battery) levelUp() bool {
	if b.Charge >= THRESHOLD[LOW] && b.Level < SAFE {
		b.Level = SAFE
>>>>>>> 60f01bf (reversed the conditional statements in levelUp())
		return true
	}
	if charge >= THRESHOLD[VERY_LOW] && b.Level < LOW {
		b.Level = LOW
		return true
	}
<<<<<<< HEAD
	if charge >= THRESHOLD[LOW] && b.Level < SAFE {
		b.Level = SAFE
=======
	if b.Charge >= THRESHOLD[CRITICAL] && b.Level < VERY_LOW {
		b.Level = VERY_LOW
>>>>>>> 60f01bf (reversed the conditional statements in levelUp())
		return true
	}
	return false
}

func (b *battery) levelDown(charge int) bool {
	// If we dropped a level,
	//    Sets b.Level to new level
	//    Returns true
	// Otherwise returns false
	if charge < THRESHOLD[CRITICAL] && b.Level > CRITICAL {
		b.Level = CRITICAL
		return true
	}
	if charge < THRESHOLD[VERY_LOW] && b.Level > VERY_LOW {
		b.Level = VERY_LOW
		return true
	}
	if charge < THRESHOLD[LOW] && b.Level > LOW {
		b.Level = LOW
		return true
	}
	return false
}
