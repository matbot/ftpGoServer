package ftp

type dataType int

const (
	// iota automates enum increment, starting at 0.
	ascii dataType = iota
	binary
)

func (c *Connection) setDataType(args []string) {
	if len(args) == 0 {
		c.respond(status501)
	}
	switch args[0] {
	case "A":
		c.dataType = ascii
	case "I":
		c.dataType = binary
	default:
		c.respond(status504)
		return
	}
	c.respond(status200)
}
