package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent represents the structure of the incoming JSON event
type MyEvent struct {
	Name string `json:"What is your name?"` // Name field represents the name of the person
	Age  int    `json:"How old are you?"`   // Age field represents the age of the person
}

// MyResponse represents the structure of the outgoing JSON response
type MyResponse struct {
	Message string `json:"Answer:"` // Message field holds the response message
}

// HandleLambdaEvent is the function executed when the Lambda function is invoked
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	// Constructing the response message with the received name and age
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	// Start the Lambda execution with the HandleLambdaEvent function
	lambda.Start(HandleLambdaEvent)
}
