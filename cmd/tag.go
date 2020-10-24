package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stve/tn/tags"
)

func init() {
	tagCmd.Flags().BoolVarP(&autoTag, "auto", "", false, "Autotag all files in the current directory by filename")
	tagCmd.Flags().StringVarP(&artist, "artist", "", "", "Set the artist name")
	tagCmd.Flags().StringVarP(&album, "album", "", "", "Set the album name")
	tagCmd.Flags().StringVarP(&title, "title", "", "", "Set the song title")
	tagCmd.Flags().StringVarP(&file, "file", "", "", "A specific file to tag")

	rootCmd.AddCommand(tagCmd)
}

var autoTag bool
var artist string
var album string
var title string
var file string

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Show currently set tags for all mp3 files",
	Run: func(cmd *cobra.Command, args []string) {
		if autoTag {
			tags.AutoTag()
			return
		}

		if file != "" {
			tags.SaveData(file, artist, album, title)
		}

		files, _ := filepath.Glob("*.mp3")
		for i := 0; i < len(files); i++ {

			fmt.Println("\n************************************")
			fmt.Println("File: " + files[i] + ":")

			err := tags.SaveData(files[i], artist, album, title)
			if err != nil {
				log.Println("Error!")
			}

		}
	},
}
