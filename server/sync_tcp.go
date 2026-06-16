package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/satwikcoder007/Redis/config"
)

func readCommand(conn net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := conn.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
func respond(conn net.Conn, cmd string) error {
		_,err := conn.Write([]byte(cmd))
		return err
}

func RunTcpServer() {
	//server loop

	log.Printf("Starting TCP server on %s:%d", config.Host, config.Port)

	var connection_client int  = 0

	listener, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	
	if err != nil{
		panic(err)
	}

	for{
		//handle incoming connections
		conn,err := listener.Accept()

		if err != nil{
			panic(err)
		}

		connection_client += 1
		log.Printf("New client connected with address: %d", connection_client)

		for {
			// handle client commands
			cmd,err := readCommand(conn)
			if err != nil {
				conn.Close()
				log.Printf("Client disconnected with address: %d", connection_client)
				connection_client -= 1
				//this handles normall client disconnection
				if err != io.EOF {
					log.Println("err",err)
				}
				break
			}
			log.Println("command",cmd)
			if err = respond(conn,cmd); err != nil {
				log.Println("err",err)
			}
		}

	}

}
