package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stve/tn/tags"
)

func init() {
	rootCmd.AddCommand(artworkCmd)
}

var artworkCmd = &cobra.Command{
	Use:   "artwork",
	Short: "Set the artwork for a file",
	Long:  `Set the artwork for a file`,
	Run: func(cmd *cobra.Command, args []string) {
		tags.Artwork()
	},
}
