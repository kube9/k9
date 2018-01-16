package app

import (
	"fmt"

	"github.com/kube9/k9/pkg/gen/k9server/client"
	"github.com/kube9/k9/pkg/version"
	"github.com/spf13/cobra"
)

const notResponding = "Not responding"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		getVersion(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("Client: %s\n", getClientVersionAsString())
	fmt.Printf("Server: %s\n", getServerVersionAsString())
}

func getClientVersionAsString() string {
	return version.Version
}

func getServerVersionAsString() string {
	resp, err := client.Default.Info.GetVersion(nil)
	if err != nil {
		return notResponding
	}
	return resp.Payload
}
