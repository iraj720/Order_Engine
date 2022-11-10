package main

import oe "rabex/api/pb/order_engine"

func possibleOrders() []oe.Order {
	possibleIncomingOrders := []oe.Order{

		// we assume that at order initiate time '1 BTC == 1000 USDT'

		// buy orders :
		{
			Base:        "BTC",
			Quote:       "USDT",
			Price:       "1000",
			FreezeAsset: "USDT",
			FreezeQty:   "1000",
			BaseQty:     "0",
			QuoteQty:    "1000",
			Side:        oe.OrderSide_Order_Side_Buy,
			Type:        oe.OrderType_Order_Type_Limit,
			Status:      oe.OrderStatus_Order_Status_New,
		},
		// we should check that 1 BTC wont be more than 1000 USDT
		{
			Base:        "BTC",
			Quote:       "USDT",
			Price:       "1000",
			FreezeAsset: "USDT",
			FreezeQty:   "1000",
			BaseQty:     "1",
			QuoteQty:    "0",
			Side:        oe.OrderSide_Order_Side_Buy,
			Type:        oe.OrderType_Order_Type_Limit,
			Status:      oe.OrderStatus_Order_Status_New,
		},
		// sell orders :
		{
			Base:        "BTC",
			Quote:       "USDT",
			Price:       "1000",
			FreezeAsset: "BTC",
			FreezeQty:   "1",
			BaseQty:     "1",
			QuoteQty:    "0",
			Side:        oe.OrderSide_Order_Side_Sell,
			Type:        oe.OrderType_Order_Type_Limit,
			Status:      oe.OrderStatus_Order_Status_New,
		},
		// we should check that 1 BTC wont be more than 1000 USDT
		{
			Base:        "BTC",
			Quote:       "USDT",
			Price:       "1000",
			FreezeAsset: "BTC",
			FreezeQty:   "1",
			BaseQty:     "0",
			QuoteQty:    "1000",
			Side:        oe.OrderSide_Order_Side_Sell,
			Type:        oe.OrderType_Order_Type_Limit,
			Status:      oe.OrderStatus_Order_Status_New,
		},
	}
	return possibleIncomingOrders
}
