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
	"time"
)

const (
    PIN = rpio.Pin(18)
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
	PIN.Output()

    if _action == "blink" {
        PIN.Low()
        for x := 0; x < 6; x++ {
            PIN.Toggle()
            time.Sleep(time.Second / 2)
        }
    }
    if _action == "on" {
        PIN.High()
    }
    if _action == "off" {
        PIN.Low()
    }
}
