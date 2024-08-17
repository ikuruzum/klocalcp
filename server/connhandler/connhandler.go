package connhandler

import (
	"fmt"
	"net"
	"strings"
)

func Handle(conn net.Conn) {

	var endpoint, err = getEndpoint(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp string
	switch endpoint {
	case "/":
		resp = "HTTP/1.1 200 OK\r\nContent-Length: 7\r\n\r\n200 OK"
	case "/areyoualive":
		resp = "HTTP/1.1 200 OK\r\nContent-Length: 1\r\n\r\n1"
	default:
		resp = "HTTP/1.1 404 Not Found\r\nContent-Length: 9\r\n\r\n404 Not Found"
	}

	_, err = conn.Write([]byte(resp))
	conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getEndpoint(conn net.Conn) (string, error) {

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		return "", err
	}

	return strings.Split(string(buf), " ")[1], nil
}
