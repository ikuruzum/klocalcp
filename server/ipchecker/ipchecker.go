package ipchecker

import (
	"time"
)

var AvailableIps = make(chan []string)

func Run() {
	for {
		AvailableIps <- checkIps(false)

		time.Sleep(60 * time.Second)
	}

}
