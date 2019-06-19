package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8001")
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	ips := []string{"192.168.1.10", "192.168.1.20", "192.168.1.30"}
	r := bufio.NewReader(con)
	IP, _ := r.ReadString('\n')
	ips = append(ips, IP)
	fmt.Println(IP)
}
