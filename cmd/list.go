package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectName string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all opened Pull Requests",
	Long:  "List command provides an list of all opened Pull Requests from Azure Devops selected Project",
	Run:   func(cmd *cobra.Command, args []string) { run(projectName) },
}

func run(projectName string) {
	url := fmt.Sprintf("https://dev.azure.com/%s/_apis/git/pullrequests?api-version=7.0", projectName)

	GetRequest(url, "crw77k4pvxjxw3cc3jek3szmbi6w2uhq5angyskijrludjgwb5fq")
}

func init() {
	listCmd.Flags().StringVarP(&projectName, "project", "p", "", "Source project to list PR's")
	rootCmd.AddCommand(listCmd)
}
