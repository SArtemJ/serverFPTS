package calculator

import (
	"errors"
)

type WalletCalculator struct{}

func NewWalletCalculator() *WalletCalculator {
	return &WalletCalculator{}
}

func (cw *WalletCalculator) Calculate(w, t int64, event string) (int64, error) {
	var result int64
	var err error

	switch event {
	case "win":
		result = w + t
	case "lost":
		result = w - t
	}

	if result < 0 {
		return 0, errors.New("ERROR - User wallet can't be negative")
	} else {
		return result, nil
	}

	return result, err
}
