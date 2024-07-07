package green

import (
	"fmt"
	"github.com/gpltaylor/copper-monkey/pkg/repository"
	"github.com/spf13/cobra"
)

var (
	data repository.AddClientPaymentRequestData
)

func NewCmdSubAddClientPaymentRequest() *cobra.Command {
	data.Action = "AddClientPaymentRequest"

	cmd := &cobra.Command{
		Use:     "addclientpaymentrequest",
		Short:   "Request for a a client payment to be processed",
		Aliases: []string{"acpr"},
		//		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Adding client payment request:")
			fmt.Print(data)

			clientPaymentRequest := repository.NewClientPaymentRequest(data)

			// Save to database
			clientPaymentRequest, err := repository.AddPendingPaymentRequest(clientPaymentRequest)
			if err != nil {
				fmt.Println("Error inserting payment")
				panic(err)
			}
			fmt.Println(clientPaymentRequest)
		},
	}

	cmd.Flags().StringVarP(&data.CustomerId, "CustomerID", "s", "nil", "The customer id")
	cmd.Flags().Float32VarP(&data.Amount, "Amount", "e", 0, "End Index")
	cmd.Flags().StringVar(&data.FirstName, "FirstName", "nil", "Customers first name")
	cmd.Flags().StringVar(&data.Surname, "Surname", "nil", "Customers surname")
	cmd.Flags().StringVar(&data.Email, "Email", "nil", "Customers email")
	return cmd
}

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func AddClientPaymentRequest(data repository.AddClientPaymentRequestData) {
	fmt.Printf("Adding client payment request: %f", data.Amount)
}
