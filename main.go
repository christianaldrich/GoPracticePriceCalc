package main

import (
	"fmt"

	filemanager "calc.com/price-calc/fileManager"
	prices "calc.com/price-calc/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index]) //run this Process in parallel by using goroutine

		// if err != nil {
		// 	fmt.Println("Couldn't process job")
		// 	fmt.Println("error.")
		// }
	}
	// fmt.Println(result)

	for index := range taxRates {
		select { // tunggu 1 case yang kasih value, kalau 1 udh masuk yg lain didiscard
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("DONE.")
		}
	}

}
