package main

import (
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
)

const (
	midiNoteOn  = 0x90 // Note On message
	midiNoteOff = 0x80 // Note Off message
)

// Параметры для маппинга MIDI нот на клавиши клавиатуры
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

func main() {
	// Открываем MIDI устройство
	midiFile, err := os.OpenFile("/dev/midi3", os.O_RDWR, 0666) // или "/dev/snd/midiC0D0"
	if err != nil {
		fmt.Println("Error opening MIDI device:", err)
		return
	}
	defer midiFile.Close()

	fmt.Println("Listening for MIDI messages...")

	for {
		var msg [3]byte
		_, err := midiFile.Read(msg[:])
		if err != nil {
			fmt.Println("Error reading MIDI message:", err)
			continue
		}

		processMIDIMessage(msg)
	}
}

func processMIDIMessage(msg [3]byte) {
	status := msg[0]
	note := msg[1]
	velocity := msg[2]

	switch status {
	case midiNoteOn:
		if velocity > 0 {
			fmt.Printf("Note On: Note: %d, Velocity: %d\n", note, velocity)
			if key, ok := midiToKeyMap[note]; ok {
				robotgo.KeyDown(key) // Нажимаем клавишу
			}
		} else {
			// Если velocity = 0 при Note On, это считается как Note Off
			fmt.Printf("Note Off: Note: %d\n", note)
			if key, ok := midiToKeyMap[note]; ok {
				robotgo.KeyUp(key) // Отпускаем клавишу
			}
		}
	case midiNoteOff:
		fmt.Printf("Note Off: Note: %d\n", note)
		if key, ok := midiToKeyMap[note]; ok {
			robotgo.KeyUp(key) // Отпускаем клавишу
		}
	default:
		// Игнорируем другие MIDI сообщения
	}
}
