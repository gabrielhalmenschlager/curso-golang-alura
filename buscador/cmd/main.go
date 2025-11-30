package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/fetcher"
	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/processor"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var showWg sync.WaitGroup
	showWg.Add(1)

	go func() {
		defer showWg.Done()
		processor.ShowPriceAVG(priceChannel)
	}()

	go fetcher.FetchPrices(priceChannel)

	showWg.Wait()

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
