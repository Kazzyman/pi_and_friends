package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// case 50: // -- AMFcompareFloat64withBigFloatsA
// comparison of float64 vs big.Float types using a Nilakantha Somayaji example
// π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...
// I just love github.com, Sublime Text, and Go.lang

func compareFloat64withBigFloats() { // case 50:
	usingBigFloats = true
	fmt.Println("\n ... just a moment, working ... \n")
	start := time.Now()
	with_float64_types()                         // runs pretty quickly through tens of billions of iters
	precision, iterBig := with_big_Float_types() // can take forever with large values of prec and iters
	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	fmt.Printf("\n\nTotal run with SetPrec at: %d and iters of %d was %s \n ", precision, iterBig, TotalRun)
}

func with_float64_types() {
	usingBigFloats = false
	// float64 constants:

	var four64 float64
	four64 = 4

	var three64 float64
	three64 = 3

	// float64 variables:

	var digitone64 float64
	digitone64 = 2

	var digittwo64 float64
	digittwo64 = 3

	var digitthree64 float64
	digitthree64 = 4

	var sum64 float64
	sum64 = three64 + (four64 / (digitone64 * digittwo64 * digitthree64))

	var nextterm64 float64

	iters_fl64 := 1
	for iters_fl64 < 10000000000 { // 10,000,000,000 ten billion
		iters_fl64++
		digitone64 = digitone64 + 2
		digittwo64 = digittwo64 + 2
		digitthree64 = digitthree64 + 2
		nextterm64 = four64 / (digitone64 * digittwo64 * digitthree64)

		if iters_fl64%2 == 0 {
			sum64 = sum64 - nextterm64
		} else {
			sum64 = sum64 + nextterm64
		}
	}
	printResultStats64comp(sum64, iters_fl64)
}

func printResultStats64comp(sum64 float64, iterBig int) {
	stringOfSum := strconv.FormatFloat(sum64, 'f', 160, 64)
	piAs100chars := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706"

	fmt.Printf("\npi as calculated per float64s is: %0.76f \n", sum64)
	fmt.Println("                                                ^ ")
	fmt.Println("                                  1 234567890123456789012345678901234567890123456789012345")
	fmt.Println("                                           10   |    20        30        40        50     ")
	fmt.Printf("pi from the web is:               %s \n", piAs100chars)

	posInPi := 0               // to be the incremented offset : piChar = piAs86chars[posInPi]
	var piChar byte            // one byte (character) of pi as string, e.g. piChar = piAs86chars[posInPi]
	var copyOfLastPosition int // an external (to the loop) copy of positionInString
	var stringVerOfCorrectDigits = []string{}
	for positionInString, charAtRangePos := range stringOfSum {
		piChar = piAs100chars[posInPi]
		if charAtRangePos == rune(piChar) {
			stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
			// fmt.Printf("we have agreement up to character pos: %d \n", positionInString)
			copyOfLastPosition = positionInString // save an external copy, of the last position found to have matched pi, as an int
		} else {
			break // to print result and info below
		}
		posInPi++
	}
	fmt.Printf("\nWe have calculated pi correctly to %d digits using %d iters and float64s types \n", copyOfLastPosition, iterBig)
	fmt.Print(" .... that is ten Billion iterations!!! \n\n")
	fmt.Print("The correctly calculated digits are: ")
	for _, oneChar := range stringVerOfCorrectDigits {
		fmt.Print(oneChar)
	}
	fmt.Print("\n")
}

func with_big_Float_types() (precision uint, iterBig int) {
	usingBigFloats = true
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

	precision = 164
	sumBig.SetPrec(precision)
	twoBig.SetPrec(precision)
	threeBig.SetPrec(precision)
	fourBig.SetPrec(precision)
	digitoneBig.SetPrec(precision)
	digittwoBig.SetPrec(precision)
	digitthreeBig.SetPrec(precision)
	nexttermBig.SetPrec(precision)

	sumBig.Add(threeBig, new(big.Float).Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig))))

	iterBig = 1
	for iterBig < 960000 { // 1,000,000,000
		// Total run with SetPrec at:  128 and iters of 50000000  was 23.1879034s    :: 3.14159265358979323846264
		// Total run with SetPrec at: 1024 and iters of 600000000 was 12m23.9554367s :: 3.14159265358979323846264338
		/* We have calculated pi correctly to 28 digits using 1000000000 iters and Prec of 128
		           The correctly calculated digits are:                                        3.141592653589793238462643383
		           Total run with SetPrec at: 128 and iters of 1000000000 was 7m57.3179415s
		   // got 31 digits in 1 hour and 26 min using this algorithm with one billion (must have been ten) iters at 128 prec
		*/
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
	}
	printResultStatsLongComp(sumBig, iterBig, int(precision))
	return precision, iterBig // still is an int ^^^^^^^^^
}

func printResultStatsLongComp(sumBig *big.Float, iterBig int, precision int) {
	piAs100chars := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706"
	stringOfSum := sumBig.Text('f', 160) // create a string version of a big result, 700 some odd chars in length

	fmt.Printf("\n\npi as calculated per big.Floats is: %0.76f \n", sumBig)
	fmt.Println("                                                   ^ ")
	fmt.Println("                                    1 234567890123456789012345678901234567890123456789012345")
	fmt.Println("                                             10    |   20        30        40        50     ")
	fmt.Printf("pi from the web is:                 %s \n", piAs100chars)

	posInPi := 0               // to be the incremented offset : piChar = piAs86chars[posInPi]
	var piChar byte            // one byte (character) of pi as string, e.g. piChar = piAs86chars[posInPi]
	var copyOfLastPosition int // an external (to the loop) copy of positionInString
	var stringVerOfCorrectDigits = []string{}
	for positionInString, charAtRangePos := range stringOfSum {
		piChar = piAs100chars[posInPi]
		if charAtRangePos == rune(piChar) {
			stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
			// fmt.Printf("we have agreement up to character pos: %d \n", positionInString)
			copyOfLastPosition = positionInString // save an external copy, of the last position found to have matched pi, as an int
		} else {
			break // to print result and info below
		}
		posInPi++
	}
	fmt.Printf("\nWe have calculated pi correctly to %d digits using %d iters and Prec of %d on big.Floats \n", copyOfLastPosition, iterBig, precision)
	fmt.Print(" ... and here it is only 960,000 iterations !!!\n\n")
	fmt.Print("The correctly calculated digits are: ")
	for _, oneChar := range stringVerOfCorrectDigits {
		fmt.Print(oneChar)
	}
	fmt.Print("\n")
	// written entirely by Richard Woolley
} // end of compareFloat64withBigFloats() set // -- AMFcompareFloat64withBigFloatsB
