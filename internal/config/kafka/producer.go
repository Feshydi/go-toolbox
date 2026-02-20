package kafka_config

import (
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func MustCreateProducer(brokers []string, groupId string, requiredAcks, batchSize int) *kafka.Producer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":     strings.Join(brokers, ","),
		"group.id":              groupId,
		"request.required.acks": requiredAcks, // RequireNone = 0; RequireOne = 1; RequireAll = -1
		"batch.size":            batchSize,
	}

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		panic(fmt.Errorf("error creating kafka.Producer: %w", err))
	}

	return p
}
