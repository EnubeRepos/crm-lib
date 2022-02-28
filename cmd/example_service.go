package cmd

import (
	"fmt"

	"github.com/EnubeRepos/crm-lib/apicrm/account"
	"github.com/EnubeRepos/crm-lib/internal/crmapi"
	"github.com/spf13/cobra"
)

var example = &cobra.Command{
	Use:   "example",
	Short: "Start Example application",
	RunE: func(cmd *cobra.Command, args []string) error {

		host := "https://app.10maisbank.com.br/api/v1/"
		client := account.NewAPIAccountService(crmapi.New(host, "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="))

		res, _ := client.GetAccount()

		fmt.Println(res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(example)
}
