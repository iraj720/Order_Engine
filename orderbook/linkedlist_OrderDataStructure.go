package orderbook

import (
	"github.com/shopspring/decimal"
)

type linkedList struct {
	head      *OrderNode
	tail      *OrderNode
	tradeChan chan *Trade
	size      int64
	limit     *Limit
}

type OrderNode struct {
	*Order
	next *OrderNode
	prev *OrderNode
}

func NewLinkedList(l *Limit) *linkedList {
	head := NewOrderNode()
	tail := NewOrderNode()
	head.next = tail
	tail.prev = head
	return &linkedList{head: head, tail: tail, limit: l, size: 0}
}

func (l *linkedList) Insert(o *Order) {
	on := &OrderNode{Order: o}
	on.prev = l.tail.prev
	on.next = l.tail
	l.tail.prev.next = on
	l.tail.prev = on
	l.size++
}

func (l *linkedList) Cancel(o *Order) {
	l.remove(l.find(o.ID))
	l.size--
	o.Status = Order_Status_Canceled
}

func (l *linkedList) GetMatchedOrders() <-chan *Trade {
	return l.tradeChan
}

func (l *linkedList) MergeOne(o *Order) {
	// SecondOrder := l.head.next
	// if SecondOrder != nil {
	// 	tradedVolume := l.matchOrder(o, SecondOrder.Order)
	// 	o.Status = Order_Status_Partially_Filled
	// 	SecondOrder.Status = Order_Status_Partially_Filled
	// 	l.limit.totalVolume -= tradedVolume
	// 	if o.Volume == 0 {
	// 		o.Status = Order_Status_Filled
	// 	} else {
	// 		SecondOrder.Status = Order_Status_Filled
	// 		SecondOrder.next.prev = l.head
	// 		l.head.next = SecondOrder.next
	// 	}
	// }

	// l.sendTradeToChannel(NewTrade(o, SecondOrder.Order))
}

func (l *linkedList) MergeFully(o *Order) {
	tradeVolume := decimal.NewFromFloat32(0.0)
	order := l.head
	o.Status = Order_Status_Partially_Filled
	for order.next != l.tail && tradeVolume.Cmp(o.Volume) == -1 {
		order = order.next
		order.Status = Order_Status_Filled
		tradeVolume = tradeVolume.Add(order.Volume)
		if tradeVolume.Cmp(o.Volume) == 1 {
			order.Volume = tradeVolume.Sub(o.Volume)
			order.Status = Order_Status_Partially_Filled
			tradeVolume = o.Volume
		}
		l.sendTradeToChannel(NewTrade(o, order.Order))
	}
	l.limit.totalVolume = l.limit.totalVolume.Sub(tradeVolume)
	o.Volume = o.Volume.Sub(tradeVolume)
	if order.Status == Order_Status_Partially_Filled {
		order.prev = l.head
		l.head.next = order
	} else if order.Status == Order_Status_Filled {
		order.next.prev = l.head
		l.head.next = order.next
	}
}

func (l *linkedList) First() *OrderNode {
	return l.head.next
}

func (l *linkedList) Last() *OrderNode {
	return l.tail.prev
}

func (l *linkedList) find(orderId string) *OrderNode {
	element := l.head.next
	for element.ID != orderId {
		element = element.next
	}
	return element
}

func (l *linkedList) remove(o *OrderNode) {
	o.next.prev = o.prev
	o.prev.next = o.next
	o = nil
}

// func (l *linkedList) matchOrder(o1, o2 *Order) decimal.Decimal {
// 	// making trade object and send it to redis
// 	tradeVolume := math.Min(float64(o1.Price), float64(o2.Price))
// 	o1.Volume -= tradeVolume
// 	o2.Volume -= tradeVolume
// 	return tradeVolume
// }

func (l *linkedList) sendTradeToChannel(t *Trade) {
	select {
	case l.limit.tradeChan <- t:
	default:
	}
}

func (l *linkedList) isEmpty() bool {
	return l.limit.totalVolume == decimal.NewFromInt(0)
}
