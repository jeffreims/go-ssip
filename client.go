package gossip

import (
	"net"
	"strings"
)

func sockClient(client net.Conn) {

	defer client.Close()

	var cmdMsg []string
	var message string

	sockreader := make([]byte, 1024)
	for {
		read, e := client.Read(sockreader)
		if e != nil {
			break
		}

		message = strings.TrimSpace(string(sockreader[:read-1]))

		cmdMsg = strings.Split(message, " ")

		// Clear the vars
		cmdMsg = cmdMsg[:0]
		message = ""
	}

	//clnt.Write([]byte(fmt.Sprintln(out.rawKey)))
}
