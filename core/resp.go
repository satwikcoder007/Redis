package core

import (
	"errors"
)

func readSimpleString(data []byte) (string,int, error) {
	pos := 0
	for ; data[pos] != '\r'; pos++ {}
	return string(data[:pos]), pos + 2, nil
}
func readInteger(data []byte) (int, int, error) {
	pos := 0
	val := 0
	if data[0] == '-' {
		pos++
	}
	for ; data[pos] != '\r'; pos++ {
		val = val*10 + int(data[pos]-'0')
	}
	if data[0] == '-' {
		val = -val
	}
	return val, pos + 2, nil
}
func readBulkString(data []byte) (string, int, error) {
	pos := 0
	length := 0
	for ; data[pos] != '\r'; pos++ {
		length = length*10 + int(data[pos]-'0')
	}
	pos += 2 
	return string(data[pos:(pos+length)]), pos + length + 2, nil
}
func readArray(data []byte) ([]interface{}, int, error) {
	
	pos := 0
	length := 0
	var delta int 
	for ; data[pos] != '\r'; pos++ {
		length = length*10 + int(data[pos]-'0')
	}
	pos += 2
	array := make([]interface{}, length)

	for i:=0; i<length;i++ {
		array[i],delta,_  = decodeOne(data[pos:])
		pos += delta+1
	}

	return array, pos, nil
}
func readError(data []byte) (string,int, error) {
	return readSimpleString(data)
}
	
func decodeOne(data []byte) (interface{},int,error){
	switch data[0] {
	case '+':
		return readSimpleString(data[1:])
	
	case ':':
		return readInteger(data[1:])
	
	case '$':
		return readBulkString(data[1:])
		
	case '*':
		return readArray(data[1:])
	
	case '-':
		return readError(data[1:])
	
	default:
		return nil,0,errors.New("invalid RESP type")
	}
}
func Decode(data []byte) (interface{}, error) {
	if len(data) == 0{
		return nil, errors.New("no data")
	}
	value,_,err := decodeOne(data)
	return value, err
}