package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stve/tn/tags"
)

func init() {
	rootCmd.AddCommand(clearCmd)
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all id3 tags",
	Run: func(cmd *cobra.Command, args []string) {
		tags.Clear()
	},
}
