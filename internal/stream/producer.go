package stream

import (
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"

	m "monitoring/internal/models"
)

func produce(stats m.DSPbyHour) error {
    w := &kafka.Writer{
        Addr:     kafka.TCP("localhost:19092"), // внешний порт из compose
        Topic:    "DSPbyHour",
        Balancer: &kafka.LeastBytes{},
    }
    defer w.Close()

    ctx := context.Background()
    for _, row := range stats.Data {
        b, _ := json.Marshal(row) //json
        if err := w.WriteMessages(ctx, kafka.Message{Value: b}); err != nil {
            return err
        }
    }
    return nil
}