package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/mkideal/cli"
)

type artworkT struct {
	cli.Helper
}

var artwork = &cli.Command{
	Name: "artwork",
	Desc: "Set the artwork for a file",
	Argv: func() interface{} { return new(artworkT) },
	Fn: func(ctx *cli.Context) error {
		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			artworkFilename := strings.Replace(files[i], ".mp3", ".jpg", 1)

			if !(fileExists(artworkFilename)) {
				continue
			}

			fmt.Println("setting artwork on file: " + files[i])
			err := setPicture(artworkFilename, files[i])
			if err != nil {
				log.Fatal("could not save artwork", err)
			}
		}

		return nil
	},
}
