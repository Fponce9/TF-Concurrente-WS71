package main

import (
	"fmt"
	"net"
)

func main() {

	newIP := "192.168.1.50"

	ips := []string{"192.168.1.10", "192.168.1.20", "192.168.1.30"}

	for i := 0; i < len(ips); i++ {
		con, _ := net.Dial("tcp", ips[i])
		defer con.Close()
		fmt.Fprintln(con, newIP)
	}

}
