package core

import (
	"testing"

	"github.com/quangvu30/matching-engine/types"
)

type TestCase struct {
	name     string
	order    types.Order
	expected []types.ResultMatching
}

func TestAddOrder(t *testing.T) {
	testcase := []TestCase{
		{name: "Add bid 100 3", order: types.Order{ID: 1, Price: 100, Qty: 3, Side: 0, Type: 0}, expected: []types.ResultMatching{}},
		{name: "Add bid 100 2", order: types.Order{ID: 2, Price: 100, Qty: 2, Side: 0, Type: 0}, expected: []types.ResultMatching{}},
		{name: "Add ask 101 1", order: types.Order{ID: 3, Price: 101, Qty: 1, Side: 1, Type: 0}, expected: []types.ResultMatching{}},
		{name: "Add ask 101 2", order: types.Order{ID: 4, Price: 101, Qty: 2, Side: 1, Type: 0}, expected: []types.ResultMatching{}},
		{name: "Add bid 102 1", order: types.Order{ID: 5, Price: 102, Qty: 1, Side: 0, Type: 0}, expected: []types.ResultMatching{{ID: 3, PMatch: 101, Filled: 1, Remain: 0}, {ID: 5, PMatch: 101, Filled: 1, Remain: 0}}},
	}
	appl := NewOrderBook()
	var msgs []types.ResultMatching
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			msgs = appl.AddOrder(tc.order)
			if len(msgs) != len(tc.expected) {
				t.Fatalf("Expected %d messages, but got %d", len(tc.expected), len(msgs))
			}
			for i, msg := range msgs {
				if msg.ID != tc.expected[i].ID || msg.PMatch != tc.expected[i].PMatch || msg.Filled != tc.expected[i].Filled || msg.Remain != tc.expected[i].Remain {
					t.Fatalf("Expected %v, but got %v", tc.expected[i], msg)
				}
			}
		})
	}
}
