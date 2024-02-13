package tags

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

const defaultCoverImage string = "cover.jpg"

func Artwork() {
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
}

// Cover sets the artwork of all mp3s in the current directory to the contents of cover.jpg
func Cover(filename string) {
	coverImage := defaultCoverImage

	if checkNotEmpty(filename) {
		coverImage = filename
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
}
