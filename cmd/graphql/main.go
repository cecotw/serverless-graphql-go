package graphql

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cecotw/serverless-graphql-go/internal/pkg/handler"
)

var graphql = new(handler.GraphQl)

func main() {
	lambda.Start(graphql.Lambda)
}
