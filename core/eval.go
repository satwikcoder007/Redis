package core

import (
	"net"
	"errors"
)

func evalPing(conn net.Conn, args []string) error {
	var response []byte
	if len(args) > 1 {
		return errors.New("invalid number of arguments for PING")
	}
	if len(args) == 0 {
		response  = Encode("PONG",true)
	}else {
		response = Encode(args[0],false)
	}
	_,err := conn.Write(response)
	return err
}

func EvalAndRespond(conn net.Conn, cmd *RedisCmd) error {

	switch cmd.Cmd {
	case "PING":
		return evalPing(conn, cmd.Args)
	default:
		return errors.New("unknown command: " + cmd.Cmd)
	}
}