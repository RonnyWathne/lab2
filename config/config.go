package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	Number int
	Name   string
}

// Save configuration c as text file 'config.txt' in the current directory.
func (c Configuration) Save() error {
	cfile, err := os.Create("config.txt")
	if err != nil {
		// couldn't create file
		return err
	}
	// successfully created file; close it at the end of this function
	defer cfile.Close()
	_, err = fmt.Fprintf(cfile, "Number=%d\nName=%s\n", c.Number, c.Name)
	return err
}
