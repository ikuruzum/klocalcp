package ipchecker

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var PORT = 5001
var MyIP = ""

// var AvailableIps = make(chan []string)
// var AvailableIps []string
type IpChecker struct {
	sync.Mutex
	Addrs []string
}

var AvailableIps = IpChecker{}

func Run() {
	for {
		AvailableIps.Lock()
		AvailableIps.Addrs = checkIps(false)
		AvailableIps.Unlock()
		fmt.Println(AvailableIps.Addrs)

		time.Sleep(60 * time.Second)
	}
}

func checkIps(widesearch bool) (availableIps []string) {
	setLocalAddr()
	var maxroute = 1
	if widesearch {
		maxroute = 255
	}
	for j := 1; j <= maxroute; j++ {
		wg := sync.WaitGroup{}
		for i := 1; i <= 255; i++ {
			wg.Add(1)
			go func() {
				ip := fmt.Sprintf("192.168.%d.%d:%d", j, i, PORT)
				if MyIP == ip {
					wg.Done()
					return
				}
				found := checkIp(ip)
				if found {
					availableIps = append(availableIps, ip)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
	//fmt.Println(availableIps)
	return availableIps
}

func checkIp(ip string) (found bool) {
	d := net.Dialer{Timeout: 2 * time.Second}
	dial, err := d.Dial("tcp", ip)
	if err == nil {
		return areyoualive(dial) //returns if available
	}
	return false //adress is empty
}

func setLocalAddr() {
	conn, err := net.Dial("udp", "192.168.1.1:1")
	if err != nil {
		fmt.Println("Local Ip Couldn't set")
		return
	}
	MyIP = conn.LocalAddr().(*net.UDPAddr).IP.String() + fmt.Sprintf(":%d", PORT)
	conn.Close()
}
