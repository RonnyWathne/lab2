package main

import (
	"fmt"
	"os"

	"github.com/uis-dat320-fall2014/lab2/config"
)

func main() {
	newconfig := config.Configuration{1, "hello"}
	err := newconfig.Save()
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
