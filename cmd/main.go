package main

import (
	"MidiReader/internal/gui"
	"MidiReader/internal/processMIDIMessage"
	"MidiReader/internal/utils/getFiles"
	"fmt"
	"os"
)

/*

START: выводим все миди и выбираем одну из них /dev/showmidi(num)

Если есть
1. загружаем конфиг / если их несколько - просим выбрать один из них
2. связываем все нажатия по конфигу
-------------------------------------------------------------------------------------------
Если конфига нет
1. просим нажимать кнопку на миди а потом на клавиатуре для связки
2. нажимаем enter для генерации имени файла и сохранения
3. запускаем скрипт сначала
*/

func main() {
	midiNames, err := getFiles.GetFiles("/dev", "midi")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	midi := gui.ShowSelect(midiNames, "MIDI devices", "Device")

	if midi == "" {
		fmt.Println("No MIDI device selected")
		os.Exit(1)
	}

	midiFile, err := os.OpenFile("/dev/"+midi, os.O_RDWR, 0666)
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
