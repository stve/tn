package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/mkideal/cli"
)

type coverT struct {
	cli.Helper
	File string `cli:"file" usage:"A specific image file to use as the cover"`
}

const defaultCoverImage string = "cover.jpg"

var cover = &cli.Command{
	Name: "cover",
	Desc: "Update the artwork for all files with cover.jpg",
	Argv: func() interface{} { return new(coverT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*coverT)

		var coverImage string

		if checkNotEmpty(argv.File) {
			coverImage = argv.File
		} else {
			coverImage = defaultCoverImage
		}

		if fileExists(coverImage) {
			files, _ := filepath.Glob("*.mp3")
			for i := 0; i < len(files); i++ {
				fmt.Println("setting cover on file: " + files[i])
				err := setPicture(coverImage, files[i])
				if err != nil {
					log.Fatal("could not save cover", err)
				}
			}
		} else {
			log.Fatal("could not find file: " + coverImage)
		}

		return nil
	},
}
