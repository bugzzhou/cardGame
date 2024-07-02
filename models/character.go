package models

import (
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

func (c *Character) PlayCard(index int) {}

func (c *Character) Restart() {}

func (c *Character) GetCard() {}
