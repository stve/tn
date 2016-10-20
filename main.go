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
	Desc: "Set the id3 tags for mp3 files",
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

// tags command

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

// artwork command

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
				err := setPictureTag(artworkFilename, files[i])
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

// cover command

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
				err := setPictureTag("cover.jpg", files[i])
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

// helper functions

func setPictureTag(coverFilename string, filename string) error {
	frontCover, err := os.Open("cover.jpg")
	if err != nil {
		return err
	}

	defer frontCover.Close()

	pic := id3v2.PictureFrame{
		Encoding:    id3v2.ENUTF8,
		MimeType:    "image/jpeg",
		Picture:     frontCover,
		PictureType: id3v2.PTFrontCover,
	}

	// Open file and find tag in it
	tag, err := id3v2.Open(filename)
	if err != nil {
		return err
	}

	defer tag.Close()
	tag.DeleteAllFrames()
	tag.AddAttachedPicture(pic)

	fmt.Println("picture saved")

	err = tag.Save()

	if err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return true
	}

	return false
}

func checkNotEmpty(str string) bool {
	if len(str) > 0 {
		return true
	}

	return false
}
