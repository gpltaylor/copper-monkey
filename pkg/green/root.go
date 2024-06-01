package green

import (
	"fmt"
	"github.com/gpltaylor/copper-monkey/pkg/repository"
	"github.com/spf13/cobra"
	"github.com/google/uuid"
)

type AddClientPaymentRequestData struct {
	RequestId     string
	CustomerId    string
	Action        string
	Status        string // Pending, Cancelled, Processed
	Amount        float32
	FirstName     string
	Surname       string
	Email         string
	DateRequested string
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
			fmt.Printf("Adding client payment request:")
			fmt.Print(data)

			clientPaymentRequest := repository.ClientPaymentRequest{
				RequestId:     uuid.New().String(),
				CustomerId:    data.CustomerId,
				Amount: 			data.Amount,
				FirstName:     data.FirstName,
				Surname:       data.Surname,
				Email:         data.Email,
				Status:        "Pending",
				Action:        "AddClientPaymentRequest",
			}
			// Save to database
			msg, err := repository.AddPendingPaymentRequest(clientPaymentRequest)
			if err != nil {
				fmt.Println("Error inserting payment")
				panic(err)
			}
			fmt.Println(msg)
		},
	}

	cmd.Flags().StringVarP(&data.CustomerId, "customerId", "s", "nil", "The customer id")
	cmd.Flags().Float32VarP(&data.Amount, "Amount", "e", 0, "End Index")
	cmd.Flags().StringVar(&data.FirstName, "FirstName", "nil", "Customers first name")
	cmd.Flags().StringVar(&data.Surname, "Surname", "nil", "Customers surname")
	cmd.Flags().StringVar(&data.Email, "Email", "nil", "Customers email")
	return cmd
}

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func AddClientPaymentRequest(data AddClientPaymentRequestData) {
	fmt.Printf("Adding client payment request: %f", data.Amount)
}
