package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azure-pullrequests",
	Short: "Azure Pullrequests provide Azure Devops PR managment",
	Long:  "Azure Pullrequests are a short and easy way to Azure Devops user manage opened Pull Requests",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
