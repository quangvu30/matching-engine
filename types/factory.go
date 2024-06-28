package types

type RawOrder struct {
	ID        uint64
	Code      string
	Price     float64
	Qty       float64
	Side      int8
	Type      int8
	StopPrice float64
}
