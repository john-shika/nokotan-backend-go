package cores

import (
	"net/url"
	"strconv"
)

type NetworkImpl interface {
	GetScheme() string
	GetProtocol() string
	GetAddress() string
	GetPort() uint16
	GetHost() string
	GetURL() *url.URL
	String() string
}

type Network struct {
	Scheme   string `mapstructure:"scheme" json:"scheme" yaml:"scheme"`
	Protocol string `mapstructure:"protocol" json:"protocol" yaml:"protocol"`
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Port     uint16 `mapstructure:"port" json:"port" yaml:"port"`
}

func NewNetwork() NetworkImpl {
	return &Network{
		Scheme:   "http",
		Protocol: "tcp",
		Address:  "localhost",
		Port:     80,
	}
}

func (n *Network) GetNameType() string {
	return "network"
}

func (n *Network) GetScheme() string {
	return n.Scheme
}

func (n *Network) GetProtocol() string {
	return n.Protocol
}

func (n *Network) GetAddress() string {
	return n.Address
}

func (n *Network) GetPort() uint16 {
	return n.Port
}

func (n *Network) GetHost() string {
	return n.Address + ":" + strconv.Itoa(int(n.Port))
}

func (n *Network) GetURL() *url.URL {
	return &url.URL{
		Scheme: n.Scheme,
		Host:   n.GetHost(),
	}
}

func (n *Network) String() string {
	return n.GetURL().String()
}
