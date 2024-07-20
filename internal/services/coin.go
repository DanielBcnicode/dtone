package services

import (
	"fmt"
	"math"
	"strconv"
)

func CoinStringToInt64(coin string) (int64, error) {
	fa, err := strconv.ParseFloat(coin, 64)
	if err != nil {
		return 0, err
	}
	ia, _ := math.Modf(fa * 100.00)
	return int64(ia), nil
}

func CoinInt64ToString(coin int64) string {
	b := float64(coin) / 100.00
	return fmt.Sprintf("%.2f", b)
}
