package orderbook

import (
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

type marketFeeder struct {
	matchHandy bool
}

type OrderFeed struct {
	Volume float64 `json:"volume"`
	Price  float64 `json:"price"`
	Sell   bool    `json:"sell"`
	Type   int     `json:"type"`
}

func NewMarketFeeder() *marketFeeder {
	return &marketFeeder{matchHandy: false}
}
func (b *marketFeeder) StartHandyMatchAfter(t time.Duration) {
	go func() {
		time.Sleep(t)
		b.matchHandy = true
	}()
}

func (b *marketFeeder) StartSending(ch chan *Order) {

	go func() {
		for {
			in := b.RandomOrderMatchSell()
			o1 := &Order{
				ID:        string(time.Now().UnixMicro()),
				Volume:    decimal.NewFromFloat32(float32(in.Volume)),
				Price:     decimal.NewFromFloat32(float32(in.Price)),
				IsBid:     !in.Sell,
				EntryTime: time.Now().UnixNano(),
				Type:      in.Type,
				Status:    Order_Status_New,
			}
			ch <- o1
		}
	}()

	go func() {
		for {
			in := b.RandomOrderMatchBuy()
			o1 := &Order{
				ID:        string(time.Now().UnixMicro()),
				Volume:    decimal.NewFromFloat32(float32(in.Volume)),
				Price:     decimal.NewFromFloat32(float32(in.Price)),
				IsBid:     !in.Sell,
				EntryTime: time.Now().UnixNano(),
				Type:      in.Type,
				Status:    Order_Status_New,
			}
			ch <- o1
		}
	}()
}

func (b *marketFeeder) RandomOrderMatchBuy() OrderFeed {
	// 1 is market
	// 0 is limit
	p1 := rand.Float32() * 100
	if b.matchHandy {
		if p1 < 50 {
			p1 += 50
		}
	} else {
		if p1 > 50 {
			p1 -= 50
		}
	}
	p := int(p1)
	return OrderFeed{Volume: float64(int(rand.Float64() * 100)), Price: float64(p), Sell: false, Type: 0}
}

func (b *marketFeeder) RandomOrderMatchSell() OrderFeed {
	// 1 is market
	// 0 is limit

	p1 := rand.Float32() * 100
	if b.matchHandy {
		if p1 > 50 {
			p1 -= 50
		}
	} else {
		if p1 < 50 {
			p1 += 50
		}
	}
	p := int(p1)
	return OrderFeed{Volume: float64(int(rand.Float64() * 100)), Price: float64(p), Sell: true, Type: 0}
}
