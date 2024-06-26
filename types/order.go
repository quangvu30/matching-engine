package types

type Order struct {
	ID    uint64
	Price float64
	Qty   float64
	Side  int8
	Type  int8
}

type ResultMatching struct {
	ID     uint64
	PMatch float64
	Filled float64
	Remain float64
}
