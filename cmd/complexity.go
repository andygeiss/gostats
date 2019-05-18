package cmd

import (
	"fmt"
	"github.com/andygeiss/gostats/internal/stats"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(complexityCmd)
}

var complexityCmd = &cobra.Command{
	Use:   "complexity",
	Short: "Print the complexity of a given Golang source directory",
	Long:  "Print the complexity of a given Golang source directory",
	Run: func(cmd *cobra.Command, args []string) {
		printHelpIfEmptyArgs(args)
		printComplexityByResult(getResultByArg0Path(args))
	},
}

func printComplexityByResult(result []stats.Statistics) {
	c := 0
	for _, stat := range result {
		c += stat.Complexity
	}
	fmt.Printf("%d\n", c)
}
