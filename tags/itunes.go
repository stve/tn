package tags

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/everdev/mack"
)

func Itunes() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files, _ := filepath.Glob("*.mp3")
	for i := 0; i < len(files); i++ {
		command := fmt.Sprintf("add (POSIX file \"%s\")", filepath.Join(pwd, files[i]))
		fmt.Println("Adding " + command)

		_, err := mack.Tell("iTunes", command)
		if err != nil {
			fmt.Println("Error adding "+files[i], err)
		}
	}
}
