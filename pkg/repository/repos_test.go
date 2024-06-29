package repository

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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


	newPayment := ClientPaymentRequest{
		RequestId:     "123",
		CustomerId:    "123",
		Action:        "AddClientPaymentRequest",
		Status:        "Pending",
		Amount:        100.00,
		FirstName:     "John",
		Surname:       "Doe",
		Email:         "john@gmail.com",
	}

	//	clientPaymentRequest := repository.NewClientPaymentRequest(data)
	// TODO : Using NewClientPaymentRequest

  msg, err := AddPendingPaymentRequest(newPayment)

	if err != nil {
		t.Errorf("Error inserting payment: " + err.Error())
	} else {
		t.Logf("Payment inserted: " + msg)
	}

	result, err := GetPaymentRequestByRequestId("123")
	if err != nil {
		t.Errorf("Error getting payment request: " + err.Error())
	} else {
		// TODO : check the values are the same
		t.Logf("Payment request: " + result.FirstName)
	}

}
