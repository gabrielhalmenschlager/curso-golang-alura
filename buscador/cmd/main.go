package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/fetcher"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for price := range priceChannel {
			fmt.Printf("Preço recebido: R$ %.2f \n", price)
		}
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

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
