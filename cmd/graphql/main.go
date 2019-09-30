package graphql

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cecotw/serverless-graphql-go/internal/app/graphql/schema"
	"github.com/cecotw/serverless-graphql-go/internal/pkg/handler"
)

var graphql = new(handler.GraphQl)

func init() {
	graphql.BuildSchema(schema.Schema, &schema.QueryResolver{})
}

func main() {
	lambda.Start(graphql.Lambda)
}
