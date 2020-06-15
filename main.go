package ftpGoServer

import (
	"flag"
	"fmt"
	"log"
	"net"
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
