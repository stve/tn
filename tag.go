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
	File   string `cli:"file" usage:"A specific file to tag"`
}

var tag = &cli.Command{
	Name: "tag",
	Desc: "Set the id3 tags for mp3 files",
	Argv: func() interface{} { return new(tagT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*tagT)

		// ctx.String("auto=%v, artist=%s, album=%s, title=%s, file=%s\n", argv.Auto, argv.Artist, argv.Album, argv.Title, argv.File)

		if argv.Auto {
			autoTag()
			return nil
		}

		if checkNotEmpty(argv.File) {
			return saveData(argv.File, argv.Artist, argv.Album, argv.Title)
		}

		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {

			fmt.Println("\n************************************")
			fmt.Println("File: " + files[i] + ":")

			err := saveData(files[i], argv.Artist, argv.Album, argv.Title)
			if err != nil {
				log.Println("Error!")
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

		err := saveData(files[i], trackInfo[0], "", trackInfo[1])
		if err != nil {
			log.Println("Error!")
		}
	}
}

func saveData(filename string, artist string, album string, title string) error {
	tag, err := id3v2.Open(filename)
	if err != nil {
		mp3File, terr := id3.Open(filename)
		if terr != nil {
			log.Fatal("Error while opening mp3 file: ", terr)
		}
		defer mp3File.Close()

		if checkNotEmpty(artist) {
			mp3File.SetArtist(artist)
		}

		if checkNotEmpty(album) {
			mp3File.SetAlbum(album)
		}

		if checkNotEmpty(title) {
			mp3File.SetTitle(title)
		}

	} else {
		defer tag.Close()

		if checkNotEmpty(artist) {
			tag.SetArtist(artist)
		}

		if checkNotEmpty(album) {
			tag.SetAlbum(album)
		}

		if checkNotEmpty(title) {
			tag.SetTitle(title)
		}

		saveerr := tag.Save()
		if saveerr != nil {
			log.Fatal("error while saving mp3 file ", saveerr)
		}
	}

	return nil
}
