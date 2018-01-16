package app

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kube9/k9/pkg/gen/k9server/client"
	"github.com/kube9/k9/pkg/gen/k9server/client/zone"
	"github.com/kube9/k9/pkg/gen/k9server/models"
	"github.com/spf13/cobra"
)

// zonesCreateCmd represents the zones_list command
var zonesCreateCmd = &cobra.Command{
	Use:   "create {zone-name}",
	Short: "Create a new zone",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		createZone(cmd, args)
	},
}

func init() {
	zonesCmd.AddCommand(zonesCreateCmd)
	AddMultipleOutput(zonesCreateCmd)
}

func createZone(cmd *cobra.Command, args []string) {
	name := &args[0]
	newZone := &models.CreateZoneRequest{
		NewZone: models.NewZone{
			Name: name,
		},
	}
	// Get data from server
	data, err := client.Default.Zone.CreateZone(zone.NewCreateZoneParams().WithZone(newZone))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Print JSON
	if len(multipleOutput) > 0 {
		if err := OutputPrinter(data.Payload.Zone); err != nil {
			os.Exit(1)
		}
		return
	}

	// Print TEXT
	printZoneAsText(data.Payload.Zone)
}

func printZoneAsText(data models.Zone) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tNAME")

	fmt.Fprintf(w, "%s\t%s\n", *data.ID, *data.Name)

	fmt.Fprintln(w)
	w.Flush()
}
