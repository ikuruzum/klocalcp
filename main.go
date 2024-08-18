package main

import (
	"klocalcp/common/klavye"
	"klocalcp/server"
)

func main() {
	go server.Start()
	go klavye.CopyPasteDinle()
}
