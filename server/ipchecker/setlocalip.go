package ipchecker

import (
	"fmt"
	"net"
)

func setLocalAddr() {
	conn, err := net.Dial("udp", "192.168.1.1:1")
	if err != nil {
		fmt.Println("Local Ip Couldn't set")
		return
	}
	MyIP = conn.LocalAddr().(*net.UDPAddr).IP.String() + fmt.Sprintf(":%d", PORT)
	conn.Close()
}
