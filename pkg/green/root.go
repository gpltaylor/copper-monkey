package green

import (
	"fmt"
	"github.com/spf13/cobra"
)

type AddClientPaymentRequestData struct {
	EventId       string
	Action        string
	DateRequested string
	Amount        string
	CustomerId    string
}

var (
	data AddClientPaymentRequestData
)


func NewCmdSubAddClientPaymentRequest() *cobra.Command {
	data.Action = "AddClientPaymentRequest"

	cmd := &cobra.Command{
		Use:     "addclientpaymentrequest",
		Short:   "Request for a a client payment to be processed",
		Aliases: []string{"acpr"},
		//		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Adding client payment request:" + data.Amount)
		},
	}

	cmd.Flags().StringVarP(&data.CustomerId, "customerId", "s", "nil", "The customer id")
	cmd.Flags().StringVarP(&data.Amount, "amount", "e", "0", "End Index")
	return cmd
}

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func AddClientPaymentRequest(data AddClientPaymentRequestData) {
	fmt.Println("Adding client payment request:" + data.Amount)
}
