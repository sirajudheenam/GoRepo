package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

func main() {
    config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
    c := new(dns.Client)

    m := new(dns.Msg)
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname of MX record>\n", os.Args[0])
		os.Exit(1)
	}
  
    m.SetQuestion(dns.Fqdn(os.Args[1]), dns.TypeMX)
    m.RecursionDesired = true

    r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
            log.Fatalf(" *** invalid answer name %s after MX query for %s\n", os.Args[1], os.Args[1])
    }
    // Stuff must be in the answer section
    for _, a := range r.Answer {
            fmt.Printf("%v\n", a)
    }
}