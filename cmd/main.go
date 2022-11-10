package main

import (
	"context"
	"encoding/json"
	"fmt"
	oe "rabex/api/pb/order_engine"
	"rabex/app/api"
	"rabex/confs"
	"rabex/orderbook"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	RedisSubChanName = "XX_Order_XX"
	RedisPubChanName = "XX_Trade_XX"
)

func main() {
	srv := grpc.NewServer()
	rc := confs.SetupRedisConnection()
	obc := api.NewOrderBookController(rc, orderbook.NewMarketManager(), RedisPubChanName, RedisSubChanName)
	go api.RegisterGrpc(srv, obc)

	go obc.PublishToRedis()
	go obc.SubscribeToRedis()

	SimulateClient()
}

func SimulateClient() {
	rc := confs.SetupRedisConnection()
	// go func() {
	// 	mf := orderbook.NewMarketFeeder()
	// 	ch := make(chan *orderbook.Order)
	// 	mf.StartHandyMatchAfter(5 * time.Second)
	// 	mf.StartSending(ch)
	// 	for order := range ch {
	// 		orderProto := cs.ConvertOrders2(order)
	// 		orderByte, _ := json.Marshal(orderProto)
	// 		rc.Publish(context.Background(), RedisSubChanName, orderByte)
	// 	}
	// }()

	go func() {
		for {
			sub := rc.Subscribe(context.Background(), RedisPubChanName)
			data, _ := sub.ReceiveMessage(context.Background())
			var trade *oe.Trade
			if data == nil {
				time.Sleep(10 * time.Second)
				continue
			}
			json.Unmarshal([]byte(data.Payload), trade)
			fmt.Println(trade.Id, trade.Price, trade.Volume, trade.MakerId, trade.TakerId)
		}
	}()
	market := &oe.Market{
		Id:    "1234",
		Base:  "BTC",
		Quote: "USDT",
	}
	order1 := &oe.NewOrderRequest{
		Order: &oe.Order{
			Market:      market,
			ID:          "123",
			Price:       "123",
			BaseQty:     "1234",
			ExecutedQty: "0",
			Type:        oe.OrderType_Order_Type_Limit,
			Side:        oe.OrderSide_Order_Side_Buy,
			Status:      oe.OrderStatus_Order_Status_New,
		},
	}

	order2 := &oe.NewOrderRequest{
		Order: &oe.Order{
			Market:      market,
			ID:          "124",
			Price:       "122",
			BaseQty:     "1234",
			ExecutedQty: "0",
			Type:        oe.OrderType_Order_Type_Limit,
			Side:        oe.OrderSide_Order_Side_Sell,
			Status:      oe.OrderStatus_Order_Status_New,
		},
	}
	conn, _ := grpc.Dial("127.0.0.1:8183", grpc.WithTransportCredentials(insecure.NewCredentials()))
	oeClient := oe.NewOrderEngineClient(conn)
	oeClient.InitiateMarket(context.Background(), &oe.InitiateMarketRequest{
		Market: market,
	})
	res, err := oeClient.NewOrder(context.Background(), order1)
	fmt.Println(res, err)
	res, err = oeClient.NewOrder(context.Background(), order2)
	fmt.Println(res, err)
	time.Sleep(100 * time.Second)
}
