package controllers

import (
	"net"
	"os"
)

func get_my_ip() string {
	var my_ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				my_ip = ipnet.IP.String() + "\n"
			}
		}
	}
	return my_ip
}
