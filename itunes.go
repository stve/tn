package main

import (
	"fmt"
	"path/filepath"

	"github.com/everdev/mack"
	"github.com/mkideal/cli"
)

type itunesT struct {
	cli.Helper
}

var itunes = &cli.Command{
	Name: "itunes",
	Desc: "Add mp3s to iTunes",
	Argv: func() interface{} { return new(itunesT) },
	Fn: func(ctx *cli.Context) error {
		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			fmt.Println("Adding " + files[i])
			err := mack.Tell("iTunes", "add (POSIX file \""+files[i]+"\"")
			if err != nil {
				fmt.Println("Error adding " + files[i])
			}
		}

		return nil
	},
}
