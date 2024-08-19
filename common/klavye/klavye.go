
package klavye

import (
	hotkey "golang.design/x/hotkey"
)

var copyComb *hotkey.Hotkey = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyC)
var pasteComb *hotkey.Hotkey = hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)

func CopyPasteDinle() (copy chan int, paste chan int) {
	copy = make(chan int)
	paste = make(chan int)
	go copyPasteDinle(copy, paste)
	return
}




