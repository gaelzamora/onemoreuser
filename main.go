package main

import (
	"context"
	"errors"
	"fmt"
	"onemoreuser/awsgo"
	"onemoreuser/database"
	"onemoreuser/models"
	"os"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaFunction)
}

func LambdaFunction(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.StartAWS()

	if !ValidParams() {
		fmt.Println("Error! you have to send 'SecretManager'")
		err := errors.New("error in parameters SecretManager")

		return event, err
	}
	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email: "+data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub: "+data.UserUUID)
		}
	}

	err := database.ReadSecret()

	if err != nil {
		fmt.Println(err.Error())
		return event, err
	}

	err = database.SignUp(data)

	return event, err
}

func ValidParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")

	return getParam
}