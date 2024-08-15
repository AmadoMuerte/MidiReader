package processMIDIMessage

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

const (
	midiNoteOn  = 0x90 // Note On
	midiNoteOff = 0x80 // Note Off
)

var midiToKeyMap = map[byte]string{
	48: "y",
	50: "u",
	52: "i",
	53: "o",
	55: "p",
	57: "h",
	59: "j",
	60: "k",
	62: "l",
	64: ";",
	65: "n",
	67: "m",
	69: ",",
	71: ".",
	72: "/",
}

func New(msg [3]byte) {
	status := msg[0]
	note := msg[1]
	velocity := msg[2]

	switch status {
	case midiNoteOn:
		if velocity > 0 {
			if key, ok := midiToKeyMap[note]; ok {
				err := robotgo.KeyDown(key)
				if err != nil {
					_ = fmt.Errorf("error releasing key: %v", err)
				}
			}
		} else {
			if key, ok := midiToKeyMap[note]; ok {
				err := robotgo.KeyUp(key)
				if err != nil {
					return
				}
			}
		}
	case midiNoteOff:
		if key, ok := midiToKeyMap[note]; ok {
			err := robotgo.KeyUp(key)
			if err != nil {
				_ = fmt.Errorf("error releasing key: %v", err)
			}
		}
	}
}
