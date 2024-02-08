package main

import (
	"fmt"
	"math/big"
	"time"
)

// case 5: // -- AMFNilakantha_Somayaji_with_big_Float_typesA
func Nilakantha_Somayaji_with_big_Float_types(selection int) {

	usingBigFloats = true

	fmt.Println("\nYou have selected the Nilakantha Somayaji method using big.Float types, and with some ")
	fmt.Print("patience one can generate 31 correct digits of pi using it.\n\n")

	fmt.Println("Enter the number of iterations (suggest between 100,000 and 100,000,000)")

	var iters int

	// /* Unix variant
	fmt.Scanf("%d", &iters)
	// Unix variant */
	//
	/* Windows variant
	    fmt.Scanf("%d", &iters)
	    fmt.Scanf("%d", &iters)
	Windows variant */

	fmt.Println("Enter the precision: (suggest between 128 and 512)")

	var precision int

	// /* Unix variant
	fmt.Scanf("%d", &precision)
	// Unix variant */
	//
	/* Windows variant
	    fmt.Scanf("%d", &precision)
	    fmt.Scanf("%d", &precision)
	Windows variant */

	if iters > 36111222 {
		fmt.Println(" ... ")
	}
	if iters > 42000000 {
		fmt.Println("... werkin ...")
	}
	if iters > 55111222 {
		fmt.Println("... working for a while ...")
	}
	if iters > 69111222 {
		fmt.Println("... will be working for quite a while ...")
	}
	if iters > 80111222 {
		fmt.Println("... a very long while ... working ...")
	}

	start := time.Now()
	var iterBig int

	// big.Float constants:

	twoBig := big.NewFloat(2)
	threeBig := big.NewFloat(3)
	fourBig := big.NewFloat(4)

	// big.Float variables:

	digitoneBig := new(big.Float)
	*digitoneBig = *twoBig

	digittwoBig := new(big.Float)
	*digittwoBig = *threeBig

	digitthreeBig := new(big.Float)
	*digitthreeBig = *fourBig

	sumBig := new(big.Float)
	nexttermBig := new(big.Float)

	// set precision to a user-specified value
	sumBig.SetPrec(uint(precision))
	twoBig.SetPrec(uint(precision))
	threeBig.SetPrec(uint(precision))
	fourBig.SetPrec(uint(precision))
	digitoneBig.SetPrec(uint(precision))
	digittwoBig.SetPrec(uint(precision))
	digitthreeBig.SetPrec(uint(precision))
	nexttermBig.SetPrec(uint(precision))

	sumBig.Add(threeBig, new(big.Float).Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig))))

	iterBig = 1
	for iterBig < iters { // 1,000,000,000 yeilds 19 digits in 13 sec

		// 256p 100,000,000 : 56s 25 digits digits June 26 2023
		// 1,280p and 1Bil : 23 min without ending
		// 128p and 100,000,000 49s gave 25 digits June 26 2023
		// Total run with SetPrec at: 128 and iters of 1,000,000,000 was 7m57.3179415s
		// got 31 digits in 1 hour and 26 min using this algorithm with one billion iters at 128 prec
		// 1,000,000,002 and 64 bits prec yielded :: 17 digits in 5m41s

		iterBig++

		digitoneBig.Add(digitoneBig, twoBig)
		digittwoBig.Add(digittwoBig, twoBig)
		digitthreeBig.Add(digitthreeBig, twoBig)

		nexttermBig.Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig)))

		if iterBig%2 == 0 { // % is modulus operator
			sumBig.Sub(sumBig, nexttermBig)
		} else {
			sumBig.Add(sumBig, nexttermBig)
		}

		if iterBig == 20111222 {
			fmt.Println(" ... doin some ... ")
		}
		if iterBig == 36111222 {
			fmt.Println(" ... werkin ... ")
		}
		if iterBig == 42000000 {
			fmt.Println("... still werkin ...")
		}
		if iterBig == 55111222 {
			fmt.Println("... been working for a while ...")
		}
		if iterBig == 69111222 {
			fmt.Println("... been working for quite a while ...")
		}
		if iterBig == 80111222 {
			fmt.Println("... it's been a very long while ... but still working ...")
		}
		if iterBig == 180111222 {
			fmt.Println("... it's been a very long while, 180,111,222 down, ... and still working ...")
		}
		if iterBig == 280111222 {
			fmt.Println("... it's been a very long while, 280,111,222 down, ... and still working ...")
		}
		if iterBig == 480111222 {
			fmt.Println("... it's been a very long while, 480,111,222 down, ... still working ...")
		}
		if iterBig == 680111222 {
			fmt.Println("... it's been a very long while, 680,111,222 down, ...  working ...")
		}
		if iterBig == 880111222 {
			fmt.Println("... it's been a very long while, down, 880,111,222, down ... still, working ...")
		}
		if iterBig == 977111222 {
			fmt.Println("... it's been a very long while, 977,111,222 already ... why am I still working? ...")
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String()
	// print to file:
	printResultStatsLong(sumBig, precision, "Nilakantha", iterBig, TotalRun, selection)
	// print to screen:
	fmt.Println(" via Nilakantha \n")
	fmt.Printf("Total run with SetPrec at: %d and iters of %d was %s \n\n ", precision, iterBig, TotalRun)
	// written entirely by Richard Woolley
} // end of Nilakantha_Somayaji_with_big_Float_types() // -- AMFNilakantha_Somayaji_with_big_Float_typesB
