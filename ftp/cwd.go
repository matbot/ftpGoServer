package ftp

import (
	"log"
	"os"
	"path/filepath"
)

// TODO: Improve implementation of dir handling and implement path permissions.
func (c *Connection) cwd(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}
	workingDir := filepath.Join(c.workingDir, args[0])
	absPath := filepath.Join(c.rootDir, workingDir)
	_, err := os.Stat(absPath)
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	c.workingDir = workingDir
	c.respond(status220)
}
