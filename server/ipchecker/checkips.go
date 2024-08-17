package ipchecker

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var PORT = 5001
var MyIP = ""

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
				//defer wg.Done()
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
	fmt.Println(availableIps)
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
