package main

import (
	"machine"
	"machine/usb/hid/keyboard"
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
	kb := keyboard.New()
	wait := 1 * time.Millisecond
	for {
		if fc1.Get() {
			println("pushing: fc1")
			kb.Down(keyboard.KeyRight)
		} else if fc2.Get() {
			println("pushing: fc2")
			kb.Down(keyboard.KeyUp)
		} else if fc3.Get() {
			println("pushing: fc3")
			kb.Down(keyboard.KeyDown)
		} else if fc4.Get() {
			println("pushing: fc4")
			kb.Down(keyboard.KeyLeft)
		} else {
			kb.Up(keyboard.KeyRight)
			kb.Up(keyboard.KeyUp)
			kb.Up(keyboard.KeyDown)
			kb.Up(keyboard.KeyLeft)
		}
		time.Sleep(32*time.Millisecond - wait*3)
	}
}
