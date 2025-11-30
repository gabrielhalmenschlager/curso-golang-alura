package main

import (
	"fmt"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/fetcher"
	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/processor"
)

func main() {
	start := time.Now()

	priceChannel := make(chan models.PriceDetail, 4)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
