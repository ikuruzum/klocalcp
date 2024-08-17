package server

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func Start() {
	listen, err := net.Listen("tcp", ":5001")
	lookForIps(false)
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		var endpoint, err = getEndpoint(conn)
		if err != nil {
			fmt.Println(err)
			return
		}

		var header = ""
		var body = ""

		switch endpoint {
		case "/":
			fmt.Print(endpoint, " HTTP/1.1 200 OK\r\n\r\n")
			body = "200 OK"
			header = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprint(len(body)) + "\r\n\r\n"

		case "/areyoualive":
			fmt.Print(endpoint, " HTTP/1.1 200 OK\r\n\r\n")
			body = "1"
			header = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprint(len(body)) + "\r\n\r\n"
		default:
			fmt.Print(endpoint, " HTTP/1.1 404 Not Found\r\n\r\n")
			body = "404 Not Found"
			header = "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprint(len(body)) + "\r\n\r\n"
		}

		_, err = conn.Write([]byte(header + body))
		if err != nil {
			fmt.Println(err)
			return
		}
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

var wg sync.WaitGroup
var ips []string

func lookForIps(widesearch bool) /*([]net.IP, error)*/ {
	var maxroute = 1
	if widesearch {
		maxroute = 255
	}
	for j := 1; j <= maxroute; j++ {
		for i := 1; i < 256; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				lookForIp(fmt.Sprintf("192.168.%d.%d:5001", j, i))
			}()
		}
		wg.Wait()
	}
	fmt.Println(ips)
}

func lookForIp(ip string) {
	d := net.Dialer{Timeout: 5 * time.Second}
	dial, err := d.Dial("tcp", ip)
	if err == nil {
		ips = append(ips, ip)

		dial.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
		var x []byte = make([]byte, 1024)
		dial.Read(x)
		fmt.Println(string(x))

		dial.Close()
		return
	}
}
