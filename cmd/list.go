package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Value struct {
  Url string `json:"url"`
  Status string `json:"status"`
  Title string `json:"title"`
}

type ResponseBody struct {
  Count int `json:"count"`
  Value []Value `json:"value"`
}

var projectName string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all opened Pull Requests",
	Long:  "List command provides an list of all opened Pull Requests from Azure Devops selected Project",
	Run:   func(cmd *cobra.Command, args []string) { run(projectName) },
}

func run(projectName string) {
	url := fmt.Sprintf("https://dev.azure.com/wilzkelvin/%s/_apis/git/pullrequests?api-version=7.0", projectName)

  bodyJSON := ResponseBody{}
  body, err :=	GetRequest(url, ":dh5fuwjelg6l33fagckzmn3jwmcvyian4x6lmcnzfrwdthg2p5nq")
  if err != nil {
    os.Exit(1)
  }
  
  json.Unmarshal([]byte(body), &bodyJSON)
  fmt.Println(bodyJSON)
}

func init() {
	listCmd.Flags().StringVarP(&projectName, "project", "p", "", "Source project to list PR's")
	rootCmd.AddCommand(listCmd)
}
