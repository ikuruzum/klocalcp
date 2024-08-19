//go:build linux
// +build linux

package platform

import (
	"klocalcp/common/klavye/hotkeys"
	"log"
)

func BindChans(copy chan int, paste chan int) {
// burada ise C ile wayland kontrolü ve wayland ise ona yönelik bir şeyler yapçez 
	err := hotkeys.Copy.Register()
	err2 :=  hotkeys.Paste.Register()
	if err != nil || err2 != nil {
		log.Fatal("HATA klavye dinlenemiyor. ")
	}
	for {
		select {
		case <-hotkeys.Copy.Keydown():
			copy <- 1
		case <-hotkeys.Copy.Keyup():
			copy <- 0
		case <-hotkeys.Paste.Keydown():
			paste <- 1
		case <-hotkeys.Paste.Keyup():
			paste <- 0
		}
	}
}
