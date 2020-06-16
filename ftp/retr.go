package ftp

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func (c *Connection) retr(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}
	path := filepath.Join(c.rootDir, c.workingDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		c.respond(status550)
	}
	c.respond(status150)

	dataConnection, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(status425)
	}
	defer dataConnection.Close()

	// TODO: read/send file by line to reduce message load.
	_, err = io.Copy(dataConnection, file)
	if err != nil {
		log.Print(err)
		c.respond(status426)
		return
	}
	io.WriteString(dataConnection, c.EOL())
	c.respond(status226)
}
