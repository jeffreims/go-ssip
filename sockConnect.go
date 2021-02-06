package gossip

import (
	"fmt"
	"net"
	"time"
)

func connectHost(h host) {

	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", h.Addr, h.Port))

		if err != nil {
			continue
		}

		defer conn.Close()

		data := make([]byte, 1024)

		for {
			_, err := conn.Read(data)
			if err != nil {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}
}
