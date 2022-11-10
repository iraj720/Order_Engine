package orderbook

import (
	"errors"
	"sync"
)

type MarketManager struct {
	markets sync.Map
	// this field will lock all of goroutines that are making trades
	// so in future we can change it and make publish channel per market then there is no syncronization
	AllOfTrades chan *Trade
}

func NewMarketManager() *MarketManager {
	return &MarketManager{
		markets:     sync.Map{},
		AllOfTrades: make(chan *Trade),
	}
}

// this will block and recieve orders forever
func (mm *MarketManager) StartNewMarket(m *Market) {
	mm.markets.Store(m.ID, m)

	go func() {
		mm.AllOfTrades <- <-m.tradeChan
	}()
	m.startMarket()
}

func (mm *MarketManager) GetMarketIfExists(marketID string) (*Market, bool) {
	market, found := mm.markets.Load(marketID)
	if !found {
		return (*Market)(nil), false
	}
	m := market.(*Market)
	return m, true
}

func (mm *MarketManager) CloseMarket(m *Market) error {
	market, found := mm.markets.Load(m.ID)
	if !found {
		return errors.New("market not found close")
	}
	markett := market.(*Market)
	err := markett.closeMarket()
	if err != nil {
		return err
	}
	mm.markets.Delete(markett.ID)
	return nil
}

func (mm *MarketManager) GetAllMarkets() []*Market {
	markets := make([]*Market, 0)
	mm.markets.Range(func(k, v interface{}) bool {
		market := v.(*Market)
		markets = append(markets, market)
		return true
	})
	return markets
}
