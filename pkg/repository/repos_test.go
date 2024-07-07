package repository

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

func TestSetupDatabase(t *testing.T) {
	// create a random table name
	tableName := "PendingClientPayments-" + strconv.Itoa(rand.Intn(1000))

	ctx := context.TODO()
	cfg, _ := GetConfig(ctx)

	svc := dynamodb.NewFromConfig(cfg)
	SetTableName(tableName)
	InitDatabase()

	if exist, _ := CheckTableExists(svc, &ctx, tableName); exist != true {
		t.Errorf("Table does not exist : " + tableName)
	} else {
		t.Logf("Table exists : " + tableName)
	}
}

// Insert into the database
func TestInsertPendingClientPayment(t *testing.T) {

	tableName := "PendingClientPayments-" + strconv.Itoa(rand.Intn(1000))
	t.Logf("Table name: " + tableName)
	SetTableName(tableName)

	data := AddClientPaymentRequestData{
		CustomerId: "123",
		FirstName:  "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Amount:     19.99,
	}

	clientPaymentRequest := NewClientPaymentRequest(data)

	ClientPaymentRequestentPaymentRequest, err := AddPendingPaymentRequest(clientPaymentRequest)
	if err != nil {
		t.Errorf("Error inserting payment: " + err.Error())
	} else {
		t.Logf("Payment inserted: " + ClientPaymentRequestentPaymentRequest.RequestId)
	}

	result, err := GetPaymentRequestByRequestId(clientPaymentRequest.RequestId)
	if err != nil {
		t.Errorf("Error getting payment request: " + err.Error())
	} else {
		// TODO : check the values are the same
		t.Logf("Payment request: " + result.FirstName)
	}

	// check that the request id has a valid GUID
	if _, err := uuid.Parse(result.RequestId); err != nil {
		t.Errorf("Payment request id is not a valid GUID: " + result.RequestId)
	}

	// Check that the new request has a pending status
	if result.Status != "Pending" {
		t.Errorf("Payment request status is not pending: " + result.Status)
	}

	// Check that the Action is set to AddClientPaymentRequest
	if result.Action != "AddClientPaymentRequest" {
		t.Errorf("Payment request action is not AddClientPaymentRequest: " + result.Action)
	}

	// TODO: Write a test to see what happens when you get a payment request that does not exist

}
