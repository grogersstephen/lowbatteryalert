package main

import (
	"os/exec"
)

// Type notification will be a wrapper for a map[string]string
//     The keys will be the flags used in the call to twmnc
//     The values will be the arguments for each flag
type notification map[string]string

func (n *notification) notify() error {
	// This will execute the twmnc command, launching our notification with the values in n.Arguments
	var args []string // args will hold the arguments to our twmnc command as a []string
	for k, v := range *n {
		if v == "" { // If an argument is not set, we won't include it
			continue
		}
		args = append(args, k, v)
	}
	cmd := exec.Command("twmnc", args...)
	err := cmd.Run()
	return err
}
