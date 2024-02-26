package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var IP string
	ports := []string{"21", "22", "25", "80", "443", "3306"}

	fmt.Println("Please type the IP address you want to scan:")
	fmt.Scanln(&IP)

	checkPort(IP, ports)
}

func checkPort(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, port))
		}
	}
}
