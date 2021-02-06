package gossip

import (
	"log"
	"net"
)

func whoAmI() (addresses []string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		// This will do for now;
		// Need better error handling
		log.Fatal(err)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			// This will do for now;
			// Need better error handling
			log.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			addresses = append(addresses, ip.String())
		}
	}
	return
}
