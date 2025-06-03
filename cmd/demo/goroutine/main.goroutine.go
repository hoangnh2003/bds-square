package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)
func main() {
	fmt.Println("Starting...")

	// 1. goroutine
	var wg sync.WaitGroup

	ids := []int{1, 2, 3, 4, 5} // productId

	start := time.Now()
	for _, id := range ids {
		wg.Add(1) // Mỗi loop add vào waitgroup
		go getProductByIdAPI(id, &wg) // detail product
	}

	// time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("Finished...", time.Since(start))
}

func getProductByIdAPI(id int, wg *sync.WaitGroup) {
	defer wg.Done() // hoàn thành 1 goroutine

	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf(">>> Data ProductId: %d: %s\n", id, string(body))
}


// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )
// func main() {
// 	fmt.Println("Starting...")

// 	ids := []int{1, 2, 3, 4, 5} // productId

// 	start := time.Now()
// 	for _, id := range ids {
// 		getProductByIdAPI(id) // detail product
// 	}

// 	fmt.Println("Finished...", time.Since(start))
// }

// func getProductByIdAPI(id int) {

// 	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Printf(">>> Data ProductId: %d: %s\n", id, string(body))
// }
