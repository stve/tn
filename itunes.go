package main

import (
	"fmt"
	"os"
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
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			command := fmt.Sprintf("add (POSIX file \"%s\")", filepath.Join(pwd, files[i]))
			fmt.Println("Adding " + command)
			err := mack.Tell("iTunes", command)
			if err != nil {
				fmt.Println("Error adding "+files[i], err)
			}
		}

		return nil
	},
}
