package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

// OutputFormat defines the way that the results of a command are output to the user.
type OutputFormat string

const (
	// OutputFormatTable represents the output formatted in a table.
	OutputFormatTable OutputFormat = "table"

	// OutputFormatJSON represents the output formatted as JSON.
	OutputFormatJSON OutputFormat = "json"
)

// AvailableOutputFormatStrings returns all the output formats available to users.
var AvailableOutputFormatStrings = []string{string(OutputFormatTable), string(OutputFormatJSON)}

// GetOutputFormat gets the selected output format based on the CLI args.
func GetOutputFormat(cliContext *cli.Context) (OutputFormat, error) {
	format := cliContext.String(FlagOutputFormat.Name)
	if format == "" || strings.EqualFold(format, string(OutputFormatTable)) {
		return OutputFormatTable, nil
	}

	if strings.EqualFold(format, string(OutputFormatJSON)) {
		return OutputFormatJSON, nil
	}

	return OutputFormatTable, fmt.Errorf("unknown output format: %s", format)
}

// OutputTable outputs the specified data as a table.
func OutputTable(data [][]string) error {
	return pterm.
		DefaultTable.
		WithHasHeader().
		WithData(data).
		Render()
}

// OutputJSON outputs the specified object as JSON.
func OutputJSON(v interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(v); err != nil {
		return err
	}

	return nil
}
