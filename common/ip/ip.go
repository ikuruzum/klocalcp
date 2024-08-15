package ip

import (
	"fmt"
	"net"
	"strings"
)
var PORT = 5001
// çoğunlukla 192.168.1
var routerIP string = ""

func init() {
	_, _ = MyLocalIP()
}

func RouterIP() string {
	return routerIP
}
func LocalIP(ld int) string {
	return routerIP + "." + fmt.Sprint(ld)
}

func MyLocalIP() (string, bool) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ip := fmt.Sprint(ipNet.IP)
			if strings.Contains(ip, "192.168.") {
				sp := strings.Split(ip, ".")
				routerIP = strings.Join(sp[0:3], ".")
				return ip, true
			}
		}
	}
	return "", false
}



//192.168.1.2

//127.0.0.1 //loopback