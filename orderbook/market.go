package orderbook

import (
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	WaitTimeBetweenFeedOrders = 1 * time.Second
)

type action struct {
	order  *Order
	action int
}

// each market is representing a pair of symbols
// the caller of package will follow his market by its ID that is unique (so all calls needs reliable market ID)

type Market struct {
	ID           string
	BaseSymbol   string
	QuoteSymbol  string
	reqChan      chan *Order  // new orders
	actionChan   chan *action // action chan
	tradeChan    chan *Trade  // trades output
	buyDs        LimitDataStructure
	sellDs       LimitDataStructure
	limitIndexes map[decimal.Decimal]*Limit
	isClosed     bool
}

func NewMarket(id, pair1, pair2 string) *Market {
	return &Market{
		ID:           id,
		BaseSymbol:   pair1,
		QuoteSymbol:  pair2,
		reqChan:      make(chan *Order),
		actionChan:   make(chan *action),
		buyDs:        NewPQ(100, false),
		sellDs:       NewPQ(100, true),
		limitIndexes: make(map[decimal.Decimal]*Limit),
		isClosed:     false,
	}
}

// each market and limit orders will be inserted to the system from here (except feeding)
func (s *Market) ExecuteOrder(o *Order) error {
	select {
	case <-time.After(time.Second * 1):
		return errors.New("unable to insert order for 1 seconds (order igonred)")
	case s.reqChan <- o:
		return nil
	}
}

func (s *Market) FeedMarket() (chan<- *Order, error) {
	ch := make(chan *Order)
	select {
	case <-time.After(time.Second * 1):
		return nil, errors.New("unable to feed")
	case s.actionChan <- &action{order: nil, action: Market_Action_Feed}:
		go func() {
			for o := range ch {
				select {
				case <-time.After(time.Second * 1):
					return
				case s.reqChan <- o:
				}
			}
		}()
		return ch, nil
	}

}

func (s *Market) CancelOrder(o *Order) error {
	select {
	case <-time.After(time.Second * 1):
		return errors.New("unable to cancel order for 1 seconds (Cancelation igonred)")
	case s.actionChan <- &action{order: o, action: Order_Action_Cancel}:
		return nil
	}
}

func (s *Market) Feed(o *Order) error {
	select {
	case <-time.After(time.Second * 1):
		return errors.New("unable to cancel order for 1 seconds (Cancelation igonred)")
	case s.reqChan <- o:
		return nil
	}
}

func (s *Market) GetTradeChan() <-chan *Trade {
	return s.tradeChan
}

func (s *Market) startMarket() {
	s.startReqListener()
}

func (s *Market) closeMarket() error {
	act := &action{(*Order)(nil), Market_Action_Close}
	select {
	case s.actionChan <- act:
	case <-time.After(2 * time.Second):
		return errors.New("unable to close market (can not insert an action to market)")
	}
	return nil
}

// each market has just one goroutine that works via dataStructure
// so there its lock free and goroutine will stay at this function listening to orders and actions
func (s *Market) startReqListener() {
	defer close(s.reqChan)
marketLoop:
	for {
		select {
		case act := <-s.actionChan:
			switch act.action {
			case Market_Action_Close:
				s.isClosed = true
				break marketLoop
			case Order_Action_Cancel:
				s.cancelOrder(act.order)
			case Market_Action_Feed:
				for {
					select {
					case o := <-s.reqChan:
						if o.IsBid {
							s.placeBuyOrder(o)
						} else {
							s.placeSellOrder(o)
						}
					case <-time.After(WaitTimeBetweenFeedOrders):
						fmt.Println("feeding is ended")
						break marketLoop
					}
				}

			}
		case o := <-s.reqChan:
			if o.IsBid {
				if o.Type == Market_Order {
					s.matchBuy(o)
				} else {
					if s.sellDs.IsEmpty() {
						s.placeBuyOrder(o)
					} else {
						s.matchBuy(o)
						if o.hasReMatchCondition() {
							s.placeBuyOrder(o)
						}
					}
				}

			} else {
				if o.Type == Market_Order {
					s.matchSell(o)
				} else {
					if s.buyDs.IsEmpty() {
						s.placeSellOrder(o)
					} else {
						s.matchSell(o)
						if o.hasReMatchCondition() {
							s.placeSellOrder(o)
						}
					}
				}
			}
		}
	}
}

func (s *Market) placeBuyOrder(o *Order) {
	limit, found := s.limitIndexes[o.Price]
	if s.buyDs.IsEmpty() || !found {
		limit = s.buyDs.NewLimitFromOrder(o)
		limit.tradeChan = s.tradeChan
		s.buyDs.Insert(limit)
		s.limitIndexes[o.Price] = limit
	}
	limit.orderQueue.Insert(o)
	limit.totalVolume = limit.totalVolume.Add(o.Volume)
}

func (s *Market) cancelOrder(o *Order) error {
	limit, found := s.limitIndexes[o.Price]
	if s.buyDs.IsEmpty() || !found {
		return fmt.Errorf("Cancel Order Not Found, order_id :  %s,", o.ID)
	}
	limit.orderQueue.Cancel(o)
	limit.totalVolume = limit.totalVolume.Sub(o.Volume)
	return nil
}

func (s *Market) placeSellOrder(o *Order) {
	limit, found := s.limitIndexes[o.Price]
	if s.sellDs.IsEmpty() || !found {
		limit = s.sellDs.NewLimitFromOrder(o)
		limit.tradeChan = s.tradeChan
		s.sellDs.Insert(limit)
		s.limitIndexes[o.Price] = limit
	}
	limit.orderQueue.Insert(o)
	limit.totalVolume = limit.totalVolume.Add(o.Volume)
}

func (s *Market) matchBuy(o *Order) {
	for s.isReadyToMatch(o) {
		minOrMax := s.sellDs.GetMinOrMax()
		minOrMax.orderQueue.MergeFully(o)
		if minOrMax.orderQueue.isEmpty() || s.sellDs.IsEmpty() {
			if s.sellDs.IsEmpty() {
				break
			} else {
				s.sellDs.DelMinOrMax()
				if s.sellDs.IsEmpty() {
					break
				}
			}
		}
	}
}

func (s *Market) matchSell(o *Order) {
	for s.isReadyToMatch(o) {
		minOrMax := s.buyDs.GetMinOrMax()
		minOrMax.orderQueue.MergeFully(o)
		if minOrMax.orderQueue.isEmpty() || s.buyDs.IsEmpty() {
			if s.buyDs.IsEmpty() {
				break
			} else {
				s.buyDs.DelMinOrMax()
				if s.buyDs.IsEmpty() {
					break
				}
			}
		}
	}
}

func (s *Market) isReadyToMatch(o *Order) bool {
	res := false
	if o.IsBid {
		firstOrderPrice := s.sellDs.GetMinOrMax().Price
		res = s.sellDs.GetMinOrMax() != nil && o.hasMatchCondition()
		switch o.Type {
		case Limit_Order:
			return res && firstOrderPrice.Cmp(o.Price) != 1
		case Market_Order:
			return res
		default: // pretend limit
			return res && firstOrderPrice.Cmp(o.Price) != 1
		}

	} else {
		firstOrderPrice := s.buyDs.GetMinOrMax().Price
		res = s.buyDs.GetMinOrMax() != nil && o.hasMatchCondition()
		switch o.Type {
		case Limit_Order:
			return res && firstOrderPrice.Cmp(o.Price) != -1
		case Market_Order:
			return res
		default: // pretend limit
			return res && firstOrderPrice.Cmp(o.Price) != -1
		}
	}
}
