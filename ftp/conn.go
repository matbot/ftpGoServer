package ftp

import "net"

type Connection struct {
	connection net.Conn
	dataType   dataType
	dataPort   *dataPort
	rootDir    string
	workingDir string
}

func NewConnection(c net.Conn, rootDir string) *Connection {
	return &Connection{
		connection: c,
		dataType:   nil,
		dataPort:   nil,
		rootDir:    "/",
		workingDir: "",
	}
}
