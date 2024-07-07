package green

import (
	"fmt"
	"github.com/gpltaylor/copper-monkey/pkg/repository"
	"github.com/spf13/cobra"
)

var (
	requestId string
)

func NewCmdSubGetClientPaymentRequest() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "getclientpaymentrequest",
		Short:   "Get a payment request using the RequestId",
		Aliases: []string{"gcpr"},
		//		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Getting client payment request:")
			requestId := cmd.Flag("RequestID").Value.String()

			clientPaymentRequest, err := repository.GetPaymentRequestByRequestId(requestId)

			if err != nil {
				fmt.Println("Error getting payment request")
				panic(err)
			}
			fmt.Println(clientPaymentRequest)
		},
	}

	cmd.Flags().StringVarP(&requestId, "RequestID", "s", "nil", "The request id")
	return cmd
}

