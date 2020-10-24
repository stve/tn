package tags

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bogem/id3v2"
	"github.com/mikkyang/id3-go"
)

func List(displayArtwork bool) {
	files, _ := filepath.Glob("*.mp3")
	for i := 0; i < len(files); i++ {
		fmt.Println("\n************************************")
		fmt.Println("File: " + files[i] + ":")

		// Open file and find tag in it
		tag, err := id3v2.Open(files[i], id3v2.Options{Parse: true})
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

			if displayArtwork {
				pictures := tag.GetFrames(tag.CommonID("Attached picture"))
				if pictures != nil {
					for _, f := range pictures {
						pic, ok := f.(id3v2.PictureFrame)
						if !ok {
							log.Fatal("Couldn't assert picture frame")
						}

						// Do some operations with picture frame:
						fmt.Println("  Cover: " + pic.Description) // For example, print description of picture frame
						// _, rerr := ioutil.ReadAll(pic.Picture) // Or read a picture from picture frame
						// if rerr != nil {
						// 	log.Fatal("Error while reading a picture from picture frame: ", rerr)
						// }
					}
				}
			}
		}
	}
}
