package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("my-urmi-exchange", "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	masq, err := ch.QueueDeclare("rex-queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(masq.Name, "myRoUte.*", "my-urmi-exchange", false, nil)
	if err != nil {
		log.Fatal(err)
	}
	masCon, err := ch.Consume(masq.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	for m := range masCon {
		time.Sleep(2 * time.Second)
		fmt.Println(m.Body)
	}
}
