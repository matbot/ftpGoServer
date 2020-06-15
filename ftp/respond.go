package ftp

import (
	"fmt"
	"log"
)

// collection of standard ftp response statuses, with messages.
const (
	status150 = "150 File status okay; opening data connection."
	status200 = "200 Command okay."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status230 = "230 User %s logged in."
	status425 = "425 Error opening data connection."
	status426 = "426 Connection closed; file transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Command-parameter mismatch/not supported."
	status550 = "550 File unavailable."
)

func (c *Connection) respond(s string) {
	log.Print(">> ", s)
	// responses are terminated by custom EOL() sequence determined by data type.
	_, err := fmt.Fprint(c.connection, s, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

func (c *Connection) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}
