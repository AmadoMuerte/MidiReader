package main

import (
	"MidiReader/internal/processMIDIMessage"
	"fmt"
	"os"
)

func main() {
	midiFile, err := os.OpenFile("/dev/midi3", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error opening MIDI device:", err)
		return
	}
	defer func(midiFile *os.File) {
		err := midiFile.Close()
		if err != nil {
			fmt.Println("Error closing MIDI device:", err)
		}
	}(midiFile)

	fmt.Println("Listening for MIDI messages...")

	for {
		var msg [3]byte
		_, err := midiFile.Read(msg[:])
		if err != nil {
			fmt.Println("Error reading MIDI message:", err)
			continue
		}
		processMIDIMessage.New(msg)
	}
}
