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
	var availableIps chan []string = make(chan []string)
	//var dials = make(chan net.Dialer)

	go ipChecker(availableIps)
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
	defer conn.Close()

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

var wg sync.WaitGroup

func ipChecker(availableIps chan []string) {
	for {
		availableIps <- lookForIps(false)

		time.Sleep(60 * time.Second)
	}

}

func lookForIps(widesearch bool) (availableIps []string) {
	var maxroute = 1
	if widesearch {
		maxroute = 255
	}
	//var ips = []string{}
	for j := 1; j <= maxroute; j++ {
		for i := 1; i < 256; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				found := lookForIp(fmt.Sprintf("192.168.%d.%d:5001", j, i))
				if found {
					fmt.Println("found ip 192.168." + fmt.Sprint(j) + "." + fmt.Sprint(i) + ":5001")
					availableIps = append(availableIps, fmt.Sprintf("192.168.%d.%d:5001", j, i))
				}
			}()
		}
		wg.Wait()
	}
	fmt.Println(availableIps)
	return availableIps
}

func lookForIp(ip string) (found bool) {
	d := net.Dialer{Timeout: 5 * time.Second}
	dial, err := d.Dial("tcp", ip)
	if err == nil {
		//ips = append(ips, ip)

		dial.Write([]byte("GET /areyoualive HTTP/1.1\r\n\r\n"))
		var x []byte = make([]byte, 1024)
		dial.Read(x)
		fmt.Println(string(x))

		dial.Close()
		return true
	}
	return false
}
