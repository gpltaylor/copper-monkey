package green

import (
	"fmt"
	"github.com/gpltaylor/copper-monkey/pkg/green"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			amount, _ := cmd.Flags().GetString("amount")

			data := green.AddClientPaymentRequestData{
				Amount: amount,
			}

			green.AddClientPaymentRequest(data)
			// Get values as flags

			// Change aboeve to that there are err and a response with a request id
			// Reply with the request id.
			// Review sub command
		},
	}
}

var rootCmd = NewRootCmd()

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Init NewRootCmd")
	rootCmd.Flags().StringP("amount", "a", "", "Amount to add")
}

// copper-monkey-green AddClientPaymentRequest --amount 12.99 --customerId {GUID}
