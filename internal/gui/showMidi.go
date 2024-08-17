package gui

import (
	"MidiReader/internal/utils/getFiles"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ShowMidi() string {
	midiNames, err := getFiles.GetFiles("/dev", "midi")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	title := "MIDI devices"
	column := "Device"

	cmd := exec.Command(
		"zenity",
		"--list",
		"--column="+column,
		"--title="+title,
	)

	for _, name := range midiNames {
		cmd.Args = append(cmd.Args, name)
	}

	midiOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	return string(strings.ReplaceAll(string(midiOutput), "\n", ""))
}
