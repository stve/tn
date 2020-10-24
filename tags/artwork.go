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

func Cover(filename string) {
	var coverImage string
	if checkNotEmpty(filename) {
		coverImage = filename
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
}
