package config

import (
	"encoding/gob"
	"os"
)

func (c Configuration) SaveGob() error {
	cfile, err := os.Create("config.gob")
	if err != nil {
		// couldn't create file
		return err
	}
	// successfully created file; close it at the end of this function
	defer cfile.Close()
	// create a gob encoder for the file
	encoder := gob.NewEncoder(cfile)
	err = encoder.Encode(c)
	return err
}

func LoadGobConfig(file string) (conf *Configuration, err error) {
	conf = &Configuration{0, ""}
	err = nil
	// TODO: Open file using os.Open()
	// TODO: Decode using a gob decoder.
	return
}
