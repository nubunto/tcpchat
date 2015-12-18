package main

import (
	"flag"
	"github.com/nubunto/tcpchat/client"
	"log"
)

func main() {
	ip := flag.String("ip", "127.0.0.1:8080", "The server IP")
	name := flag.String("name", "", "Your name on the server")
	flag.Parse()
	if *name != "" {
		client.Connect(*ip, *name)
	} else {
		log.Fatalln("You must specify a name with -name.")
	}
}