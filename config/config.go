package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strings"
	"encoding/gob"
)

type Configuration struct {
	Number int
	Name   string
}

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

func LoadConfig(file string) (Configuration, error) {
	newconf := Configuration{0, ""}
	file, err := os.Open(file) //open file in read mode
	if err != nil {
		return newconf, err //stop end return error if Open went wrong
	}
	defer file.Close()
	bufferedReader := bufio.NewReader(file) //create a buffered reader wrapping file
	var err1 error = nil
	for err == nil && err1 == nil { //This for-loop will break when an error occures
		var line string
		line, err = bufferedReader.ReadString('\n') // read until next \n.
		// Returns error (io.EOF), when end of file is reached.
		if line != "" {
			err1 = parseline(line, &newconf)
		}
	}
	if err != io.EOF { //something went wrong when reading
		return newconf, err
	}
	return newconf, err1
}

func parseline(line string, newconf *Configuration) (err error) {
	err = nil
	//parse the lines into the Configuration.
	//You can e.g. use: strings.HasPrefix(line, "somestring)
	//To switch bewtween name and number and
	// fmt.Sscanf(line, "format", &newconf.Myfield)
	// to parse the string.
	return
}

func (C Configuration) MkGobConf() error {
	cfile, err := os.Create("config.gob")
	if err != nil {
		return err //stop end return error of Create went wrong
	}
	defer cfile.Close()              //close the file at the end of this function
	encoder := gob.NewEncoder(cfile) //create a gob encoder
	err = encoder.Encode(C)
	return err
}

func GetGobConf(file string) (conf Configuration, err error) {
	conf = Configuration{0, ""}
	err = nil
	//Open the file using os.Open.
	//Decode using a gob Decoder.
	return
}
