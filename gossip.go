package gossip

type host struct {
	Addr string
	Port int
}

type Gossip struct {
	// Hosts to connect with
	Hosts []host
}

func (g *Gossip) AddHost(addr string, port int) {
	g.Hosts = append(g.Hosts, host{Addr: addr, Port: port})
}

func (g *Gossip) Connect() {
	myAddrs := whoAmI()
	_ = myAddrs
	for _, host := range g.Hosts {
		connectHost(host)
	}
}
