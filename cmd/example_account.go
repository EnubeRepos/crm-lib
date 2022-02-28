package cmd

import (
	"fmt"

	"github.com/EnubeRepos/crm-lib/internal/core/services/accountsrv"
	"github.com/EnubeRepos/crm-lib/internal/crmapi"
	"github.com/spf13/cobra"
)

var workerAccount = &cobra.Command{
	Use:   "account",
	Short: "Start Example application",
	RunE: func(cmd *cobra.Command, args []string) error {

		srv := accountsrv.NewAPIAccountService(crmapi.CRMAPIClient{
			BaseURL: "https://ssaarg-dev.cloudanalytics.me/api/v1/",
			Token:   "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug==",
		})

		result, _ := srv.GetAccount()

		fmt.Println(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(workerAccount)
}
