package main

import (
	"machine"
	"time"
)

const (
	fc1 = machine.D0
	fc2 = machine.D1
	fc3 = machine.D2
	fc4 = machine.D3
)

func main() {
	for _, fc := range []machine.Pin{fc1, fc2, fc3, fc4} {
		fc.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
	for {
		if fc1.Get() {
			println("pushing: fc1")
		} else if fc2.Get() {
			println("pushing: fc2")
		} else if fc3.Get() {
			println("pushing: fc3")
		} else if fc4.Get() {
			println("pushing: fc4")
		} else {
			// nothing to do
		}
		time.Sleep(time.Millisecond * 10)
	}
}
