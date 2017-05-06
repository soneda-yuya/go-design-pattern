package prototype

import (
	"errors"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ItemInfoGetter interface {
	GetInfo() string
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtsCache struct {}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

func (sc *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}