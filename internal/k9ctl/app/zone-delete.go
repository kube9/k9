package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/kube9/k9/pkg/gen/k9server/client"
	"github.com/kube9/k9/pkg/gen/k9server/client/zone"
	"github.com/spf13/cobra"
)

// zonesDeleteCmd represents the zones_delete command
var zonesDeleteCmd = &cobra.Command{
	Use:   "delete {zone-id}",
	Short: "Delete a zone",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		deleteZone(cmd, args)
	},
}

func init() {
	zonesCmd.AddCommand(zonesDeleteCmd)
}

func deleteZone(cmd *cobra.Command, args []string) {
	id := args[0]
	// Get data from server
	_, err := client.Default.Zone.DeleteZone(zone.NewDeleteZoneParams().WithZoneID(id))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
