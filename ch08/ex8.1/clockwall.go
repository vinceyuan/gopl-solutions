package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type host struct {
	name    string
	address string
	port    string
}

func main() {
	hosts := make([]host, 0)
	if len(os.Args) > 1 {
		for i, str := range os.Args {
			if i == 0 {
				continue
			}
			split1 := strings.Split(str, "=")
			name := split1[0]
			split2 := strings.Split(split1[1], ":")
			address := split2[0]
			port := split2[1]
			hosts = append(hosts, host{name, address, port})
		}
	} else {
		hosts = append(hosts, host{"local", "localhost", "8000"})
	}
	fmt.Println(hosts)

	conns := make([]net.Conn, len(hosts))
	for i, host := range hosts {
		hoststr := host.address + ":" + host.port
		conn, err := net.Dial("tcp", hoststr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		conns[i] = conn
		for j := 0; j < i; j++ {
			fmt.Print("\t\t\t")
		}
		fmt.Print(host.name)
	}
	fmt.Println()

	for {
		for i, conn := range conns {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err == io.EOF {
				os.Exit(0)
			}
			if err == nil && n > 0 {
				os.Stdout.WriteString("\r")
				for j := 0; j < i; j++ {
					os.Stdout.WriteString("\t\t\t")
				}
				os.Stdout.Write(b)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
