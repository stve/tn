package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
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
			if fileExists(artworkFilename) {
				fmt.Println("setting artwork on file: " + files[i])
				err := setPicture(files[i], artworkFilename)
				if err != nil {
					log.Fatal("could not save artwork", err)
				}
			} else {
				continue
			}
		}

		return nil
	},
}

var art = &cli.Command{
	Name: "art",
	Desc: "View the artwork for a file",
	Argv: func() interface{} { return new(artworkT) },
	Fn: func(ctx *cli.Context) error {

		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			tag, err := id3v2.Open(files[i])
			if err != nil {
				return err
			}

			pictures := tag.GetFrames(tag.CommonID("Attached picture"))
			if pictures != nil {
				for _, f := range pictures {
					pic, ok := f.(id3v2.PictureFrame)
					if !ok {
						log.Fatal("Couldn't assert picture frame")
					}

					// Do some operations with picture frame:
					fmt.Println(pic.Description) // For example, print description of picture frame
					// _, rerr := ioutil.ReadAll(pic.Picture) // Or read a picture from picture frame
					// if rerr != nil {
					// 	log.Fatal("Error while reading a picture from picture frame: ", rerr)
					// }
				}
			}

			if err = tag.Close(); err != nil {
				log.Fatal("Error while closing a tag:", err)
			}
		}

		return nil
	},
}
