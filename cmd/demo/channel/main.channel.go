package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	OrderId string
	Title   string
	Price   int
}

func publisher(channel chan<- Message, orders []Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, order := range orders {
		fmt.Printf("Pub:::%s\n", order.OrderId)
		channel <- order
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func subscriber(channel <-chan Message, userName string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range channel {
		fmt.Printf("userName %s::Order::%s :: Title %s:: Price::%d\n",
			userName, msg.OrderId, msg.Title, msg.Price)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	orderChannel := make(chan Message)

	orders := []Message{
		{OrderId: "Order-01", Title: "Tips Go", Price: 30},
		{OrderId: "Order-02", Title: "Tips Nodejs", Price: 30},
		{OrderId: "Order-03", Title: "Tips Java", Price: 50},
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go publisher(orderChannel, orders, &wg)
	go subscriber(orderChannel, "User", &wg)

	wg.Wait()
	fmt.Println("End pub sub...")
}
