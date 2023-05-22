package goethclient

import (
	"math"
	"math/big"
)

func WeiToString(wei *big.Int) string {
	w := new(big.Float).SetInt(wei)
	w = new(big.Float).Quo(w, big.NewFloat(math.Pow10(EthDecimals)))
	return w.String()
}

func TftUnitsToString(units *big.Int) string {
	w := new(big.Float).SetInt(units)
	w = new(big.Float).Quo(w, big.NewFloat(math.Pow10(TftDecimals)))
	return w.String()
}
