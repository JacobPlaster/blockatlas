package blockbook

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error) {
	tokens, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(tokens, coinIndex), nil
}

func NormalizeTokens(srcTokens []Token, coinIndex uint) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0, len(srcTokens))
	for _, srcToken := range srcTokens {
		if srcToken.Balance == "0" || srcToken.Balance == "" {
			continue
		}
		token := NormalizeToken(&srcToken, coinIndex)
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}

func NormalizeToken(srcToken *Token, coinIndex uint) blockatlas.Token {
	return blockatlas.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Contract,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     blockatlas.GetEthereumTokenTypeByIndex(coinIndex),
	}
}
