package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func AddPendingPaymentRequest() (string, error) {
	ctx := context.TODO()
	cfg, err := GetConfig(ctx)

	svc := dynamodb.NewFromConfig(cfg)

	newPayment := PendingClientPayment{
		AccountNumber: "123456",
		FirstName:     "John",
		Surname:       "Connerxx",
	}

	err = InsertPendingClientPayment(svc, ctx, newPayment)
	if err != nil {
		fmt.Println("Error inserting payment")
		panic(err)
	}

	return "New Client Pending Payment Request make" , err
}

func InsertPendingClientPayment(svc *dynamodb.Client, ctx context.Context, payment PendingClientPayment) (error) {

	// What does this look like? what type of errors get thrown?
	item, err := attributevalue.MarshalMap(payment)
	if err != nil {
		return err
	}

	// Table name is a string vaule. This needs to be controlled using Config
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("PendingClientPayments"),
	}

	_, err = svc.PutItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

