package main

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

func main() {
	var NAMESPACE_CONNECTION_STRING = os.Getenv("NAMESPACE_CONNECTION_STRING")
	if NAMESPACE_CONNECTION_STRING == "" {
		panic("NAMESPACE_CONNECTION_STRING is not set")
	}

	var EVENT_HUB_NAME = os.Getenv("EVENT_HUB_NAME")
	if EVENT_HUB_NAME == "" {
		panic("EVENT_HUB_NAME is not set")
	}

	// create an Event Hubs producer client using a connection string to the namespace and the event hub
	producerClient, err := azeventhubs.NewProducerClientFromConnectionString(NAMESPACE_CONNECTION_STRING, EVENT_HUB_NAME, nil)

	if err != nil {
		panic(err)
	}

	defer producerClient.Close(context.TODO())

	// create sample events
	events := createEventsForSample()

	// create a batch object and add sample events to the batch
	newBatchOptions := &azeventhubs.EventDataBatchOptions{}

	batch, err := producerClient.NewEventDataBatch(context.TODO(), newBatchOptions)
	if err != nil {
		panic("could not create batch object")
	}

	for i := 0; i < len(events); i++ {
		err = batch.AddEventData(events[i], nil)
		if err != nil {
			panic("could not add event data")
		}
	}

	// send the batch of events to the event hub
	producerClient.SendEventDataBatch(context.TODO(), batch, nil)
}

func createEventsForSample() []*azeventhubs.EventData {
	return []*azeventhubs.EventData{
		{
			Body: []byte("hello"),
		},
		{
			Body: []byte("world"),
		},
	}
}
