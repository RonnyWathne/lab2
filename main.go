// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"os"

	"github.com/uis-dat320-fall2014/lab2/config"
)

func main() {
	cfg1 := config.Configuration{1, "hello"}
	if err := cfg1.Save(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg2 := config.Configuration{2, "lab"}
	if err := cfg2.SaveGob(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(cfg1)
	fmt.Println(cfg2)

	// TODO: Load Configuration objects back from disk
}
