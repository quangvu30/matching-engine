package core

import (
	"github.com/gammazero/deque"
	"github.com/quangvu30/matching-engine/types"
)

const (
	LIMIT      int8 = 1
	MARKET     int8 = 0
	STOP       int8 = 2
	STOP_LIMIT int8 = 3
)

type MarketValue struct {
	Value float64
	Qty   float64
	Bid   float64
	Ask   float64
}

type OrderManager struct {
	Products map[string]*OrderBook
	Queues   map[string]*deque.Deque[types.RawOrder]
	Req      map[string]chan types.Order
	Res      map[string]chan types.ResultMatching
	market   map[string]MarketValue
}

func NewOrderManager() *OrderManager {
	return &OrderManager{
		Products: make(map[string]*OrderBook),
		Queues:   make(map[string]*deque.Deque[types.RawOrder]),
		Req:      make(map[string]chan types.Order),
		Res:      make(map[string]chan types.ResultMatching),
		market:   make(map[string]MarketValue),
	}
}

/*
* Factory creates a new OrderBook for each product
* and initializes the Queues and Pipe for each product
 */
func (om *OrderManager) Factory(orders []types.RawOrder) {
	for _, order := range orders {
		if _, ok := om.Products[order.Code]; !ok {
			om.Products[order.Code] = NewOrderBook()
			om.Queues[order.Code] = deque.New[types.RawOrder]()
			om.Req[order.Code] = make(chan types.Order, 100)
			om.Res[order.Code] = make(chan types.ResultMatching)
			om.market[order.Code] = MarketValue{Value: 0, Qty: 0}
		}

		switch order.Type {
		case MARKET:
			om.market[order.Code] = MarketValue{Value: om.market[order.Code].Value + order.Price, Qty: om.market[order.Code].Qty + order.Qty}
		case LIMIT:
		}
		om.Queues[order.Code].PushBack(order)
	}
}
