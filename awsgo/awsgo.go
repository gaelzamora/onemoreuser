package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Ctf aws.Config
var err error

func StartAWS() {
	Ctx = context.TODO()
	Ctf, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error to attempt configure "+err.Error())
	}
}
