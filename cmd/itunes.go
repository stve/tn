package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stve/tn/tags"
)

func init() {
	rootCmd.AddCommand(itunesCmd)
}

var itunesCmd = &cobra.Command{
	Use:   "itunes",
	Short: "Add mp3s to iTunes",
	Run: func(cmd *cobra.Command, args []string) {
		tags.Itunes()
	},
}
