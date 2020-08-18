package tokensearcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func Test_assetsMap(t *testing.T) {
	tx1 := blockatlas.Tx{
		Coin: 60,
		From: "A",
		To:   "B",
		Meta: blockatlas.NativeTokenTransfer{
			TokenID: "ABC",
			From:    "A",
			To:      "C",
		},
	}

	tx2 := blockatlas.Tx{
		Coin: 60,
		From: "D",
		To:   "V",
		Meta: blockatlas.TokenTransfer{
			TokenID: "EFG",
			From:    "D",
			To:      "F",
		},
	}

	tx3 := blockatlas.Tx{
		Coin: 60,
		From: "Q",
		To:   "L",
		Meta: blockatlas.AnyAction{
			TokenID: "HIJ",
		},
	}

	result := assetsMap(blockatlas.Txs{tx1, tx2, tx3})
	assert.Equal(t, result["60_A"], []string{"c60_tABC"})
	assert.Equal(t, result["60_C"], []string{"c60_tABC"})
	assert.Equal(t, result["60_D"], []string{"c60_tEFG"})
	assert.Equal(t, result["60_F"], []string{"c60_tEFG"})
	assert.Equal(t, result["60_Q"], []string{"c60_tHIJ"})
	assert.Equal(t, result["60_L"], []string{"c60_tHIJ"})
}

func Test_associationsToAdd(t *testing.T) {
	o := make(map[string][]string)
	n := make(map[string][]string)

	o["A"] = []string{"1", "2", "3"}
	o["B"] = []string{"3", "4", "5"}

	n["A"] = []string{"1", "2", "5"}
	n["B"] = []string{"3", "9", "8"}

	result := associationsToAdd(o, n)

	assert.Equal(t, result["A"], []string{"5"})
	assert.Equal(t, result["B"], []string{"9", "8"})
}
