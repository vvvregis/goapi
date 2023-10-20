package main

import (
	"awesomeProject/repository"
	"github.com/ClickHouse/clickhouse-go/v2"
	randomDataTime "github.com/duktig-solutions/go-random-date-generator"
	"math/rand"
	"strings"
)

const testDataCount = 20

// Добавляет тестовые данные
func addTestData(connect clickhouse.Conn) error {
	for i := 0; i < testDataCount; i++ {

		randomDate, err := getRandomDate()
		if err != nil {
			return err
		}
		event := repository.Event{
			EventID:   getRandomInt(),
			EventType: getRandomString(),
			UserID:    getRandomInt(),
			EventTime: randomDate,
			Payload:   getRandomString(),
		}

		err = event.AddEvent(connect)

		if err != nil {
			return err
		}
	}

	return nil

}

// Генерирует случайную строку
func getRandomString() string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

// генерирует случайное число
func getRandomInt() int64 {
	return rand.Int63()
}

// генерирует случайную дату
func getRandomDate() (string, error) {
	return randomDataTime.GenerateDateTime(
		"1970-01-01 00:00:00",
		"2023-08-21 17:08:26",
	)
}
