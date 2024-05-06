package green

import (
	"fmt"
)

type AddClientPaymentRequestData struct {
	EventId       string
	Action        string
	DateRequested string
	Amount        string
}

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func AddClientPaymentRequest(data AddClientPaymentRequestData) {
	fmt.Println("Adding client payment request:" + data.Amount)
}
