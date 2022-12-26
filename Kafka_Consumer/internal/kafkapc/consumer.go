package kafkapc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type DebObject struct {
	Schema struct {
		Type   string `json:"type,omitempty"`
		Fields []struct {
			Type   string `json:"type,omitempty"`
			Fields []struct {
				Type     string `json:"type,omitempty"`
				Optional bool   `json:"optional,omitempty"`
				Default  int    `json:"default,omitempty"`
				Field    string `json:"field,omitempty"`
			} `json:"fields,omitempty"`
			Optional bool   `json:"optional,omitempty"`
			Name     string `json:"name,omitempty"`
			Field    string `json:"field,omitempty"`
		} `json:"fields,omitempty"`
		Optional bool   `json:"optional,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"schema,omitempty"`
	Payload struct {
		Before interface{} `json:"before,omitempty"`
		After  struct {
			ID        int    `json:"id,omitempty"`
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
			Email     string `json:"email,omitempty"`
		} `json:"after,omitempty"`
		Source struct {
			Version   string      `json:"version,omitempty"`
			Connector string      `json:"connector,omitempty"`
			Name      string      `json:"name,omitempty"`
			TsMs      int64       `json:"ts_ms,omitempty"`
			Snapshot  string      `json:"snapshot,omitempty"`
			Db        string      `json:"db,omitempty"`
			Sequence  string      `json:"sequence,omitempty"`
			Schema    string      `json:"schema,omitempty"`
			Table     string      `json:"table,omitempty"`
			TxID      int         `json:"txId,omitempty"`
			Lsn       int         `json:"lsn,omitempty"`
			Xmin      interface{} `json:"xmin,omitempty"`
		} `json:"source,omitempty"`
		Op          string      `json:"op,omitempty"`
		TsMs        int64       `json:"ts_ms,omitempty"`
		Transaction interface{} `json:"transaction,omitempty"`
	} `json:"payload,omitempty"`
}

func StartConsumer(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		debObj := DebObject{}
		if err := json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&debObj); err != nil {
			fmt.Println("Debezium decode error")
		}
		fmt.Println("output", debObj.Payload.After)
		// fmt.Println("output", debObj)

		// fmt.Println("received: ", string(msg.Value))

	}
}
