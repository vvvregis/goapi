package main

import (
	"awesomeProject/repository"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	"log"
)

// connect открываем один раз
var connect clickhouse.Conn

// нужно ли заполнять тестовыми данными в данный прогон
const useTestData bool = false

// инициализация подключения к БД
func init() {
	connect = repository.Connect()
}

// главная горутина
// Еще я бы изменил тип EventID на UUID
func main() {
	if useTestData {
		err := addTestData(connect)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("test data add successfully")
	}

	router := gin.Default()
	router.POST("/add", add)
	err := router.Run("localhost:4040")
	if err != nil {
		log.Fatal(err)
	}
}

// функция обрабатывает роут добавления евента
func add(context *gin.Context) {
	var event repository.Event
	err := context.BindJSON(&event)
	if err != nil {
		log.Fatal(err)
	}

	context.IndentedJSON(200, event)

	if err != nil {
		log.Fatal(err)
	}

	err = event.AddEvent(connect)

	if err != nil {
		log.Fatal(err)
	}
}
