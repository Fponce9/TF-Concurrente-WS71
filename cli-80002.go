package main

import (
    "fmt"
    "net"

)

func main() {
    con, _ := net.Dial("tcp", "localhost:8001")
    defer con.Close()

    fmt.Fprintln(con, "46456!")
}
