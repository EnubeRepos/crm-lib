package cmd

import "github.com/spf13/cobra"

var example = &cobra.Command{
	Use:   "example",
	Short: "Start Example application",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	rootCmd.AddCommand(example)
}
