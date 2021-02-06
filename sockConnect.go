package gossip

import (
	"fmt"
	"net"
)

func connectHost(h host) {

	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", h.Addr, h.Port))

		if err != nil {
			continue
		}

		defer conn.Close()

		go func() {
			for {
				data := <-h.Write
				conn.Write(data)
			}
		}()

		data := make([]byte, 1024)

		for {
			_, err := conn.Read(data)
			if err != nil {
				break
			}
			h.Read <- data
		}
	}
}
