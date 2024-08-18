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
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	go ipchecker.Run()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go connhandler.Handle(conn)
	}
}
