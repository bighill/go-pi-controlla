package main

import "fmt"

func wsTrigger(trigger string) {
    if trigger == "un" {
        trigger_un()
    }
    if trigger == "deux" {
        trigger_deux()
    }
    if trigger == "trois" {
        trigger_trois()
    }

    pi(trigger)
}

func trigger_un() {
    fmt.Println("un was triggered")
}

func trigger_deux() {
    fmt.Println("deux was triggered")
}

func trigger_trois() {
    fmt.Println("trois was triggered")
}
