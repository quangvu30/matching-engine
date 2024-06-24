package core

import (
	"fmt"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
	"github.com/gammazero/deque"
	"github.com/quangvu30/matching-engine/types"

	mutils "github.com/quangvu30/matching-engine/utils"
)

const (
	B int8 = 0
	S int8 = 1
)

type OrderBook struct {
	Buy  map[float64]*treemap.Map
	BP   *deque.Deque[float64]
	Sell map[float64]*treemap.Map
	SP   *deque.Deque[float64]
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		SP:   deque.New[float64](),
		BP:   deque.New[float64](),
		Buy:  make(map[float64]*treemap.Map),
		Sell: make(map[float64]*treemap.Map),
	}
}

func (ob *OrderBook) AddOrder(order types.Order) []types.PipeMsg {
	msgs := []types.PipeMsg{}
	if order.Side == B {
		// kiem tra gia mua da ton tai chua
		tree, exist := ob.Buy[order.Price]
		if exist {
			// them gia mua vao tree
			tree.Put(order.ID, order.Qty)
		} else {
			// them gia mua moi
			ob.Buy[order.Price] = treemap.NewWith(utils.UInt64Comparator)
			ob.Buy[order.Price].Put(order.ID, order.Qty)
			mutils.InsertAsc(ob.BP, order.Price)

			// tim gia ban nho nhat
			if ob.SP.Len() == 0 {
				return nil
			}
			pSellMin := ob.SP.Front()
			for order.Price >= pSellMin {
				treeBuy := ob.Buy[order.Price]
				ordBuyId, _ := treeBuy.Min()
				// lay ra tree cua gia ban nho nhat
				treeSell := ob.Sell[pSellMin]
				// lay ra order dau tien cua tree
				ordSellId, qtySell := treeSell.Min()
				qtySellF := qtySell.(float64)
				// so sanh so luong mua va ban
				if order.Qty > qtySellF {
					// cap nhat so luong mua
					order.Qty -= qtySellF
					treeBuy.Put(ordBuyId, order.Qty)

					// xoa order khoi tree sell
					treeSell.Remove(ordSellId)
					if treeSell.Empty() {
						delete(ob.Sell, pSellMin)
						mutils.RemoveAsc(ob.SP, pSellMin)
					}

					// tao message
					msgs = append(msgs, []types.PipeMsg{
						{
							ID:     ordSellId.(uint64),
							PMatch: order.Price,
							Filled: qtySellF,
							Remain: 0,
						},
						{
							ID:     order.ID,
							PMatch: pSellMin,
							Filled: qtySellF,
							Remain: order.Qty,
						},
					}...)

					// set lai pSellMin
					if ob.SP.Len() == 0 {
						break
					}
					pSellMin = ob.SP.Front()
					continue
				}

				if order.Qty == qtySellF {
					// xoa order khoi tree sell
					treeSell.Remove(ordSellId)
					if treeSell.Empty() {
						delete(ob.Sell, pSellMin)
						mutils.RemoveAsc(ob.SP, pSellMin)
					}

					// xoa order khoi tree buy
					treeBuy.Remove(ordBuyId)
					if treeBuy.Empty() {
						delete(ob.Buy, order.Price)
						mutils.RemoveAsc(ob.BP, order.Price)
					}

					// tao message
					return []types.PipeMsg{
						{
							ID:     ordSellId.(uint64),
							PMatch: order.Price,
							Filled: qtySellF,
							Remain: 0,
						},
						{
							ID:     order.ID,
							PMatch: pSellMin,
							Filled: qtySellF,
							Remain: 0,
						},
					}
				} else {
					// cap nhat so luong ban
					treeSell.Put(ordSellId, qtySellF-order.Qty)

					// xoa order khoi tree buy
					treeBuy.Remove(ordBuyId)
					if treeBuy.Empty() {
						delete(ob.Buy, order.Price)
						mutils.RemoveAsc(ob.BP, order.Price)
					}

					// tao message
					return []types.PipeMsg{
						{
							ID:     ordSellId.(uint64),
							PMatch: order.Price,
							Filled: order.Qty,
							Remain: qtySellF - order.Qty,
						},
						{
							ID:     order.ID,
							PMatch: pSellMin,
							Filled: order.Qty,
							Remain: 0,
						},
					}
				}
			}
		}
	} else {
		tree, exist := ob.Sell[order.Price]
		if exist {
			tree.Put(order.ID, order.Qty)
		} else {
			ob.Sell[order.Price] = treemap.NewWith(utils.UInt64Comparator)
			ob.Sell[order.Price].Put(order.ID, order.Qty)
			mutils.InsertAsc(ob.SP, order.Price)

			if ob.BP.Len() == 0 {
				return nil
			}
			pBuyMax := ob.BP.Back()
			for order.Price <= pBuyMax {
				treeSell := ob.Sell[order.Price]
				ordSellId, _ := treeSell.Min()
				treeBuy := ob.Buy[pBuyMax]
				ordBuyId, qtyBuy := treeBuy.Min()
				qtyBuyF := qtyBuy.(float64)
				if order.Qty > qtyBuyF {
					order.Qty -= qtyBuyF
					treeSell.Put(ordSellId, order.Qty)

					treeBuy.Remove(ordBuyId)
					if treeBuy.Empty() {
						delete(ob.Buy, pBuyMax)
						mutils.RemoveAsc(ob.BP, pBuyMax)
					}

					msgs = append(msgs, []types.PipeMsg{
						{
							ID:     ordBuyId.(uint64),
							PMatch: order.Price,
							Filled: qtyBuyF,
							Remain: 0,
						},
						{
							ID:     order.ID,
							PMatch: pBuyMax,
							Filled: qtyBuyF,
							Remain: order.Qty,
						},
					}...)

					if ob.BP.Len() == 0 {
						break
					}
					pBuyMax = ob.BP.Back()
					continue
				}

				if order.Qty == qtyBuyF {
					treeBuy.Remove(ordBuyId)
					if treeBuy.Empty() {
						delete(ob.Buy, pBuyMax)
						mutils.RemoveAsc(ob.BP, pBuyMax)
					}

					treeSell.Remove(ordSellId)
					if treeSell.Empty() {
						delete(ob.Sell, order.Price)
						mutils.RemoveAsc(ob.SP, order.Price)
					}

					return []types.PipeMsg{
						{
							ID:     ordBuyId.(uint64),
							PMatch: order.Price,
							Filled: qtyBuyF,
							Remain: 0,
						},
						{
							ID:     order.ID,
							PMatch: pBuyMax,
							Filled: qtyBuyF,
							Remain: 0,
						},
					}
				} else {
					treeBuy.Put(ordBuyId, qtyBuyF-order.Qty)

					treeSell.Remove(ordSellId)
					if treeSell.Empty() {
						delete(ob.Sell, order.Price)
						mutils.RemoveAsc(ob.SP, order.Price)
					}

					return []types.PipeMsg{
						{
							ID:     ordBuyId.(uint64),
							PMatch: order.Price,
							Filled: order.Qty,
							Remain: qtyBuyF - order.Qty,
						},
						{
							ID:     order.ID,
							PMatch: pBuyMax,
							Filled: order.Qty,
							Remain: 0,
						},
					}
				}
			}
		}
	}
	return msgs
}

