package main

import(
    "fmt"
    "net"
    "bufio"
    "encoding/json"
    "os"
)

var hosts []string = []string{}
        
const (
    CONN_TYPE = "tcp"
)


func server() {

         ln, _ := net.Listen(CONN_TYPE, "localhost:8001")
        defer ln.Close()
        con, _ := ln.Accept()
        defer con.Close()
        r := bufio.NewReader(con)
        msg, _ := r.ReadString('\n')

        hosts  = append(hosts,msg)
        fmt.Print(hosts)

        dat, _ := json.MarshalIndent(hosts, "", " ")

        jsonFile, _:= os.Create("./data.json")
        jsonFile.Write(dat)

}

func main() {
    for{
        server()
    }
}
