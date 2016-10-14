package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
	"github.com/mikkyang/id3-go"
	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(tag),
		cli.Tree(tags),
		cli.Tree(cover),
		cli.Tree(artwork),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help information")

// root command
type rootT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

var root = &cli.Command{
	Desc: "this is root command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}

// child command
type childT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

// tag command

type tagT struct {
	cli.Helper
	Auto   bool   `cli:"auto" usage:"short and long format flags both are supported"`
	Artist string `cli:"artist" usage:"Artist name"`
	Album  string `cli:"album" usage:"Album name"`
	Title  string `cli:"title" usage:"Song name"`
}

var tag = &cli.Command{
	Name: "tag",
	Desc: "tag command",
	Argv: func() interface{} { return new(tagT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*tagT)

		ctx.String("auto=%v, artist=%s, album=%s, title=%s\n", argv.Auto, argv.Artist, argv.Album, argv.Title)

		// if auto flag
		// use filename
		if argv.Auto {
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
				} else {
					defer tag.Close()

				}
			}

		} else {
			// Open file and find tag in it
			tag, err := id3v2.Open("file.mp3")
			if err != nil {
				log.Fatal("Error while opening mp3 file: ", err)
			}
			defer tag.Close()

			// if artist flag
			tag.SetArtist(argv.Artist)

			// if album flag
			tag.SetAlbum(argv.Album)

			// if song flag
			tag.SetTitle(argv.Title)

			if err = tag.Save(); err != nil {
				log.Fatal("Error while saving a tag: ", err)
			}
		}

		return nil
	},
}

// tags command

type tagsT struct {
	cli.Helper
}

var tags = &cli.Command{
	Name: "tags",
	Desc: "tags command",
	Argv: func() interface{} { return new(tagsT) },
	Fn: func(ctx *cli.Context) error {

		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {
			fmt.Println("\n************************************")
			fmt.Println("File: " + files[i] + ":")

			// Open file and find tag in it
			tag, err := id3v2.Open(files[i])
			if err != nil {
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

var artwork = &cli.Command{
	Name: "artwork",
	Desc: "artwork command",
	Argv: func() interface{} { return new(childT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childT)
		ctx.String("Hello, child command, I am %s\n", argv.Name)
		return nil
	},
}

var cover = &cli.Command{
	Name: "cover",
	Desc: "cover command",
	Argv: func() interface{} { return new(childT) },
	Fn: func(ctx *cli.Context) error {

		argv := ctx.Argv().(*childT)
		ctx.String("Hello, child command, I am %s\n", argv.Name)
		return nil
	},
}
