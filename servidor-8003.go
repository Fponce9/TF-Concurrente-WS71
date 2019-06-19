package main

import (
    "fmt"
    "net"
	"encoding/json"
)

var hosts []string = []string{"transacción 1",
							"transacción 2",
							"transacción 3",
							"transacción 4",
							"transacción 5"}

							

func handle(con net.Conn, id int) {
    defer con.Close()
    for {

			poio,_:=json.Marshal(hosts)
			poio2:=string(poio)
			fmt.Fprintln(con, poio2)
		
    }
    fmt.Printf("Con%d cerrada!\n", id)
}

func main() {
    ln, _ := net.Listen("tcp", "10.142.112.70:8000")
    defer ln.Close()
    cont := 0
    for {
        con, _ := ln.Accept()
        go handle(con, cont)
        cont++
    }
}
