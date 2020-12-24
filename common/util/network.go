package util

import (
	"errors"
	"net"
)

var ErrGetLocalIP = errors.New("Fail to get Local IP")

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", ErrGetLocalIP
}

func CheckPortRange(port int) bool {
	return port >= 1 && port <= 65535
}
