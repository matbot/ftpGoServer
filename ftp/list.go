package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func (c *Connection) list(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workingDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workingDir)
	}
	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	c.respond(status150)

	dataConnection, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(status425)
		return
	}
	for _, file := range files {
		_, err := fmt.Fprint(dataConnection, file.Name(), c.EOL())
		if err != nil {
			log.Print(err)
			c.respond(status426)
		}
	}
	_, err = fmt.Fprintf(dataConnection, c.EOL())
	if err != nil {
		log.Print(err)
		c.respond(status426)
	}
	c.respond(status226)
}
