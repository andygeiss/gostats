package cmd

import (
	"fmt"
	"github.com/andygeiss/gostats/internal/stats"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nodesCmd)
}

var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Measures the number of nodes created in the abstract syntax tree of a source code",
	Long:  "Measures the number of nodes created in the abstract syntax tree of a source code",
	Run: func(cmd *cobra.Command, args []string) {
		printHelpIfEmptyArgs(args)
		printNodesByResult(getResultByArg0Path(args))
	},
}

func printNodesByResult(result []stats.Statistics) {
	num := 0
	for _, stat := range result {
		num += stat.Nodes
	}
	fmt.Printf("%d\n", num)
}
