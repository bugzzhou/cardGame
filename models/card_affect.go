package models

import "fmt"

type AffectFunc1 struct{}
type AffectFunc2 struct{}

func (a *AffectFunc1) Check(c1, c2 *Character) bool {
	return true
}

func (a *AffectFunc1) Run(c1, c2 *Character) {
	defaultDamage := 7
	power := 0
	mul := 1.0

	c1Buff := c1.Buffs
	c2Debuff := c2.Debuffs

	for _, v := range c1Buff {
		if v.TypeId == 1 {
			power = v.Value
			break
		}
	}

	for _, v := range c2Debuff {
		if v.TypeId == 102 {
			mul *= 1.5
		}
	}

	fmt.Println(int(float64(defaultDamage+power) * mul))

	c2.Health -= int(float64(defaultDamage+power) * mul)
}

func (a *AffectFunc2) Check(c1, c2 *Character) bool {
	return true
}

func (a *AffectFunc2) Run(c1, c2 *Character) {
	c1.Shield += 5
}
