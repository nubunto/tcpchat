package main

import (
	"flag"
	"github.com/nubunto/tcpchat/server"
)

func main() {
	wait := make(chan struct{})
	ip := flag.String("ip", "127.0.0.1:8080", "The server IP")
	flag.Parse()
	go server.Start(*ip)
	<-wait
}
