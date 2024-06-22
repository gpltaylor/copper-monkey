# README - Connecting to DynamoDB using Go and AWS SDK
The  following is a sample demonstarating how to connect to DynamoDB using Go and AWS SDK.

## REF

- https://dev.to/aws-heroes/dynamodb-local-in-docker-25i


## Setup
run the Db using Docker

```bash
docker run --rm -d --name dynamodb -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath /home/dynamodblocal
```

Setup alias to allow us to use this DB  whilst using AWS SDK

```bash
# alias to run `sqlite3` on this file
alias sql='docker exec -it dynamodb \
 sqlite3 /home/dynamodblocal/shared-local-instance.db'

# alias to run AWS CLI with linked to the DynamoDB entrypoint and exposing the current directory as /aws (which is the container home directory) 
alias aws='docker run --rm -it --link dynamodb:dynamodb -v $PWD:/aws \
 -e AWS_DEFAULT_REGION=xx -e AWS_ACCESS_KEY_ID=xx -e AWS_SECRET_ACCESS_KEY=xx \
 public.ecr.aws/aws-cli/aws-cli --endpoint-url http://dynamodb:8000'
```


Next we create the table and add some data

```
aws dynamodb create-table \
    --table-name PendingClientPayments \
    --attribute-definitions \
        AttributeName=BankAccount,AttributeType=S \
        AttributeName=Surname,AttributeType=S \
    --key-schema \
        AttributeName=BankAccount,KeyType=HASH \
        AttributeName=Surname,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --tags Key=Owner,Value=blueTeam

aws dynamodb put-item \
    --table-name PendingClientPayments \
    --item '{
        "BankAccount": {"S": "1234567890"},
        "FirstName": {"S": "John"},
        "Surname": {"S": "Doe"}
    }'
```

then query the data

```
aws dynamodb scan --table-name PendingClientPayments
```

then delete the table

```
aws dynamodb delete-table --table-name PendingClientPaymentsV1
```


## Go Code

```go
input := &dynamodb.ScanInput{
    TableName: aws.String("PendingClientPayments"),
}

result, err := svc.Scan(ctx, input)
if err != nil {
    return nil, err
}
```

## Casting Struct
The following shows how  to convert a scan result to a struct array

```go
type Client struct {
    BankAccount string
    FirstName    string
    Surname      string
}

var clients []Client

for _, i := range result.Items {
    client := Client{}
    err = dynamodbattribute.UnmarshalMap(i, &client)
    if err != nil {
        return nil, err
    }
    clients = append(clients, client)
}
```



