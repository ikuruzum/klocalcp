package klavye

import (
	"log"

	"golang.design/x/hotkey"
)

var CopyComb *hotkey.Hotkey = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyC)
var PasteComb *hotkey.Hotkey = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)

func CopyPasteDinle() (copy chan int, paste chan int) {
	copy = make(chan int)
	paste = make(chan int)
	go copyPasteDinle(copy, paste)
	return
}

func copyPasteDinle(copy chan int, paste chan int) {
	err := CopyComb.Register()
	err2 := PasteComb.Register()
	if err != nil || err2 != nil {
		log.Fatal("HATA klavye dinlenemiyor. ")
	}
	for {
		select {
		case <-CopyComb.Keydown():
			copy <- 1
		case <-CopyComb.Keyup():
			copy <- 0
		case <-PasteComb.Keydown():
			paste <- 1
		case <-PasteComb.Keyup():
			paste <- 0
		}
	}
}
