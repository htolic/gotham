package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
	var noChange bool = true
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

	saveFile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer saveFile.Close()

	reader := bufio.NewReader(saveFile)
	var bs []byte
	buf := bytes.NewBuffer(bs)

	for _, ip = range scanips {
		_, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		buf.WriteString(fmt.Sprintf("%s\n", ip.String()))
		fmt.Println(ip.String())
		noChange = true

		for _, port := range ports {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			result := ""
			isOpen := scan(fmt.Sprintf("%s:%d", ip, port))
			if isOpen {
				result = fmt.Sprintf("* %d/tcp open\n", port)
			} else {
				result = fmt.Sprintf("* %d/tcp close\n", port)
			}
			buf.WriteString(result)

			if line != result {
				fmt.Print(result)
				noChange = false
			}
		}

		if noChange {
			fmt.Println("No change")
		}
	}

	saveFile.Truncate(0)
	saveFile.Seek(0, 0)
	saveFile.WriteString(buf.String())
}
