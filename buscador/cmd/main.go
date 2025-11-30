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
	var wg, showWg sync.WaitGroup
	wg.Add(3)
	showWg.Add(1)

	go func() {
		defer showWg.Done()
		processor.ShowPriceAVG(priceChannel)
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite3()
	}()

	wg.Wait()
	close(priceChannel)
	showWg.Wait()

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
