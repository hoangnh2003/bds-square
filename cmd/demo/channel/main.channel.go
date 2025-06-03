package main

import (
	"fmt"
	"time"
)

type Message struct {
	OrderId string
	Title string
	Price int
}

func publisher(channel chan<- Message, orders []Message) {
	for _, order := range orders {
		fmt.Printf("Pub:::%s\n", order.OrderId)
		channel <- order
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func subscriber(channel <-chan Message, userName string) {
	for msg := range channel {
		fmt.Printf("userName %s::Order::%s :: Title %s:: Price::%d\n",
			userName, msg.OrderId, msg.Title, msg.Price)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// 1 - channel order order
	orderChannel := make(chan Message)

	// 2 - Simulate orders
	orders := []Message{
		{OrderId: "Ordeer-01", Title: "Tips Go", Price: 30},
		{OrderId: "Ordeer-02", Title: "Tips Nodejs", Price: 30},
		{OrderId: "Ordeer-03", Title: "Tips Java", Price: 50},
	}

	// send order to pub
	go publisher(orderChannel, orders)
	go subscriber(orderChannel, "User")

	//sleep
	time.Sleep(time.Second * 2)
	fmt.Println("End pub sub...")
}