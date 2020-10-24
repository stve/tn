package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stve/tn/tags"
)

func init() {
	tagsCmd.Flags().BoolVarP(&displayArtwork, "artwork", "a", false, "Include artwork in output (when present)")

	rootCmd.AddCommand(tagsCmd)
}

var displayArtwork bool

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Show currently set tags for all mp3 files",
	Run: func(cmd *cobra.Command, args []string) {
		tags.List(displayArtwork)
	},
}
