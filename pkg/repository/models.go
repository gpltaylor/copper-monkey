package repository

import (
	"github.com/google/uuid"
)

/**
 * DTO to store non-trusted data inputs, used as the input for the API/CLI command
 * This DTO is then ETL'd into a trusted data object
 */
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

// Write some test around this logic
func NewClientPaymentRequest(data AddClientPaymentRequestData) ClientPaymentRequest {
	return ClientPaymentRequest{
		RequestId:  uuid.New().String(),
		Status:     "Pending",
		Action:     "AddClientPaymentRequest",
		CustomerId: data.CustomerId,
		Amount:     data.Amount,
		FirstName:  data.FirstName,
		Surname:    data.Surname,
		Email:      data.Email,
	}
}
