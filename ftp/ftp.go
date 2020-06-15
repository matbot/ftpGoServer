package ftp

import (
	"bufio"
	"log"
	"strings"
)

func Serve(c *Connection) {
	// status 220 to client for successful connection.
	c.respond(status220)

	s := bufio.NewScanner(c.connection)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD":
			c.cwd(args)
		case "LIST":
			c.list(args)
		case "PORT":
			c.port(args)
		case "USER":
			c.user(args)
		case "QUIT":
			// status 221 acknowledges disconnect from client.
			c.respond(status221)
			return
		case "RETR":
			c.retr(args)
		case "TYPE":
			c.setDataType(args)
		default:
			c.respond(status502)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
