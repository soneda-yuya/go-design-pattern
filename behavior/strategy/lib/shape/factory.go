package shape

import (
	"fmt"
	"os"
	"../strategy"
)

const (
  TEXT_STRATEGY  = "text"
  IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (strategy.PrintStrategy, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			strategy.PrintOutput{
				Writer:    os.Stdout,
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}