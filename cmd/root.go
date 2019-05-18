package cmd

import (
	"fmt"
	"github.com/andygeiss/gostats/configs"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   configs.AppInfo.Name,
	Short: configs.AppInfo.Short,
	Long:  configs.AppInfo.Long,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
