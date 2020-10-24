package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stve/tn/tags"
)

func init() {
	rootCmd.AddCommand(coverCmd)
}

var coverCmd = &cobra.Command{
	Use:   "cover",
	Short: "Update the artwork for all files with cover.jpg",
	Run: func(cmd *cobra.Command, args []string) {
		tags.Cover("")
	},
}
