package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
	id3 "github.com/mikkyang/id3-go"
	"github.com/mkideal/cli"
)

type tagT struct {
	cli.Helper
	Auto   bool   `cli:"auto" usage:"short and long format flags both are supported"`
	Artist string `cli:"artist" usage:"Artist name"`
	Album  string `cli:"album" usage:"Album name"`
	Title  string `cli:"title" usage:"Song name"`
}

var tag = &cli.Command{
	Name: "tag",
	Desc: "Set the id3 tags for mp3 files",
	Argv: func() interface{} { return new(tagT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*tagT)

		ctx.String("auto=%v, artist=%s, album=%s, title=%s\n", argv.Auto, argv.Artist, argv.Album, argv.Title)

		if argv.Auto {
			autoTag()
		} else {
			// Open file and find tag in it
			tag, err := id3v2.Open("file.mp3")
			if err != nil {
				log.Fatal("Error while opening mp3 file: ", err)
			}
			defer tag.Close()

			// if artist flag
			if checkNotEmpty(argv.Artist) {
				tag.SetArtist(argv.Artist)
			}

			// if album flag
			if checkNotEmpty(argv.Album) {
				tag.SetAlbum(argv.Album)
			}

			// if song flag
			if checkNotEmpty(argv.Title) {
				tag.SetTitle(argv.Title)
			}

			sverr := tag.Save()
			if sverr != nil {
				log.Fatal("error while saving mp3 file ", sverr)
			}
		}

		return nil
	},
}

// Use filename to derive the artist and title
func autoTag() {
	files, _ := filepath.Glob("*.mp3")
	for i := 0; i < len(files); i++ {

		fmt.Println("\n************************************")
		fmt.Println("File: " + files[i] + ":")

		if strings.Index(files[i], " - ") == -1 {
			fmt.Println("Unable to determine track name/artist, skipping")
			continue
		}

		info := strings.Replace(files[i], ".mp3", "", 1)
		trackInfo := strings.Split(info, " - ")

		fmt.Println("  Artist: " + trackInfo[0])
		fmt.Println("  Title: " + trackInfo[1])

		tag, err := id3v2.Open(files[i])
		if err != nil {
			mp3File, terr := id3.Open(files[i])
			if terr != nil {
				log.Fatal("Error while opening mp3 file: ", terr)
			}
			defer mp3File.Close()

			mp3File.SetArtist(trackInfo[0])
			mp3File.SetTitle(trackInfo[1])

			// _, serr := mp3File.Save()
			// if serr != nil {
			// 	log.Fatal("error while saving mp3 file ", serr)
			// }
		} else {
			defer tag.Close()

			tag.SetArtist(trackInfo[0])
			tag.SetTitle(trackInfo[1])

			saveerr := tag.Save()
			if saveerr != nil {
				log.Fatal("error while saving mp3 file ", saveerr)
			}

		}
	}
}

func checkNotEmpty(str string) bool {
	return (len(str) > 0)
}
