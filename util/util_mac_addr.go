package util

import (
	"net"
)

// GetMACAddr 获取MAC地址
func GetMACAddr() string {
	var addr string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				if ok && ipnet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					addr = netInterfaces[i].HardwareAddr.String()
				}
			}
		}
	}
	return addr
}
