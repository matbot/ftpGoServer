package ftp

import (
	"fmt"
	"strings"
)

// TODO: enhance user functionality with authentication and permissions.
func (c *Connection) user(args []string) {
	c.respond(fmt.Sprintf(status230, strings.Join(args, " ")))
}
