package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/jsonpath"
)

var multipleOutput string

//AddMultipleOutput adds the -o|--output option to any cmd to export to json
func AddMultipleOutput(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&multipleOutput, "output", "o", "", "json| jsonpath='{}'")
}

//OutputPrinter receives an interface and dump the data using the --output flag.
//ATM only json or jsonpath. In the future yaml
func OutputPrinter(data interface{}) error {
	var re = regexp.MustCompile(`^jsonpath\=(.*)`)

	if multipleOutput == "json" {
		return dumpJSON(data, "")
	}

	if re.MatchString(multipleOutput) {
		return dumpJSON(data, re.ReplaceAllString(multipleOutput, "$1"))
	}

	return fmt.Errorf("Couldn't found output printer")
}

// dumpJSON dump the data variable to the stdout as json.
// If somethings fail, it'll return an error
// If jsonPath is passed, it'll run the json query over data var.
func dumpJSON(data interface{}, jsonPath string) error {

	if len(jsonPath) == 0 {
		result, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't marshal to json: '%s'\n", err)
			return err
		}
		fmt.Println(string(result))
		return nil
	}

	parser := jsonpath.New("").AllowMissingKeys(true)
	if err := parser.Parse(jsonPath); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse jsonpath expression: '%s'\n", err)
		return err
	}

	buf := new(bytes.Buffer)
	if err := parser.Execute(buf, data); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse jsonpath expression: '%s'\n", err)
		return err

	}
	fmt.Println(buf.String())
	return nil
}
