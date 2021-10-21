package main

import (
	"fmt"
	"os"

	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loopDir(startingDir string) {
	c, err := os.ReadDir(startingDir)
	check(err)

	for _, entry := range c {
		if entry.IsDir() {
			if entry.Name() == "Solved" {
				fmt.Println("deleting: " + startingDir + "/" + entry.Name())
				defer os.RemoveAll(startingDir + "/" + entry.Name())
			} else {
				loopDir(startingDir + "/" + entry.Name())
			}
		}
	}
}
func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	loopDir(filepath.Dir(ex))

}
