package klavye

import (
	"log"

	"golang.design/x/hotkey"
)

var copyComb *hotkey.Hotkey = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyC)
var pasteComb *hotkey.Hotkey = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)

func CopyPasteDinle() (copy chan int, paste chan int) {
	copy = make(chan int)
	paste = make(chan int)
	go copyPasteDinle(copy, paste)
	return
}




func copyPasteDinle(copy chan int, paste chan int) {
	err := copyComb.Register()
	err2 := pasteComb.Register()
	if err != nil || err2 != nil {
		log.Fatal("HATA klavye dinlenemiyor. ")
	}
	for {
		select {
		case <-copyComb.Keydown():
			copy <- 1
		case <-copyComb.Keyup():
			copy <- 0
		case <-pasteComb.Keydown():
			paste <- 1
		case <-pasteComb.Keyup():
			paste <- 0
		}
	}
}

