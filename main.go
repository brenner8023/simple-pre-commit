package main

import (
	"fmt"
	"os"
	core "simple-pre-commit/src"
)

func main() {
	err := core.CheckConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gitRoot, err := core.GetGitProjectRoot()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = core.SetHook(gitRoot)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
