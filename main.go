package main

import (
	"github.com/quangvu30/matching-engine/core"
	"github.com/quangvu30/matching-engine/types"
)

func main() {
	pipe1 := make(chan []types.PipeMsg)
	appl := core.NewOrderBook(pipe1)
	appl.AddOrder(types.Order{ID: 166666666, Price: 100, Qty: 10, Side: 0, Type: 0})
	appl.AddOrder(types.Order{ID: 166666667, Price: 101, Qty: 10, Side: 1, Type: 0})
	appl.AddOrder(types.Order{ID: 166666668, Price: 102, Qty: 10, Side: 0, Type: 0})
	appl.AddOrder(types.Order{ID: 166666669, Price: 100, Qty: 10, Side: 1, Type: 0})
	appl.AddOrder(types.Order{ID: 166666670, Price: 104, Qty: 10, Side: 0, Type: 0})
	appl.GetOrderBook()
}
