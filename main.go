package main

import (
	"flag"

	"github.com/satwikcoder007/Redis/config"
	"github.com/satwikcoder007/Redis/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}
func main() {
	setupFlags()
	server.RunTcpServer()
}
