package cmd

import (
	"github.com/andygeiss/gostats/internal/stats"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(tableCmd)
}

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Print stats about a given Golang source directory as a ASCII table",
	Long:  "Print stats about a given Golang source directory as a ASCII table",
	Run: func(cmd *cobra.Command, args []string) {
		printHelpIfEmptyArgs(args)
		printTableWithResult(getResultByArg0Path(args))
	},
}

func printTableWithResult(result []stats.Statistics) {
	var c, n int
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Ident", "Nodes", "Complexity", "NCR"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
	for _, stat := range result {
		if stat.Complexity == 0 {
			stat.Complexity++
		}
		table.Append([]string{stat.Ident, strconv.Itoa(stat.Nodes), strconv.Itoa(stat.Complexity), strconv.Itoa(stat.Nodes/stat.Complexity)})
		c += stat.Complexity
		n += stat.Nodes
	}
	table.SetFooter([]string{"Total", strconv.Itoa(n), strconv.Itoa(c), strconv.Itoa(n/c)})
	table.Render()
}
