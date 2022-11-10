package orderbook

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderDataStructure interface {
	Cancel(o *Order)
	Insert(o *Order)
	MergeFully(o *Order)
	MergeOne(o *Order)
	First() *OrderNode
	Last() *OrderNode
	isEmpty() bool
}

type LimitDataStructure interface {
	IsEmpty() bool
	Insert(key *Limit)
	Size() int
	DelMinOrMax()
	Delete(p decimal.Decimal)
	GetMinOrMax() *Limit
	NewLimitFromOrder(o *Order) *Limit
	GetTradeChan() <-chan *Trade
}

// TODO : seperate bussiness models from data structure objects
type Limit struct {
	orderQueue  OrderDataStructure
	tradeChan   chan *Trade
	Price       decimal.Decimal
	totalVolume decimal.Decimal
}

type Order struct {
	limit          *Limit
	Price          decimal.Decimal
	ExecutedVolume decimal.Decimal
	Volume         decimal.Decimal
	ID             string
	EntryTime      int64
	FulfillTime    int64
	Type           int
	Status         int
	IsBid          bool
}

type Trade struct {
	Maker        *Order
	Taker        *Order
	Price        decimal.Decimal
	Volume       decimal.Decimal
	ID           int64
	CreationTime time.Time
}

func NewOrder() *Order {
	return &Order{}
}

func NewOrderNode() *OrderNode {
	return &OrderNode{Order: &Order{}}
}

func NewTrade(taker, maker *Order) *Trade {
	return &Trade{
		Volume:       decimal.Min(taker.Volume, maker.Volume),
		Price:        taker.Price,
		Taker:        taker,
		Maker:        maker,
		CreationTime: time.Now(),
	}
}

func (o *Order) hasMatchCondition() bool {
	return o.Price.Cmp(decimal.NewFromInt(0)) == 1 && o.Volume.Cmp(decimal.NewFromInt(0)) == 1 && o.Status != Order_Status_Canceled &&
		o.Status != Order_Status_Expired && o.Status != Order_Status_Partially_Cancelled && o.Status != Order_Status_Filled
}

func (o *Order) hasReMatchCondition() bool {
	if o.Type == Market_Order {
		return false
	}
	return o.hasMatchCondition()
}
