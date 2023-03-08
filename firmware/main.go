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

type keycodes struct {
	key1 keyboard.Keycode
	key2 keyboard.Keycode
	key3 keyboard.Keycode
	key4 keyboard.Keycode
}

func New(key1, key2, key3, key4 keyboard.Keycode) *keycodes {
	return &keycodes{key1, key2, key3, key4}
}

var keys *keycodes

func init() {
	keys = New(
		keyboard.KeyRight,
		keyboard.KeyUp,
		keyboard.KeyDown,
		keyboard.KeyLeft,
	)
}

func main() {
	for _, fc := range []machine.Pin{fc1, fc2, fc3, fc4} {
		fc.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
	kb := keyboard.New()
	wait := 1 * time.Millisecond
	for {
		if fc1.Get() {
			println("pushing: fc1")
			kb.Down(keys.key1)
		} else if fc2.Get() {
			println("pushing: fc2")
			kb.Down(keys.key2)
		} else if fc3.Get() {
			println("pushing: fc3")
			kb.Down(keys.key3)
		} else if fc4.Get() {
			println("pushing: fc4")
			kb.Down(keys.key4)
		} else {
			kb.Up(keys.key1)
			kb.Up(keys.key2)
			kb.Up(keys.key3)
			kb.Up(keys.key4)
		}
		time.Sleep(32*time.Millisecond - wait*3)
	}
}
