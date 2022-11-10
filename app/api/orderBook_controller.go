package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"rabex/api/pb/order_engine"
	oe "rabex/api/pb/order_engine"
	cs "rabex/internal/conversions"
	"rabex/orderbook"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

type orderBookController struct {
	redisConn     *redis.Client
	marketManager *orderbook.MarketManager
	redisPubTrade string
	redisSubOrder string
	order_engine.UnimplementedOrderEngineServer
}

func NewOrderBookController(rc *redis.Client, mm *orderbook.MarketManager,
	redisPubTradeChanName string, redisSubOrderChanName string) *orderBookController {
	return &orderBookController{
		redisConn:     rc,
		marketManager: mm,
		redisPubTrade: redisPubTradeChanName,
		redisSubOrder: redisSubOrderChanName,
	}
}

func (o *orderBookController) NewOrder(ctx context.Context, req *oe.NewOrderRequest) (*oe.NewOrderResponse, error) {
	market, isExist := o.marketManager.GetMarketIfExists(req.Order.Market.Id)
	if !isExist {
		return (*oe.NewOrderResponse)(nil), errors.New("market does not exist")
	}
	order := cs.ConvertOrders(req.Order)
	err := market.ExecuteOrder(order)
	return &oe.NewOrderResponse{Order: &oe.Order{
		ID:          order.ID,
		Price:       order.Price.String(),
		BaseQty:     order.Volume.String(),
		ExecutedQty: order.ExecutedVolume.String(),
		Status:      oe.OrderStatus(order.Status),
	}}, err
}

func (o *orderBookController) CancelOrder(ctx context.Context, req *oe.CancelOrderRequest) (*oe.CancelOrderResponse, error) {
	market, isExist := o.marketManager.GetMarketIfExists(req.Order.Market.Id)
	if !isExist {
		return (*oe.CancelOrderResponse)(nil), errors.New("market doesnt exist")
	}

	market.CancelOrder(cs.ConvertOrders(req.Order))
	return (*oe.CancelOrderResponse)(nil), nil
}

func (o *orderBookController) InitiateMarket(ctx context.Context, req *oe.InitiateMarketRequest) (*oe.InitiateMarketResponse, error) {
	go o.marketManager.StartNewMarket(orderbook.NewMarket(req.Market.Id, req.Market.Base, req.Market.Quote))
	return &oe.InitiateMarketResponse{}, nil
}

func (o *orderBookController) GetAllMarkets(ctx context.Context, req *oe.GetAllMarketsRequest) (*oe.GetAllMarketsResponse, error) {
	output := make([]*order_engine.Market, 0)
	for _, val := range o.marketManager.GetAllMarkets() {
		output = append(output, &order_engine.Market{
			Id:    val.ID,
			Base:  val.BaseSymbol,
			Quote: val.QuoteSymbol,
		})
	}
	return &oe.GetAllMarketsResponse{
		Markets: output,
	}, nil
}

func (o *orderBookController) SubscribeToRedis() {
	sub := o.redisConn.Subscribe(context.Background(), o.redisSubOrder)
	for {
		data, _ := sub.ReceiveMessage(context.Background())
		var orderReq *order_engine.Order
		if data == nil {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		json.Unmarshal([]byte(data.Payload), orderReq)
		if orderReq.MetaData == oe.OrderMetaData_Order_Feeding_ended {
			break
		}
		order := cs.ConvertOrders(orderReq)
		market, exist := o.marketManager.GetMarketIfExists(orderReq.Market.Id)
		if !exist {
			continue
		}
		market.ExecuteOrder(order)
	}
}

func (o *orderBookController) PublishToRedis() {
	for {
		trade := <-o.marketManager.AllOfTrades
		msgtrade := cs.ConvertTrades(trade)
		res, err := json.Marshal(msgtrade)
		if err != nil {
			o.redisConn.Publish(context.Background(), o.redisPubTrade, fmt.Errorf("unable to marshal trade maker : %s, taker : %s", trade.Maker.ID, trade.Taker.ID))
			continue
		}
		o.redisConn.Publish(context.Background(), o.redisPubTrade, res)
	}
}

func RegisterGrpc(s *grpc.Server, oes oe.OrderEngineServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8183))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	order_engine.RegisterOrderEngineServer(s, oes)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func testGrpcMethods() {
	// conn, err := grpc.Dial("127.0.0.1:8181", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	// c := NewGreeterClient(conn)

	// // Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()
	// r, err := c.SayHello(ctx)
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())
	// RegisterGreeterServer(&grpc.Server{}, UnimplementedGreeterServer{})
}
