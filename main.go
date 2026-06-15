package main 

import (
	"flag"
	"log"

	"my-redis/config"

)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}
func main() {
	setupFlags()
	log.Printf("Starting server on %s:%d", config.Host, config.Port)
}