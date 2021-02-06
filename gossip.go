package gossip

type host struct {
	Addr        string
	Port        int
	Write, Read chan []byte
}

type Gossip struct {
	// Hosts to connect with
	Hosts []host
}

func (g *Gossip) AddHost(addr string, port int) {
	var w = make(chan []byte)
	var r = make(chan []byte)
	g.Hosts = append(g.Hosts, host{Addr: addr, Port: port, Write: w, Read: r})
}

func (g *Gossip) Connect() {
	myAddrs := whoAmI()
	_ = myAddrs
	for _, host := range g.Hosts {
		connectHost(host)
	}
}
