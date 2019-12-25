package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WalletCalculator(t *testing.T) {
	c := NewWalletCalculator()

	result, err := c.Calculate(100, 10, "win")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(110), result)
}
