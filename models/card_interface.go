package models

type CardInter interface {
}

var _ CardInter = (*Card)(nil)
