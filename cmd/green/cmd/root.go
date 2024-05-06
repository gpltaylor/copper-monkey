package green

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/gpltaylor/copper-monkey/pkg/green"
)


var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
		amount := args[0]

		data := green.AddClientPaymentRequestData{
			Amount: amount,
		}

		green.AddClientPaymentRequest(data)
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

// copper-monkey-green AddClientPaymentRequest --amount 12.99 --customerId {GUID}

