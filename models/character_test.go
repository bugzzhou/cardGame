package models

import (
	"fmt"
	comm "sts/common"
	"testing"
)

const DefaultDrawNum = 5

func InitCharacter(tag bool) Character {
	var buff []Effect
	var deBuff []Effect

	if tag {
		buff = []Effect{
			{
				TypeId: 1,
				Value:  10,
			},
		}
		deBuff = []Effect{
			{
				TypeId: 102,
			},
		}
	}

	Card1 := Card{
		ID: 1, Name: "攻击", CardType: "attack", Cost: 1, AffectFunc: &AffectFunc1{},
	}
	Card2 := Card{
		ID: 2, Name: "防御", CardType: "defence", Cost: 1, AffectFunc: &AffectFunc2{},
	}
	return Character{
		Name:        "a",
		HealthLimit: 99,
		Health:      99,
		Energy:      3,

		Buffs:   buff,
		Debuffs: deBuff,

		DrawNum: DefaultDrawNum,
		AllCards: []Card{
			Card1, Card2, Card1, Card2, Card1,
			Card2, Card1, Card2, Card1, Card2,
		},
		DrawCards: []Card{
			Card1, Card2, Card1, Card2, Card1,
			Card2, Card1, Card2, Card1, Card2,
		},
	}
}

func TestDrawCards(t *testing.T) {
	initC := InitCharacter(true)
	initD := InitCharacter(false)

	displayNum(initC, initD)

	initC.DrawCard(5)
	randC := comm.R.Intn(5)
	initC.PlayCard(randC, &initD)

	initD.DrawCard(5)
	randD := comm.R.Intn(5)
	initD.PlayCard(randD, &initC)

	displayNum(initC, initD)

	initC.EndTurn()

}

func displayNum(c, d Character) {
	// fmt.Printf("AllCards' num is: %d\n", len(c.AllCards))
	// fmt.Printf("DrawCards' num is: %d\n", len(c.DrawCards))
	// fmt.Printf("HandCards' num is: %d\n", len(c.HandCards))
	// fmt.Printf("DiscardCards' num is: %d\n", len(c.DiscardCards))
	fmt.Printf("%v %v %v %v\n", c.Health, c.Shield, d.Health, d.Shield)
}
