package main

import (
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
	"strings"
	"time"
)

func scan(socket string) bool {
	portscan, err := net.DialTimeout("tcp", socket, 2*time.Second)
	if err != nil {
		return false
	}
	defer portscan.Close()
	return true
}

func main() {
	var ports = [...]int{22, 80, 8080, 443, 8443}

	var ip netip.Addr = netip.Addr{}
	var prefix netip.Prefix = netip.Prefix{}
	var scanips []netip.Addr
	var filename string
	var err error

	cmdArg := ""
	if len(os.Args) > 1 {
		cmdArg = os.Args[1]
	} else {
		log.Fatal("Provide at least one command argument")
	}

	prefix, err = netip.ParsePrefix(cmdArg)
	if err != nil {
		ip, err = netip.ParseAddr(cmdArg)
		if err != nil {
			log.Fatal(err)
		}
		scanips = append(scanips, ip)
		filename = ip.String()
	} else {
		for ip = prefix.Addr(); prefix.Contains(ip); ip = ip.Next() {
			scanips = append(scanips, ip)
		}
		// scanips = scanips[2 : len(scanips)-1] // exclude network, gateway and broadcast
		filename = strings.Replace(prefix.String(), "/", "_", -1)
	}

	saveFile, err := os.Open(filename)
	if err != nil {
		saveFile, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer saveFile.Close()

	for _, ip = range scanips {
		_, err := fmt.Fprintln(saveFile, ip)
		if err != nil {
			log.Fatal(err)
		}

		for _, port := range ports {
			isOpen := scan(fmt.Sprintf("%s:%d", ip, port))
			if isOpen {
				fmt.Fprintf(saveFile, "* %d/tcp open\n", port)
			} else {
				fmt.Fprintf(saveFile, "* %d/tcp close\n", port)
			}
		}
	}
}
