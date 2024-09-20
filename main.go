package main

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/NeowayLabs/wabbit"
	"github.com/cheshir/go-mq"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/hallgren/eventsourcing"
)

func main() {
	fmt.Println("Using dependencies to satisfy go mod tidy.")

	// Using github.com/IBM/sarama
	var producer sarama.SyncProducer
	_ = producer

	// Using github.com/NeowayLabs/wabbit
	var conn wabbit.Conn
	_ = conn

	// Using github.com/cheshir/go-mq
	var mqConfig mq.Config
	_ = mqConfig

	// Using github.com/cosmos/cosmos-sdk
	var cliCtx client.Context
	_ = cliCtx

	// Using github.com/hallgren/eventsourcing
	var eventStore eventsourcing.EventRepository
	_ = eventStore
}
