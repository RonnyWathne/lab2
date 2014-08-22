package main

import (
	"fmt"
	"config"
	"os"
)

func main() {
	newconfig := config.Configuration{1, "hello"}
	err := newconfig.MkConfFile()
	errcheck(err)
	err = newconfig.MkGobConf()
	errcheck(err)
	fmt.Println("looking good", newconfig)
}

func errcheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
