package cmd

import (
	"fmt"
	"github.com/andygeiss/gostats/internal/stats"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ncrCmd)
}

var ncrCmd = &cobra.Command{
	Use:   "ncr",
	Short: "Measures the number of nodes per complexity",
	Long:  "Measures the number of nodes per complexity",
	Run: func(cmd *cobra.Command, args []string) {
		printHelpIfEmptyArgs(args)
		printComplexityByResult(getResultByArg0Path(args))
	},
}

func printNodeComplexityRateByResult(result []stats.Statistics) {
	c := 0
	n := 0
	for _, stat := range result {
		c += stat.Complexity
		n += stat.Nodes
	}
	fmt.Printf("%d\n", n/c)
}
