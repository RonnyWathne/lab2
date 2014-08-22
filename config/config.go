package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strings"
)

type Configuration struct {
	Number int
	Name   string
}

// Save configuration c as a text file 'config.txt' in the current directory.
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

// LoadConfig loads a text-based configuration file, and returns the
// corresponding Configuration object.
func LoadConfig(file string) (Configuration, error) {
	conf := Configuration{0, ""}
	f, err := os.Open(file)
	if err != nil {
		// couldn't open file for reading
		return conf, err
	}
	// successfully opened file; close it at the end of this function
	defer f.Close()
	bufferedReader := bufio.NewReader(f) //create a buffered reader wrapping file
	var err1 error = nil
	for err == nil && err1 == nil { //This for-loop will break when an error occures
		var line string
		line, err = bufferedReader.ReadString('\n') // read until next \n.
		// Returns error (io.EOF), when end of file is reached.
		if line != "" {
			err1 = parseline(line, &conf)
		}
	}
	if err != io.EOF { //something went wrong when reading
		return conf, err
	}
	return conf, err1
}

func parseline(line string, conf *Configuration) (err error) {
	err = nil
	//parse the lines into the Configuration.
	//You can e.g. use: strings.HasPrefix(line, "somestring)
	//To switch bewtween name and number and
	// fmt.Sscanf(line, "format", &conf.Myfield)
	// to parse the string.
	return
}
