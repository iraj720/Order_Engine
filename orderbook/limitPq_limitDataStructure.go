package orderbook

import "github.com/shopspring/decimal"

// mininum oriented Priority Queue
type indexMinPQ struct {
	keys      []Limit
	tradeChan chan *Trade
	n         int
	isMinPq   bool
}

func NewPQ(size int, isMinPQ bool) *indexMinPQ {
	return &indexMinPQ{
		keys:      make([]Limit, size+1),
		n:         0,
		tradeChan: make(chan *Trade),
		isMinPq:   isMinPQ,
	}
}

func (pq *indexMinPQ) Size() int {
	return pq.n
}

func (pq *indexMinPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *indexMinPQ) GetMinOrMax() *Limit {
	return pq.Top()
}

func (pq *indexMinPQ) Insert(key *Limit) {
	if pq.n+1 == cap(pq.keys) {
		panic("pq is full")
	}
	pq.n++
	pq.keys[pq.n] = *key

	// restore order
	if pq.isMinPq {
		pq.swimMin(pq.n)
	} else {
		pq.swimMax(pq.n)
	}
}

func (pq *indexMinPQ) Delete(p decimal.Decimal) {
	pq.DeleteByIdx(pq.find(p))
}

func (pq *indexMinPQ) DeleteByIdx(i int) {
	// replace key with the last key
	pq.keys[i].Price = decimal.NewFromInt(0)
	pq.keys[i].totalVolume = decimal.NewFromInt(0)
	pq.keys[i], pq.keys[pq.n] = pq.keys[pq.n], pq.keys[i]
	pq.n--

	// restore order
	if pq.isMinPq {
		pq.sinkMin(i)
	} else {
		pq.sinkMax(i)
	}
}

func (pq *indexMinPQ) find(p decimal.Decimal) int {
	element := pq.keys[1]
	c := 0
	for element.Price != p && c <= pq.n {
		element = pq.keys[c]
		c *= 2
	}
	return c
}

func (pq *indexMinPQ) Top() *Limit {
	if pq.IsEmpty() {
		panic("pq is empty")
	}

	return &pq.keys[1]
}

// removes minimal element and returns it's index
func (pq *indexMinPQ) DelMinOrMax() {
	pq.DeleteByIdx(1)
}

// helpers
func (pq *indexMinPQ) checkIndex(i int) {
	if i < 0 || i+1 >= cap(pq.keys) {
		panic("invalid index")
	}
}

func (pq *indexMinPQ) swimMax(i int) {
	k := pq.n
	for k > 1 && pq.keys[k].Price.Cmp(pq.keys[k/2].Price) == 1 {
		// swap keys
		pq.keys[k], pq.keys[k/2] = pq.keys[k/2], pq.keys[k]
		k = k / 2
	}
}

func (pq *indexMinPQ) sinkMax(i int) {
	k := i
	for 2*k <= pq.n {
		c := 2 * k

		// select minimum of two children
		if c < pq.n && pq.keys[c+1].Price.Cmp(pq.keys[c].Price) == 1 {
			c++
		}

		if pq.keys[k].Price.Cmp(pq.keys[c].Price) == -1 {
			// swap keys
			pq.keys[k], pq.keys[c] = pq.keys[c], pq.keys[k]
			k = c
		} else {
			break
		}
	}
}

func (pq *indexMinPQ) swimMin(i int) {
	k := pq.n
	for k > 1 && pq.keys[k].Price.Cmp(pq.keys[k/2].Price) == -1 {
		// swap keys
		pq.keys[k], pq.keys[k/2] = pq.keys[k/2], pq.keys[k]
		k = k / 2
	}
}

func (pq *indexMinPQ) sinkMin(i int) {
	k := i
	for 2*k <= pq.n {
		c := 2 * k

		// select minimum of two children
		if c < pq.n && pq.keys[c+1].Price.Cmp(pq.keys[c].Price) == -1 {
			c++
		}

		if pq.keys[k].Price.Cmp(pq.keys[c].Price) == 1 {
			// swap keys
			pq.keys[k], pq.keys[c] = pq.keys[c], pq.keys[k]
			k = c
		} else {
			break
		}
	}
}

func (pq *indexMinPQ) NewLimitFromOrder(o *Order) *Limit {
	limit := &Limit{Price: o.Price, totalVolume: decimal.NewFromInt(0), tradeChan: pq.tradeChan}
	ll := NewLinkedList(limit)
	limit.orderQueue = ll
	return limit
}

func (pq *indexMinPQ) GetTradeChan() <-chan *Trade {
	return pq.tradeChan
}
