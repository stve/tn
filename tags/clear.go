package tags

import (
	"log"
	"path/filepath"
)

func Clear() {
	files, _ := filepath.Glob("*.mp3")
	for i := 0; i < len(files); i++ {
		err := clearTags(files[i])
		if err != nil {
			log.Fatal("could not reset file", err)
		}
	}
}
