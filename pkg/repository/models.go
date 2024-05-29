package repository

type PendingClientPayment struct {
	FirstName     string `dynamodbav:"FirstName"`
	Surname       string `dynamodbav:"Surname"`
	AccountNumber string `dynamodbav:"BankAccount"`
}

