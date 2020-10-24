package tags

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
	"github.com/mikkyang/id3-go"
)

// Use filename to derive the artist and title
func AutoTag() {
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

		err := SaveData(files[i], trackInfo[0], "", trackInfo[1])
		if err != nil {
			log.Println("Error!")
		}
	}
}

func SaveData(filename string, artist string, album string, title string) error {
	tag, err := id3v2.Open(filename, id3v2.Options{Parse: false})
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
