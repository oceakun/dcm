package network

import (
	"net"
)

type NetworkInterface struct {
	Name          string
	MTU           int
	HardwareAddr  string
	Flags         net.Flags
	Addresses     []string
}

func GetNetworkInterfaces() []NetworkInterface {
	var interfaces []NetworkInterface

	ifs, err := net.Interfaces()
	if err != nil {
		return interfaces
	}

	for _, iface := range ifs {
		var addrs []string
		addressList, _ := iface.Addrs()
		for _, addr := range addressList {
			addrs = append(addrs, addr.String())
		}

		interfaces = append(interfaces, NetworkInterface{
			Name:          iface.Name,
			MTU:           iface.MTU,
			HardwareAddr:  iface.HardwareAddr.String(),
			Flags:         iface.Flags,
			Addresses:     addrs,
		})
	}

	return interfaces
}
