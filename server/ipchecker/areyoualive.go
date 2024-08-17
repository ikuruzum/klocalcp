package ipchecker

import (
	"net"
	"strings"
)

func areyoualive(dial net.Conn) (alive bool) {
	dial.Write([]byte("GET /areyoualive HTTP/1.1\r\n\r\n"))
	var resp []byte = make([]byte, 1024)
	dial.Read(resp)
	dial.Close()
	return strings.Contains(string(resp), "\r\n\r\n1")
}
