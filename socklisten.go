package gossip

import (
	"fmt"
	"log"
	"net"
)

//sockListen
func sockListen() {

	sockServ, e := net.Listen("tcp", fmt.Sprintf("%s:%d", "", 8086))

	if e != nil {
		log.Fatal("Error setting up a listening connection", e)
	}

	defer sockServ.Close()

	for {
		if client, e := sockServ.Accept(); e == nil {
			go sockClient(client)
		}
	}
}

// schedulerListen
func schedulerListen(host string, w chan string) {

	sockServ, e := net.Listen("tcp", fmt.Sprintf("%s:7998", ""))
	if e != nil {
		log.Fatal("Error setting up a listening connection", e)
	}

	defer sockServ.Close()

	for {
		if client, e := sockServ.Accept(); e == nil {
			_ = client
		}
	}
}
