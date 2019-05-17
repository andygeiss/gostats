package main

import (
	"fmt"
	"github.com/andygeiss/gostats/internal/stats"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <path>\n\nPrint statistics about a specific Golang source directory.\n\n", os.Args[0])
		os.Exit(0)
	}

	path := os.Args[1]
	result, err := stats.NewDefaultService(path).Measure()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(255)
	}

	c := 0
	n := 0
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Ident", "Nodes", "Complexity"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})

	for _, stat := range result {
		table.Append([]string{stat.Ident, strconv.Itoa(stat.Nodes), strconv.Itoa(stat.Complexity)})
		c += stat.Complexity
		n += stat.Nodes
	}

	table.SetFooter([]string{"Total", strconv.Itoa(n), strconv.Itoa(c)})

	table.Render()
}