func (ob *OrderBook) RemoveOrder(order types.Order) {
	if order.Side == B {
		tree := ob.Buy[order.Price]
		tree.Remove(order.ID)
		if tree.Empty() {
			delete(ob.Buy, order.Price)
			mutils.RemoveAsc(ob.BP, order.Price)
		}
	} else {
		tree := ob.Sell[order.Price]
		tree.Remove(order.ID)
		if tree.Empty() {
			delete(ob.Sell, order.Price)
			mutils.RemoveAsc(ob.SP, order.Price)
		}
	}
}

func (ob *OrderBook) GetAskDepth(size int) [][]float64 {
	depth := [][]float64{}
	for i := 0; i < ob.SP.Len(); i++ {
		p := ob.SP.At(i)
		tree := ob.Sell[p]
		for _, v := range tree.Values() {
			depth = append(depth, []float64{p, v.(float64)})
			if len(depth) == size {
				return depth
			}
		}
	}
	return depth
}

func (ob *OrderBook) GetBidDepth(size int) [][]float64 {
	depth := [][]float64{}
	for i := ob.BP.Len() - 1; i >= 0; i-- {
		p := ob.BP.At(i)
		tree := ob.Buy[p]
		for _, v := range tree.Values() {
			depth = append(depth, []float64{p, v.(float64)})
			if len(depth) == size {
				return depth
			}
		}
	}
	return depth
}

func (ob *OrderBook) GetOrderBook() {
	fmt.Println("=============Buy============")
	for k, v := range ob.Buy {
		fmt.Println("Price: ", k)
		for _, k := range v.Keys() {
			qty, _ := v.Get(k)
			fmt.Println("ID: ", k, " Qty: ", qty)
		}
	}
	fmt.Println("=============Sell============")
	for k, v := range ob.Sell {
		fmt.Println("Price: ", k)
		for _, k := range v.Keys() {
			qty, _ := v.Get(k)
			fmt.Println("ID: ", k, " Qty: ", qty)
		}
	}
}
