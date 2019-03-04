package main

import (
	. "is_git"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("Path not provided")
		os.Exit(1)
	}

	path := os.Args[1]
	if is, err := IsGitDir(path); err != nil {
		fmt.Printf("Failed to check if is git dir: %s", err.Error())
		os.Exit(1)
	} else if is {
		fmt.Println(`is_git`)
	} else {
		fmt.Println(`not_git`)
	}
}

