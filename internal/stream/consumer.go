package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"

	m "monitoring/internal/models"
)

func Consumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"redpanda:9092"}, // внутренняя docker-сеть
		Topic:   "DSPbyHour",
		GroupID: "logic-service",
	})
	defer r.Close()

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil { log.Fatal(err) }

		var	row m.SingleDSP
		if err := json.Unmarshal(message.Value, &row); err != nil {
			log.Printf("bad json: %v", err); continue
		}

		// --- ваша бизнес-логика ---
		fmt.Printf("DSP %s, %02d-й час: %.2f €\n",
			row.DSPname, row.Hour, row.Amount)
		// TODO: вызвать API склада, отправить письмо, обновить кэш …
	}
}