package gui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ShowSelect(elements []string, title, column string) string {
	cmd := exec.Command(
		"zenity",
		"--list",
		"--column="+column,
		"--title="+title,
	)

	for _, name := range elements {
		cmd.Args = append(cmd.Args, name)
	}

	midiOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	return string(strings.ReplaceAll(string(midiOutput), "\n", ""))
}
