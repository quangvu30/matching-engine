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
		{name: "Add bid 100 3", order: types.Order{ID: 1, Price: 100, Qty: 3, Side: 0}, expected: []types.ResultMatching{}},
		{name: "Add bid 100 2", order: types.Order{ID: 2, Price: 100, Qty: 2, Side: 0}, expected: []types.ResultMatching{}},
		{name: "Add ask 101 1", order: types.Order{ID: 3, Price: 101, Qty: 1, Side: 1}, expected: []types.ResultMatching{}},
		{name: "Add ask 101 2", order: types.Order{ID: 4, Price: 101, Qty: 2, Side: 1}, expected: []types.ResultMatching{}},
		{name: "Add bid 102 1", order: types.Order{ID: 5, Price: 102, Qty: 1, Side: 0}, expected: []types.ResultMatching{{ID: 3, PMatch: 101, Filled: 1, Remain: 0}, {ID: 5, PMatch: 101, Filled: 1, Remain: 0}}},
	}
	appl := NewOrderBook()
	var msgs []types.ResultMatching
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			msgs = appl.AddLimitOrder(tc.order)
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

func TestRemoveOrder(t *testing.T) {
	testcase := []TestCase{
		{name: "Add bid 100 3", order: types.Order{ID: 1, Price: 100, Qty: 3, Side: 0}, expected: []types.ResultMatching{}},
		{name: "Add bid 100 2", order: types.Order{ID: 2, Price: 100, Qty: 2, Side: 0}, expected: []types.ResultMatching{}},
		{name: "Add ask 101 1", order: types.Order{ID: 3, Price: 101, Qty: 1, Side: 1}, expected: []types.ResultMatching{}},
		{name: "Add ask 101 2", order: types.Order{ID: 4, Price: 101, Qty: 2, Side: 1}, expected: []types.ResultMatching{}},
		{name: "Add bid 102 1", order: types.Order{ID: 5, Price: 102, Qty: 1, Side: 1}, expected: []types.ResultMatching{}},
		{name: "Remove ID 3", order: types.Order{ID: 3, Price: 101, Qty: 1, Side: 1}, expected: []types.ResultMatching{}},
	}
	appl := NewOrderBook()
	var msgs []types.ResultMatching
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "Remove ID 3" {
				appl.RemoveOrder(tc.order)
				res := appl.GetAskDepth(5)
				if len(res) != 2 {
					t.Fatalf("Expected 2, but got %d", len(res))
				}
				if res[0][0] != 101 || res[0][1] != 2 {
					t.Fatalf("Expected [101 2], but got %v", res[0])
				}
			} else {
				msgs = appl.AddLimitOrder(tc.order)
				if len(msgs) != len(tc.expected) {
					t.Fatalf("Expected %d messages, but got %d", len(tc.expected), len(msgs))
				}
			}
		})
	}
}
