package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	tableName = "PendingClientPaymentsV1"
)

func AddPendingPaymentRequest(paymentRequest ClientPaymentRequest) (string, error) {
	InitDatabase()

	ctx := context.TODO()
	cfg, err := GetConfig(ctx)

	svc := dynamodb.NewFromConfig(cfg)

	newPayment := ClientPaymentRequest{
		FirstName:     paymentRequest.FirstName,
		Surname:       paymentRequest.Surname,
		Email:         paymentRequest.Email,
		Amount:        paymentRequest.Amount,
		CustomerId:    paymentRequest.CustomerId,
		Action:        paymentRequest.Action,
		Status:        paymentRequest.Status,
		RequestId:     paymentRequest.RequestId,
		DateRequested: time.Now().Format(time.RFC3339), // move this to config
	}

	err = InsertPendingClientPayment(svc, ctx, newPayment)
	if err != nil {
		fmt.Println("Error inserting payment")
		panic(err)
	}

	return "New Client Pending Payment Request make", err
}

func InsertPendingClientPayment(svc *dynamodb.Client, ctx context.Context, payment ClientPaymentRequest) error {

	// What does this look like? what type of errors get thrown?
	item, err := attributevalue.MarshalMap(payment)
	if err != nil {
		return err
	}

	// Table name is a string vaule. This needs to be controlled using Config
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

/*
*
  - Create the table in DynamoDB
    RequestId     string  `dynamodbav:"RequestId"`
    CustomerId    string  `dynamodbav:"CustomerId"`
    Action        string  `dynamodbav:"Action"`
    Status        string  `dynamodbav:"Status"`
    Amount        float32 `dynamodbav:"Amount"`
    FirstName     string  `dynamodbav:"FirstName"`
    Surname       string  `dynamodbav:"Surname"`
    Email         string  `dynamodbav:"Email"`
    DateRequested string  `dynamodbav:"DateRequested"`
*/
func InitDatabase() {
	ctx := context.TODO()
	cfg, err := GetConfig(ctx)

	svc := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("RequestId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("RequestId"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	}

	// Check if the table exists
	exists, err := CheckTableExists(svc, &ctx, tableName)
	if err != nil {
		fmt.Println("Error checking if table exists")
		panic(err)
	}

	if !exists {
		fmt.Println("Table does not exist, creating table")
		_, err = svc.CreateTable(ctx, input)
		if err != nil {
			fmt.Println("Got error calling CreateTable:")
			fmt.Println(err)
			return
		}
	}
}

func CheckTableExists(svc *dynamodb.Client, ctx *context.Context, tableName string) (bool, error) {
	_, err := svc.DescribeTable(*ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err != nil {

		// Check if the error is a ResourceNotFoundException
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			return false, nil // Table does not CheckTableExists(
		}
		return false, err // Other error occurred
	}
	return true, nil // Table CheckTableExists()
}
