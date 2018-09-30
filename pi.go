/*

A blinker example using go-rpio library.
Requires administrator rights to run

Toggles a LED on physical pin PIN 
Connect a LED with resistor from pin PIN to ground.

*/

package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
)

const (
    PIN18 = rpio.Pin(18)
    PIN23 = rpio.Pin(23)
)

func pi(_action string) {
    fmt.Println(_action)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	PIN18.Output()
	PIN23.Output()

    if _action == "eighteen__on" {
        PIN18.High()
    }
    if _action == "eighteen__off" {
        PIN18.Low()
    }

    if _action == "twentythree__on" {
        PIN23.High()
    }
    if _action == "twentythree__off" {
        PIN23.Low()
    }

}
