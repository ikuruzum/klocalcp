package klavye

import (
	ps 	"klocalcp/common/klavye/platform"
)

func CopyPasteDinle() (copy chan int, paste chan int) {
	copy = make(chan int)
	paste = make(chan int)
	go ps.BindChans(copy, paste)
	return
}
