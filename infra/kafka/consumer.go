package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

// Simula um construtor
func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

// Criando um consumidor no Kafka
func (c Consumer) Consume(msgChannel chan *ckafka.Message) error {

	// Criando consumidor no Kafka e validando-o
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}

	// Laço para ler o que tem no Kafka
	for {
		msgKafka, err := consumer.ReadMessage(-1)
		if err == nil {
			// Atribuindo as mensagens recebidas no Kafka para no nosso Canal
			msgChannel <- msgKafka
		}
	}
}
