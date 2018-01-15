package app

import "github.com/spf13/cobra"

// zonesCmd represents the zones command
var zonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Manage zones",
}

func init() {
	rootCmd.AddCommand(zonesCmd)
}
