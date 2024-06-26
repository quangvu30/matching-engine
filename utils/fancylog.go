package utils

import (
	"fmt"

	"github.com/quangvu30/matching-engine/types"
)

func Log(msgs []types.ResultMatching) {
	fmt.Println("-------- Pipe Messages --------")
	for _, msg := range msgs {
		fmt.Printf("ID: %d, PMatch: %f, Filled: %f, Remain: %f\n", msg.ID, msg.PMatch, msg.Filled, msg.Remain)
	}
}
