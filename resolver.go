package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func exitHelp() {
	panic("./resolver <DNS server> <host name> (only SRV, A and AAAA are implemented)")
}

func exitErr(err error) {
	panic(err.Error())
}

func printHostResults(addrs []string) {
	fmt.Println("Result of HOST lookup is:")
	for _, v := range addrs {
		fmt.Printf("\t%s\n", v)
	}
}

func printSrvResults(cname string, addrs []*net.SRV) {
	fmt.Println("Result of SRV lookup is:")
	for _, v := range addrs {
		fmt.Printf("\t%s: %#v\n", cname, v)
	}
}

func getResolver(dnsServer string) *net.Resolver {
	splitted := strings.Split(dnsServer, ":")
	if len(splitted) == 1 {
		dnsServer += ":53"
	} else if len(splitted) > 2 {
		panic("Invalid dnsServer")
	}
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, dnsServer)
		},
	}
}

func main() {
	if len(os.Args) < 4 {
		exitHelp()
	}
	lookupServer := os.Args[1]
	lookupType := os.Args[2]
	lookupHost := os.Args[3]

	r := getResolver(lookupServer)

	switch lookupType {
	case "SRV":
		cname, addrs, err := r.LookupSRV(context.Background(), "", "", lookupHost)
		if err != nil {
			exitErr(err)
		}
		printSrvResults(cname, addrs)
	case "A", "AAAA", "HOST":
		addrs, err := r.LookupHost(context.Background(), lookupHost)
		if err != nil {
			exitErr(err)
		}
		printHostResults(addrs)
	default:
		exitHelp()
	}

}
