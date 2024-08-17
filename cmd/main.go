package main

import (
	"MidiReader/internal/gui"
	"MidiReader/internal/processMIDIMessage"
	"fmt"
	"os"
)

/*

TODO:
Если есть
1. загружаем конфиг / если их несколько - просим выбрать один из них
2. связываем все нажатия по конфигу
-------------------------------------------------------------------------------------------
Если конфига нет
1. находим миди клавиатуру /dev/showmidi(num)  (если их несколько то просим выбрать одну из них
2. просим нажимать кнопку на миди а потом на клавиатуре для связки
3. нажимаем enter для генерации имени файла и сохранения
4. запускаем скрипт сначала
*/

func main() {
	midi := gui.ShowMidi()

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
