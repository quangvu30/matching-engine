package utils

import (
	"math/rand"

	"github.com/quangvu30/matching-engine/types"
)

func RandomOrder(i int) types.Order {
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
