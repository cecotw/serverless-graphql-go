package schema

import (
	"github.com/cecotw/serverless-graphql-go/internal/app/graphql/todo"
)

// QueryResolver : Query Resolver
type QueryResolver struct {
	*todo.Resolver
}

// Schema Schema
var Schema = `
schema {
  query: Query
}
type Query {}
type Mutation {}
extend type Query {
  todo(id: ID!): Todo
  todos: [Todo]
}
extend type Mutation {
  createTodo(input: TodoInput!): Todo!
}
type Todo {
 id: ID!
 message: String
 isComplete: Boolean
}
input TodoInput {
  id: ID
  message: String
  isComplete: Boolean!
}
`
