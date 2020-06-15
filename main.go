package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/matbot/ftpGoServer/ftp"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", "public", "root directory")
	flag.Parse()
}

func main() {
	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConnection(c, absPath))
}
