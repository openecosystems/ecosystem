package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/slackhq/nebula/config"
	"github.com/slackhq/nebula/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	configStr := `
tun:
  user: true
  disabled: true

static_host_map:
  '192.168.100.1': ['localhost:4242']

listen:
  host: 0.0.0.0
  port: 4242

lighthouse:
  am_lighthouse: true
  interval: 60
  hosts:
    - '192.168.100.1'

punchy:
  punch: true

firewall:
  outbound:
    # Allow all outbound traffic from this node
    - port: any
      proto: any
      host: any

  inbound:
    # Allow icmp between any nebula hosts
    - port: any
      proto: icmp
      host: any
    - port: any
      proto: any
      host: any

pki:
  ca: /etc/nebula/ca.crt
  cert: /etc/nebula/host.crt
  key: /etc/nebula/host.key
`
	var c config.C
	if err := c.LoadString(configStr); err != nil {
		return err
	}
	s, err := service.New(&c)
	if err != nil {
		return err
	}

	ln, err := s.Listen("tcp", ":1234")
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %s", err)
			break
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
			}
		}(conn)

		log.Printf("got connection")

		_, err = conn.Write([]byte("hello world\n"))
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			_, err = fmt.Fprintf(conn, "echo: %q\n", message)
			if err != nil {
				return err
			}
			log.Printf("got message %q", message)
		}

		if err := scanner.Err(); err != nil {
			log.Printf("scanner error: %s", err)
			break
		}
	}

	err = s.Close()
	if err != nil {
		return err
	}
	if err := s.Wait(); err != nil {
		return err
	}
	return nil
}
