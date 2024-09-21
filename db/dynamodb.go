package db

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Message represents the structure of a chat message
type Message struct {
	StartupIdea string `json:"startup-idea"` // Partition key
	Content     string `json:"content"`      // Message content
	Timestamp   string `json:"timestamp"`    // Message timestamp
}

// DynamoDB session and service client
var sess = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("us-east-1"), // Replace with your AWS region
}))

var svc = dynamodb.New(sess)

// SaveMessageToDynamoDB saves the received message to the DynamoDB table
func SaveMessageToDynamoDB(content string, startupIdea string) error {
	// Create a new message with the current timestamp
	msg := Message{
		StartupIdea: startupIdea,                     // Partition key
		Content:     content,                         // Chat message content
		Timestamp:   time.Now().Format(time.RFC3339), // Timestamp
	}

	// Marshal the message to a DynamoDB-compatible format
	av, err := dynamodbattribute.MarshalMap(msg)
	if err != nil {
		log.Printf("Error marshalling message: %v", err) // Log error
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// Define the PutItem input
	input := &dynamodb.PutItemInput{
		TableName: aws.String("scr-chat-logs"), // Your table name
		Item:      av,                          // Marshaled message
	}

	// Save the item (message) to DynamoDB
	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Error saving message to DynamoDB: %v", err) // Log error
		return fmt.Errorf("failed to save message to DynamoDB: %v", err)
	}

	log.Printf("Message saved successfully: %v", msg) // Log success
	return nil
}
