package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/quangvu30/matching-engine/core"
	"github.com/quangvu30/matching-engine/types"
)

func main() {
	appl := core.NewOrderBook()
	fmt.Println("-------- Random Orders --------")
	// var ms []types.PipeMsg
	start := time.Now()
	for i := 0; i < 100; i++ {
		order := randomOrder(i)
		appl.AddLimitOrder(order)
	}
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	ask := appl.GetAskDepth(5)
	bid := appl.GetBidDepth(5)

	fmt.Println("-------- Ask Depth --------")
	for _, a := range ask {
		fmt.Println(a)
	}

	fmt.Println("-------- Bid Depth --------")
	for _, b := range bid {
		fmt.Println(b)
	}

}

func randomOrder(i int) types.Order {
	// Generate a random price between 1 and 100
	price := rand.Intn(10) + 1

	// Generate a random quantity between 1 and 10
	qty := rand.Intn(100) + 1

	// Generate a random side (0 for buy, 1 for sell)
	side := rand.Intn(2)
	if side == 0 {
		return types.Order{
			ID:    uint64(100000000 + i),
			Price: float64(price),
			Qty:   float64(qty),
			Side:  0,
		}
	} else {
		return types.Order{
			ID:    uint64(100000000 + i),
			Price: float64(price),
			Qty:   float64(qty),
			Side:  1,
		}
	}
}
