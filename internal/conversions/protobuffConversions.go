package conversions

import (
	oe "rabex/api/pb/order_engine"
	"rabex/orderbook"

	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertOrders(o *oe.Order) *orderbook.Order {
	return &orderbook.Order{
		ID:             o.ID,
		IsBid:          o.Side == oe.OrderSide_Order_Side_Buy,
		Type:           int(o.Type),
		Price:          decimal.RequireFromString(o.Price),
		Volume:         decimal.RequireFromString(o.BaseQty),
		ExecutedVolume: decimal.RequireFromString(o.ExecutedQty),
		Status:         int(o.Status),
	}
}

func ConvertOrders2(o *orderbook.Order) *oe.Order {
	return &oe.Order{
		ID:          o.ID,
		Price:       o.Price.String(),
		ExecutedQty: o.ExecutedVolume.String(),
		Status:      oe.OrderStatus(o.Status),
	}
}

func ConvertTrades(t *orderbook.Trade) *oe.Trade {
	return &oe.Trade{
		Id:           string(t.ID),
		MakerId:      t.Maker.ID,
		TakerId:      t.Taker.ID,
		Volume:       t.Volume.String(),
		Price:        t.Price.String(),
		CreationTime: timestamppb.New(t.CreationTime),
	}
}
