package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "agarwalnitin-test")
	if err != nil {
		fmt.Println("Got error", err)
	}

	fmt.Println("Got client", client)
	topicName := "testTopic"
	topic := client.Topic(topicName)

	exists, err := topic.Exists(ctx)
	if err != nil {
		fmt.Println("Got error", err)
	}

	fmt.Println("Got the topic", exists)

	if !exists {
		log.Printf("Topic %v doesn't exist - creating it", topicName)
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			fmt.Println("Got error", err)
		}
	}

	// See the other examples to learn how to use the Client.
}
