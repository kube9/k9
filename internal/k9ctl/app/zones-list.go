package app

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kube9/k9/pkg/gen/k9server/client"
	"github.com/kube9/k9/pkg/gen/k9server/models"
	"github.com/spf13/cobra"
)

// zonesListCmd represents the zones_list command
var zonesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all zones",
	Run: func(cmd *cobra.Command, args []string) {
		listZones(cmd, args)
	},
}

func init() {
	zonesCmd.AddCommand(zonesListCmd)
	AddMultipleOutput(zonesListCmd)
}

func listZones(cmd *cobra.Command, args []string) {
	// Get data from server
	data, err := client.Default.Zone.GetZones(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Print JSON
	if len(multipleOutput) > 0 {
		if err := OutputPrinter(data.Payload.ZoneList); err != nil {
			os.Exit(1)
		}
		return
	}

	// Print TEXT
	printZonesAsText(data.Payload.ZoneList)
}

func printZonesAsText(data models.ZoneList) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tNAME")

	for _, z := range data {
		fmt.Fprintf(w, "%s\t%s\n", *z.ID, *z.Name)
	}

	fmt.Fprintln(w)
	w.Flush()
}
