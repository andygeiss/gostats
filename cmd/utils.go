package cmd

import (
	"github.com/andygeiss/check"
	"github.com/andygeiss/gostats/internal/stats"
	"os"
)

func getResultByArg0Path(args []string) []stats.Statistics {
	result, err := stats.NewDefaultService(args[0]).Measure()
	check.Fatal(err)
	return result
}

func printHelpIfEmptyArgs(args []string) {
	if len(args) == 0 {
		check.Fatal(rootCmd.Help())
		os.Exit(0)
	}
}
