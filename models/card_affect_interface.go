package models

//用于存放所有牌的作用函数

type AffectFuncInter interface {
	Check(c1, c2 *Character) bool
	Run(c1, c2 *Character)
}

var _ AffectFuncInter = (*AffectFunc1)(nil)
var _ AffectFuncInter = (*AffectFunc2)(nil)
