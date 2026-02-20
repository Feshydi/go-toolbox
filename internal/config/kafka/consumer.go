package kafka_config

import (
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func MustCreateConsumer(brokers []string, groupId, topic string, sessionTimeout time.Duration, autoOffsetReset string, commitInterval time.Duration) *kafka.Consumer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(brokers, ","),
		"group.id":                 groupId,
		"session.timeout.ms":       int(sessionTimeout.Milliseconds()),
		"auto.offset.reset":        autoOffsetReset,
		"enable.auto.offset.store": false,
		"auto.commit.interval.ms":  int(commitInterval.Milliseconds()),
	}

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		panic(fmt.Errorf("error creating kafka.Consumer: %w", err))
	}

	if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
		panic(fmt.Errorf("error subscribing to topic %s: %w", topic, err))
	}

	return c
}
