package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/mkideal/cli"
)

type coverT struct {
	cli.Helper
}

var cover = &cli.Command{
	Name: "cover",
	Desc: "Update the artwork for all files with cover.jpg",
	Argv: func() interface{} { return new(coverT) },
	Fn: func(ctx *cli.Context) error {
		if fileExists("cover.jpg") {
			files, _ := filepath.Glob("*.mp3")
			for i := 0; i < len(files); i++ {
				fmt.Println("setting cover on file: " + files[i])
				err := setPicture("cover.jpg", files[i])
				if err != nil {
					log.Fatal("could not save cover", err)
				}
			}
		} else {
			log.Fatal("could not find file: cover.jpg")
		}

		return nil
	},
}
