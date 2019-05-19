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
	Short: "Measures the number of linearly independent paths rough a source code",
	Long:  "Measures the number of linearly independent paths rough a source code",
	Run: func(cmd *cobra.Command, args []string) {
		printHelpIfEmptyArgs(args)
		printComplexityByResult(getResultByArg0Path(args))
	},
}

func printComplexityByResult(result []stats.Statistics) {
	num := 0
	for _, stat := range result {
		num += stat.Complexity
	}
	fmt.Printf("%d\n", num)
}
