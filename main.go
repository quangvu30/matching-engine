package main

import (
	"fmt"

	"github.com/quangvu30/matching-engine/core"
	"github.com/quangvu30/matching-engine/types"
)

func main() {
	appl := core.NewOrderBook()
	var ms []types.PipeMsg
	ms = appl.AddOrder(types.Order{ID: 166666666, Price: 100, Qty: 5, Side: 1, Type: 0})
	fmt.Println(ms)
	ms = appl.AddOrder(types.Order{ID: 166666667, Price: 101, Qty: 5, Side: 1, Type: 0})
	fmt.Println(ms)
	ms = appl.AddOrder(types.Order{ID: 166666668, Price: 102, Qty: 10, Side: 1, Type: 0})
	fmt.Println(ms)
	ms = appl.AddOrder(types.Order{ID: 166666669, Price: 103, Qty: 10, Side: 0, Type: 0})
	fmt.Println(ms)
	ms = appl.AddOrder(types.Order{ID: 166666670, Price: 104, Qty: 10, Side: 0, Type: 0})
	fmt.Println(ms)

	fmt.Println("-------- Order Book --------")
	appl.GetOrderBook()
}
