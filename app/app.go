package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func CreateApp() *cli.App {
	fmt.Println("\nIPFAS - IP Finder as well as Server Names :)")
	app := cli.NewApp()

	app.Name = "IPFAS - IP Finder as Server Names"
	app.Usage = "./ipfas ip -host www.google.com"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find ips related to a domain",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "nameservers",
			Usage:  "Find nameservers related to a domain",
			Flags:  flags,
			Action: findNameServers,
		},
	}
	return app
}

func findIps(c *cli.Context) {
	hostname := c.String("host")
	fmt.Println("\nList of Ips for this Domain: ")
	fmt.Println("")

	ips, err := net.LookupIP(hostname)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	fmt.Println()
}

func findNameServers(c *cli.Context) {
	hostname := c.String("host")
	fmt.Println("\nList of NS for this Domain: ")
	fmt.Println("")

	nameservers, err := net.LookupNS(hostname)
	if err != nil {
		log.Fatal(err)
	}

	for _, nameserver := range nameservers {
		fmt.Println(nameserver.Host)
	}

	fmt.Println()
}
