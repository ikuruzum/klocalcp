package server

import (
	"fmt"
	"klocalcp/server/connhandler"
	"klocalcp/server/ipchecker"
	"net"
)

var PORT = ipchecker.PORT

func Start() {
	listen, err := net.Listen("tcp", ":"+fmt.Sprint(PORT))

	go ipchecker.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go connhandler.Handle(conn)
	}
}
