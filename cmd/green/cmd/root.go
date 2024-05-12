package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"github.com/gpltaylor/copper-monkey/pkg/substr"
	"github.com/gpltaylor/copper-monkey/pkg/green"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			// REturn hello world
			fmt.Println("Hello World")
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
	rootCmd.AddCommand(substr.NewCmdSubstr())
	rootCmd.AddCommand(green.NewCmdSubAddClientPaymentRequest())

}

// copper-monkey-green AddClientPaymentRequest --amount 12.99 --customerId {GUID}
