package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
)

type Event struct {
	EventID   int64
	EventType string
	UserID    int64
	EventTime string
	Payload   string
}

// AddEvent Добавляет евент
func (e Event) AddEvent(connect clickhouse.Conn) error {
	ctx := context.Background()
	err := connect.AsyncInsert(
		ctx,
		"INSERT INTO events(eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)",
		false,
		e.EventID,
		e.EventType,
		e.UserID,
		e.EventTime,
		e.Payload,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetEvents Получает евенты в заданном интервале
func GetEvents(connect clickhouse.Conn, startData string, endData string) error {
	ctx := context.Background()
	var events []Event
	err := connect.Select(ctx, &events, "SELECT eventID as EventID, eventType as EventType, userID as UserID, "+
		"toString(eventTime) as EventTime, payload as Payload FROM events WHERE eventTime > ? AND EventTime < ?",
		startData, endData)

	fmt.Println(err)
	if err != nil {
		return err
	}

	fmt.Println(events)

	return nil
}
