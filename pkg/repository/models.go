package repository


type ClientPaymentRequest struct {
	RequestId     string  `dynamodbav:"RequestId"`
	CustomerId    string  `dynamodbav:"CustomerId"`
	Action        string  `dynamodbav:"Action"`
	Status        string  `dynamodbav:"Status"`
	Amount        float32 `dynamodbav:"Amount"`
	FirstName     string  `dynamodbav:"FirstName"`
	Surname       string  `dynamodbav:"Surname"`
	Email         string  `dynamodbav:"Email"`
	DateRequested string  `dynamodbav:"DateRequested"`
}

// TODO: setup request id and date requested , status and action
func NewClientPaymentRequest() (ClientPaymentRequest, error) {
	return ClientPaymentRequest{

	} , nil
}
