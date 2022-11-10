package orderbook

import (
	"fmt"
	"time"
)

func startOrderBook() {
	mm := NewMarketManager()
	mf := NewMarketFeeder()
	market := NewMarket("1", "ETH", "USDT")
	mm.StartNewMarket(market)
	orderChan := make(chan *Order)
	mf.StartSending(orderChan)

	go func() {
		for val := range market.GetTradeChan() {
			printTrade(val)
		}
	}()

	mf.StartHandyMatchAfter(5 * time.Second)
}

func printTrade(t *Trade) {
	fmt.Println(t.Price, " ", t.Volume, " ", t.Maker.Volume, " ", t.Taker.Volume)
}
