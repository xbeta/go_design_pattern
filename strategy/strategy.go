package strategy

import (
	"math/rand"
)

const (
	GUU = iota
	CHO
	PAA
)

var hands []*hand

func init() {
	hands = []*hand{
		&hand{GUU},
		&hand{CHO},
		&hand{PAA},
	}
}

type hand struct {
	handValue int
}

func getHand(handValue int) *hand {
	return hands[handValue]
}

func (self *hand) isStrongerThan(h *hand) bool {
	return self.fight(h) == 1
}

func (self *hand) isWeakerThan(h *hand) bool {
	return self.fight(h) == -1
}

func (self *hand) fight(h *hand) int {
	if self == h {
		return 0
	} else if (self.handValue+1)%3 == h.handValue {
		return 1
	} else {
		return -1
	}
}

// Goでは関数がファーストクラスであるため、戦略がひとつのメソッドのみで
// 完結する場合は、関数を渡すことでStrategyパターンとすることもできる。
// type strategy func() *hand
type strategy interface {
	nextHand() *hand
	study(win bool)
}

type winningStrategy struct {
	seed     int64
	won      bool
	prevHand *hand
}

func (self *winningStrategy) nextHand() *hand {
	if !self.won {
		// rand.Seed(self.seed)
		self.prevHand = getHand(rand.Intn(3))
	}
	return self.prevHand
}

func (self *winningStrategy) study(win bool) {
	self.won = win
}

type player struct {
	name                           string
	strategy                       strategy
	wincount, losecount, gamecount int
}

func (self *player) nextHand() *hand {
	return self.strategy.nextHand()
}

func (self *player) win() {
	self.wincount++
	self.gamecount++
}

func (self *player) lose() {
	self.losecount++
	self.gamecount++
}

func (self *player) even() {
	self.gamecount++
}