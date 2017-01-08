package main

import (
	"log"
	"path/filepath"

	"github.com/mkideal/cli"
)

type clearT struct {
	cli.Helper
}

var clear = &cli.Command{
	Name: "clear",
	Desc: "Clear all id3 tags",
	Argv: func() interface{} { return new(clearT) },
	Fn: func(ctx *cli.Context) error {
		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			err := clearTags(files[i])
			if err != nil {
				log.Fatal("could not reset file", err)
			}
		}

		return nil
	},
}
