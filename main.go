package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	var ports []string
	var host string

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file content:", err)
		return
	}
	err = json.Unmarshal(jsonBytes, &ports)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Please type the IP address you want to scan:")
	fmt.Scanln(&host)

	for _, port := range ports {
		_, err = connectPort(host, port)
		conn, _ := connectPort(host, port)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, port))
		}
	}

}

func connectPort(host string, port string) (net.Conn, error) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	return conn, err
}
