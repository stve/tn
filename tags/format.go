package tags

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var fileFormats = [2]string{".wav", ".aif"}

func filesOfType(dir string, ext string, overwriteExisting bool) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var filesToConvert []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.Contains(file.Name(), ext) {
			continue
		}

		filePath := path.Join(dir, file.Name())
		if _, err := os.Stat(filePath); err != nil {
			continue
		}

		// If force is true, we want to convert regardless of if a file exists or not
		if !overwriteExisting {
			if _, err := os.Stat(strings.Replace(filePath, ext, ".mp3", 1)); err == nil {
				fmt.Printf("%s already exists, skipping...\n", file.Name())
				continue
			}
		}

		filesToConvert = append(filesToConvert, filePath)
	}

	return filesToConvert
}

func convert(path string, wg *sync.WaitGroup) {
	extension := filepath.Ext(path)
	newPath := strings.Replace(path, extension, ".mp3", 1)
	cmd := exec.Command("lame", "--silent", "-b 320", "-h", "-V2", path, newPath)

	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Build:", err)
		fmt.Println("Build:", string(output))
		os.Exit(1)
	}

	fmt.Printf("Done converting '%s' \n", path)
	wg.Done()
}

// Format - a command to format files
func Format(overwriteExisting bool) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	var files []string
	for _, format := range fileFormats {
		fmt.Printf("Detecting %s files...\n", format)

		filesToConvert := filesOfType(dir, format, overwriteExisting)
		if len(filesToConvert) > 0 {
			files = append(files, filesToConvert...)
		}
	}

	if len(files) == 0 {
		fmt.Println("No files were found, exiting...")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(files))
	for _, p := range files {
		go convert(p, &wg)
	}

	wg.Wait()
}
