// +build !solution

// Leave an empty line above this comment.
package config

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// LoadConfig loads a text-based configuration file, and returns the
// corresponding Configuration object.
func LoadConfig(file string) (*Configuration, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	// create empty Configuration object
	conf := &Configuration{}
	// TODO: Parse each line in lines
	fmt.Println(lines)
	return conf, nil
}

func (c *Configuration) parse(line string) (err error) {
	// TODO: find keys=value in line
	//       and store value to the correct part of the c object
	// TIPS: strconv.Atoi()
	return
}
