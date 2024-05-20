package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	// load crt file from disk
	crt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatalln(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{crt},
		InsecureSkipVerify: true,
	}
	l, err := ldap.DialURL("ldaps://localhost:6636", ldap.DialWithTLSConfig(tlsConfig))
	if err != nil {
		log.Fatalln(err)
	}

	err = l.ExternalBind()
	if err != nil {
		l.Close()
		log.Fatalln(err)
	}
	l.SetTimeout(10 * time.Second)

	whoami, err := l.WhoAmI(nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(whoami.AuthzID)
}
