package main

import (
	"context"
	"flag"
	"github.com/gokrazy/gokrazy"
	"log"
	"os/signal"
	"strings"
	"syscall"

	"github.com/brutella/dnssd"
	dnssdlog "github.com/brutella/dnssd/log"
)

func main() {
	dnssdlog.Debug.Enable()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	/* TODO: Add flags for other customizable fields
	dnssd.Service{
		Name:   "",
		Type:   "",
		Domain: "",
		Host:   "",
		Text:   nil,
		TTL:    0,
		Port:   0,
		IPs:    nil,
		Ifaces: nil,
	}
	*/

	var Host string
	var Domain string
	var Ifaces []string

	var IfacesStr string

	flag.StringVar(&Host, "host", "gokrazy", "mDNS host")
	flag.StringVar(&Domain, "domain", "local", "mDNS domain")
	flag.StringVar(&IfacesStr, "ifaces", "eth0,wlan0", "broadcast interfaces (comma-separated)")
	flag.Parse()
	
	Ifaces = strings.Split(IfacesStr, ",")

	// Wait until network interfaces have a chance to work.
	gokrazy.WaitForClock()

	resp, err := dnssd.NewResponder()
	if err != nil {
		log.Fatalf("new responder: %s", err)
	}
	// Just need the A record.
	_, err = resp.Add(dnssd.Service{
		Domain: Domain,
		Host:   Host,
		Ifaces: Ifaces,
	})
	if err != nil {
		log.Fatalf("add service: %s", err)
	}
	err = resp.Respond(ctx)
	if err != nil {
		log.Fatalf("respond: %s", err)
	}
}
