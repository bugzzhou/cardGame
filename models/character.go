package models

import (
	"fmt"
	comm "sts/common"
)

type Character struct {
	Name        string
	HealthLimit int
	Health      int
	EnergyLimit int
	Energy      int

	Power     int
	Dexterity int
	Shield    int

	Buffs   []Effect
	Debuffs []Effect

	DrawNum      int
	AllCards     []Card
	DrawCards    []Card
	DiscardCards []Card
	HandCards    []Card
	Lotions      []Lotion
	Props        []Prop
}

func (c *Character) DrawCard(drawNum int) {
	// 弃牌堆 - > 抽牌堆
	if len(c.DrawCards) < drawNum {
		c.DrawCards = append(c.DrawCards, c.DiscardCards...)
		c.DiscardCards = nil
		comm.R.Shuffle(len(c.DrawCards), func(i, j int) {
			c.DrawCards[i], c.DrawCards[j] = c.DrawCards[j], c.DrawCards[i]
		})
	}

	// 抽牌堆 - > 手牌
	if len(c.DrawCards) <= drawNum {
		c.HandCards = append(c.HandCards, c.DrawCards...)
		c.DrawCards = nil
	} else {
		c.HandCards = append(c.HandCards, c.DrawCards[:drawNum]...)
		c.DrawCards = c.DrawCards[drawNum:]
	}
}

func (c *Character) EndTurn() {
	c.DiscardCards = append(c.DiscardCards, c.HandCards...)
	c.HandCards = nil
}

// TODO bugzzhou 添加虚无、消耗 的牌类型，用于在play、end的时候进行特殊处理（放到消耗牌堆 & 还需要增加消耗牌堆字段）
func (c *Character) PlayCard(index int, enemy *Character) {
	if len(c.HandCards) < index {
		fmt.Printf("failed to play card, out of length of the handCard\n")
		return
	}

	c.HandCards[index].AffectFunc.Run(c, enemy)

	c.DiscardCards = append(c.DiscardCards, c.HandCards[index])
	c.HandCards = append(c.HandCards[:index], c.HandCards[index+1:]...)
}

func (c *Character) Restart() {}

func (c *Character) GetCard() {}
