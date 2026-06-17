package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"github.com/satwikcoder007/Redis/config"
	"github.com/satwikcoder007/Redis/core"
)

func readCommand(conn net.Conn) (*core.RedisCmd, error) {
	buf := make([]byte, 512)

	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	tokens, err := core.Decode(buf[:n])
	if err != nil {
		return nil, err
	}

	arr := tokens.([]interface{})

	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = v.(string)
	}

	return &core.RedisCmd{
		Cmd:  strings.ToUpper(strs[0]),
		Args: strs[1:],
	}, nil
}
func respondError (conn net.Conn, err error) {
	conn.Write([]byte(fmt.Sprintf("-%s\r\n", err.Error())))
}
func respond(conn net.Conn, cmd *core.RedisCmd) {
	//log.Printf("Received command: %s with args: %v", cmd.Cmd, cmd.Args)
	if err := core.EvalAndRespond(conn, cmd); err!= nil {
		respondError(conn, err)
	}
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
				respondError(conn, err)
				conn.Close()
				log.Printf("Client disconnected with address: %d", connection_client)
				connection_client -= 1
				//this handles normall client disconnection
				if err != io.EOF {
					log.Println("err",err)
				}
				break
			}
			respond(conn, cmd)
		}

	}

}
