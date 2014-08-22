// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"os"

	"github.com/uis-dat320-fall2014/lab2/config"
)

func main() {
	cfg := config.Configuration{1, "hello"}
	if err := cfg.Save(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := cfg.SaveGob(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("looking good", cfg)
}
