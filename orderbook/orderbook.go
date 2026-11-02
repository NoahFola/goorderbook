package orderbook

import (
	"math/rand"
	"time"
)

type Limit struct {
	price       float64
	orders      []*Order
	totalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		price:  price,
		orders: []*Order{},
	}
}

func (l *Limit) fillOrder()

type Order struct {
	id         int64
	size       float64
	timestamp  int64
	isBid      bool
	limitIndex int
}

func NewOrder(isBid bool, size float64) *Order {
	return &Order{
		id:        rand.Int63n(100000),
		size:      size,
		timestamp: time.Now().UnixNano(),
		isBid:     isBid,
	}
}

func (l *Limit) addOrder(o *Order) {
	o.limitIndex = len(l.orders)
	l.orders = append(l.orders, o)
	l.totalVolume += o.size
}

func (l *Limit) deleteOrder(o *Order) {
	l.orders[o.limitIndex] = l.orders[len(l.orders)-1]
	l.orders = l.orders[:len(l.orders)-1]
	if !o.isFilled() {
		l.totalVolume -= o.size
	}
}

func NewBidOrder(size float64) *Order {
	return NewOrder(false, size)
}

func NewAskOrder(size float64) *Order {
	return NewOrder(false, size)
}

func (o *Order) isFilled() bool {
	return o.size == 0
}

func maxMinOrder(a, b *Order) (*Order, *Order) {
	if a.size > b.size {
		return a, b
	}
	return b, a
}
