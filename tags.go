package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bogem/id3v2"
	id3 "github.com/mikkyang/id3-go"
	"github.com/mkideal/cli"
)

type tagsT struct {
	cli.Helper
}

var tags = &cli.Command{
	Name: "tags",
	Desc: "Show currently set tags for all mp3 files",
	Argv: func() interface{} { return new(tagsT) },
	Fn: func(ctx *cli.Context) error {

		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			fmt.Println("\n************************************")
			fmt.Println("File: " + files[i] + ":")

			// Open file and find tag in it
			tag, err := id3v2.Open(files[i])
			if err != nil {
				fmt.Println("id3v2 failed, falling back")
				mp3File, terr := id3.Open(files[i])
				if terr != nil {
					log.Fatal("Error while opening mp3 file: ", terr)
				}
				defer mp3File.Close()

				// Read tags
				fmt.Println("  Artist: " + mp3File.Artist())
				fmt.Println("  Title: " + mp3File.Title())
				fmt.Println("  Album: " + mp3File.Album())
			} else {
				defer tag.Close()

				// Read tags
				fmt.Println("  Artist: " + tag.Artist())
				fmt.Println("  Title: " + tag.Title())
				fmt.Println("  Album: " + tag.Album())
			}
		}

		return nil
	},
}
