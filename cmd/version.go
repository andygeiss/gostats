package cmd

import (
	"fmt"
	"github.com/andygeiss/gostats/configs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version identifier",
	Long:  "Print the version identifier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(configs.AppInfo.Version)
	},
}
