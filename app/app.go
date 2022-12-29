package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the CLI application ready to be executed
func Generate() *cli.App {
	// Creating new application
	app := cli.NewApp()
	// Defining app settings
	app.Name = "CLI Application"
	app.Usage = "Search IPs and server names on internet"

	// Setting the flags which can be used on the commands
	flags := []cli.Flag{
		// Here, we define both parameter name and default value
		cli.StringFlag{
			Name:  "host",
			Value: "mhsw.com.br",
		},
	}

	// Defining the application commands
	app.Commands = []cli.Command{
		// IP search command
		{
			// Command name
			Name: "ip",
			// Command usage description
			Usage: "Search IPs addresses for hosts on the internet",
			// Flags which can be provided for the command
			Flags: flags,
			// Action to be taken when command is executed
			Action: searchIps,
		},
		// Server name command
		{
			// Command name
			Name: "servers",
			// Command usage description
			Usage: "Search the server names for hosts on the internet",
			// Flags which can be provided for the command
			Flags: flags,
			// Action to be taken when command is executed
			Action: searchServers,
		},
	}

	return app
}

// Defining search IPs function
func searchIps(c *cli.Context) {
	// Getting the 'host' parameter value
	host := c.String("host")

	// Calling the function to get the IP from a host
	ips, error := net.LookupIP(host)
	// If an error occurs
	if error != nil {
		log.Fatal(error)
	}

	// Now, for each IP, we'll show it to the user
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Defining search server names function
func searchServers(c *cli.Context) {
	// Getting the 'host' parameter value
	host := c.String("host")

	// Calling the function to get the server name from a host
	servers, error := net.LookupNS(host)
	// If an error occurs
	if error != nil {
		log.Fatal(error)
	}

	// Now, for each server name, we'll show it to the user
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
