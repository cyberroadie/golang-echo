/*
 * User: cyberroadie
 * Date: 12/11/2011
 */
package main

import (
	"fmt"
	"flag"
	"net"
	"os"
	"syslog"
	"io"
)

type Settings struct {
	string PortNumber 
	string Hostname
	string Message
	bool Verbose
	bool Sctp
}

var settings = Settings{
	PortNumber: flag.String("p", "4242", "port number to listen to")
	Hostname: flag.String("h", "127.0.0.1", "address to listen on") 
}

func test(err os.Error, mesg string) {
    if err!=nil {
        syslog, _ := syslog.New(syslog.LOG_ERR, "echo server")
        _, err = io.WriteString(syslog, mesg)
        os.Exit(-1);
    }
}

func main() {
	flag.Parse()
	fmt.Printf("Listening on: %v:%v\n", *settings.Hostname, *settings.PortNumber)

    netlisten, err := net.Listen("tcp", *settings.Hostname, *settings.PortNumber)
    test(err, "Listen() error")
    defer netlisten.Close()

    conn, err := netlisten.Accept();
    test(err, "Accept() error")

    buf := make([]byte, 1024)
    conn.Read(buf)
    fmt.Println(string(buf))

}
