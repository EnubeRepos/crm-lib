package cmd

import (
	"log"

	"github.com/EnubeRepos/CRM_framework/internal/context"
	"github.com/spf13/cobra"
)

var lead = &cobra.Command{
	Use:   "lead",
	Short: "Start Import Lead application",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx, err := context.NewConfig()

		if err != nil {
			return err
		}

		log.Println(ctx.Config.API_PARTNER_REFERRALS)

		log.Println("Started Get Lead")

		// for _, host := range a.Config.HOST_API_CRM {
		// 	a.Config.HOST_CRM = host
		// 	ctx.Logger.Println("HOST Session : ", a.Config.HOST_CRM)

		// 	listOfDistributors, err := a.CRMGetDistributor(ctx)
		// 	if err != nil {
		// 		ctx.Logger.Error("Error getting list of distributors. Error: ", err)
		// 		continue
		// 	}

		// 	for _, distributor := range listOfDistributors.Distributor {
		// 		leads, _, _ := a.REFGetLead(ctx, distributor.ID)
		// 		ctx.Logger.Println(leads)

		// 		for _, lead := range leads.Value {
		// 			ctx.Logger.Println(lead)

		// 		}

		// 	}
		// }
		log.Println("Finished")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lead)
}
