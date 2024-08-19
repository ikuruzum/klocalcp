//go:build !linux

package klavye

import "log"



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

