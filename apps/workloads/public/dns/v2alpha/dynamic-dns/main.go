package main

import (
	"context"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

func main() {
	//bounds := []sdkv2betalib.Binding{
	//	&nebulav1.Binding{},
	//}
	//
	//_, err := sdkv2betalib.NewSpecYamlSettingsProvider()
	//if err != nil {
	//	fmt.Println("SpecError:", err)
	//	return
	//}
	//
	//path, handler := cryptographyv2alphapbconnect.NewEncryptionServiceHandler(&cryptographyv2alphasrv.EncryptionServiceHandler{})
	//server := sdkv2betalib.NewRawServer(context.Background(), bounds, path, &handler)
	//
	//server.ListenAndServe()

	server := sdkv2betalib.NewServer(context.Background(), nil)

	server.ListenAndServe()
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"net"
//
//	"github.com/slackhq/nebula/config"
//	"github.com/slackhq/nebula/service"
//)
//
//func main() {
//	if err := run(); err != nil {
//		log.Fatalf("%+v", err)
//	}
//}
//
//func run() error {
//	configStr := `
//tun:
//  disabled: false
//  dev: nebula1
//  drop_local_broadcast: false
//  drop_multicast: false
//  tx_queue: 500
//  mtu: 1300
//  routes:
//
//listen:
//  host: 0.0.0.0
//  port: 4242
//
//lighthouse:
//  am_lighthouse: true
//
//punchy:
//  punch: true
//  respond: true
//
//firewall:
//  outbound:
//    # Allow all outbound traffic from this node
//    - port: any
//      proto: any
//      host: any
//
//  inbound:
//    # Allow icmp between any nebula hosts
//    - port: any
//      proto: icmp
//      host: any
//    - port: any
//      proto: any
//      host: any
//
//pki:
//  ca: /etc/nebula/ca.crt
//  cert: /etc/nebula/host.crt
//  key: /etc/nebula/host.key
//`
//	var c config.C
//	if err := c.LoadString(configStr); err != nil {
//		return err
//	}
//	s, err := service.New(&c)
//	if err != nil {
//		return err
//	}
//
//	ln, err := s.Listen("tcp", ":1234")
//	if err != nil {
//		return err
//	}
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			log.Printf("accept error: %s", err)
//			break
//		}
//		defer func(conn net.Conn) {
//			_ = conn.Close()
//		}(conn)
//
//		log.Printf("got connection")
//
//		_, err = conn.Write([]byte("hello world\n"))
//		if err != nil {
//			return err
//		}
//
//		scanner := bufio.NewScanner(conn)
//		for scanner.Scan() {
//			message := scanner.Text()
//			_, err = fmt.Fprintf(conn, "echo: %q\n", message)
//			if err != nil {
//				return err
//			}
//			log.Printf("got message %q", message)
//		}
//
//		if err := scanner.Err(); err != nil {
//			log.Printf("scanner error: %s", err)
//			break
//		}
//	}
//
//	err = s.Close()
//	if err != nil {
//		return err
//	}
//	if err := s.Wait(); err != nil {
//		return err
//	}
//	return nil
//}
