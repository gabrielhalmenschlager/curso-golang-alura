package processor

import "fmt"

func ShowPriceAVG(priceChannel chan float64) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("Preço recebido: R$ %.2f | Preço até agora: R$ %.2f \n", price, avgPrice)
	}
}
