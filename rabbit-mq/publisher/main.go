package main

import (
	"log"

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
	massages := []struct {
		routhingKey string
		body        string
	}{
		{"myRoUte.cat", "cat sayes meow"},
		{"myRoute", "i dog name lusi"},
		{"myRoUte.dog", "dog sayes boh"},
		{"myRoUte.cat", "i hate cat"},
	}

	for _, m := range massages {
		mes := amqp.Publishing{
			ContentType: "text/plan",
			Body:        []byte(m.body),
		}
		err := ch.Publish("my-urmi-exchange", m.routhingKey, false, false, mes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
