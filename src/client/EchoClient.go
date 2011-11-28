package main

import (
    "flag"
    "net"
)

var hostName *string = flag.String("h", "127.0.0.1", "address to listen on")
var portNumber *string = flag.String("p", "4242", "port number to listen to")

func main() {
    flag.Parse()

    conn, _ := net.Dial("tcp", *hostName + ":" + *portNumber)
    conn.Write([]byte("test"))
    conn.Close()

}

