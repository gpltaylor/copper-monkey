
package green

import (
	"fmt"
	"github.com/gpltaylor/copper-monkey/pkg/repository"
	"github.com/spf13/cobra"
)


func NewCmdSubDeleteClientPaymentRequest() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "deleteclientpaymentrequest",
		Short:   "Delete a payment request using the RequestId",
		Aliases: []string{"dcpr"},
		//		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Deleting client payment request:")
			requestId := cmd.Flag("RequestID").Value.String()

			err := repository.DeletePendingPaymentRequest(requestId)

			if err != nil {
				fmt.Println("Error deleting payment request")
				panic(err)
			}
			fmt.Println("Payment request sucessfully deleted")
		},
	}

	cmd.Flags().StringVarP(&requestId, "RequestID", "s", "nil", "The request id")
	return cmd
}

