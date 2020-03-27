package workers

import (
	"encoding/json"
	"log"

	"github.com/Mekawy5/chatapp/domain/application"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

func Consume(db *gorm.DB) {
	q := tools.NewRabbitClient()
	q.Setup()
	defer q.Conn.Close()

	apps, err := q.Channel.Consume(
		tools.APPQ,          //queueName
		"chat-app-consumer", //consumerTag
		true,                // noAck
		false,               // exclusive
		false,               // noLocal
		false,               // noWait
		nil,                 // arguments
	)
	handleConsumerErr(err, "applications")
	consumeApps(apps, db)

}

func consumeApps(apps <-chan amqp.Delivery, db *gorm.DB) {
	for d := range apps {
		// parse application info
		var app application.Application
		json.Unmarshal(d.Body, &app)

		appRepo := application.NewApplicationRepository(db)
		appService := application.NewApplicationService(appRepo)

		appService.Save(application.NewApplication(app))
		//TODO to log received application
	}
}

func handleConsumerErr(err error, consumer string) {
	if err != nil {
		log.Fatal("Error consuming " + consumer)
		log.Fatal(err)
		// os.Exit(0)
	}
}
