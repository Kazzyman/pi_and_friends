package main

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
)

// case 2: // -- AMFBBPFA
func BBPF(selection int) {
	usingBigFloats = true
	// Richard's modifications start: -------------------------------------
	useAlternateFile := "BBPF"

	numberOfDigits := 1 // just to create it
	fmt.Println("\nYou selected #", selection, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
	fmt.Println("This version uses channels: GOMAXPROCS(numCPU), and big floats; how many digits of π would you like?")

	fmt.Scan(&numberOfDigits)

	if numberOfDigits > 25000 {
		fmt.Printf("\nYou requested %d digits of pi. Which is kinda a lot. 25,000 would have taken a minute. \n", numberOfDigits)
		fmt.Println("The amount you asked for could take much longer. Even though I am hammering all of your cores.\n")
	}
	// Richard's modifications end: ---------------------------------------

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	p := uint((int(math.Log2(10)))*numberOfDigits + 3)

	// more of Richards's mods follow: ------------------------------------------------
	fmt.Println("log2 etc. (based on numberOfDigits) has set p at: ", p)

	// I am adding some precision for cases that go for a whole lot of digits

	additionalAmount := 1.2 // just have to create this outside these ifs down thar

	if numberOfDigits <= 12 {
		additionalAmount = float64(p) + (float64(p) * 0.127)
	} else if numberOfDigits <= 500 {
		additionalAmount = float64(p) + (float64(p) * 0.125)
	} else if numberOfDigits <= 5000 {
		additionalAmount = float64(p) + (float64(p) * 0.120)
	} else if numberOfDigits <= 10000 {
		additionalAmount = float64(p) + (float64(p) * 0.119)
	} else if numberOfDigits > 10000 {
		additionalAmount = float64(p) + (float64(p) * 0.115)
	}

	wholePart := uint(math.Floor(additionalAmount))
	p = wholePart
	if p > 85400 {
		fmt.Printf("\n ... we have adjusted that to: %d ... working ...", p)
	} else {
		fmt.Printf("\n ... we have adjusted that to: %d", p)
	}
	// Richard's mod end: -------------------------------------------------------------

	result := make(chan *big.Float, numberOfDigits)
	worker := workers2(p)

	pi := new(big.Float)

	// Richards modification of SetPrec: --------------------
	pi.SetPrec(p) // this gives a lot more results

	for i := 0; i < numberOfDigits; i++ {
		go worker(i, result)
	}
	for i := 0; i < numberOfDigits; i++ {
		pi.Add(pi, <-result)
	}
	// -----------------------------------------------------

	// Richards's mods follow: --------------------------------------------------------
	fmt.Printf("\n\na peek at pi formatted 250f is: %[1]*.[2]*[3]f \n", 1, 250, pi)
	printResultStatsLong(pi, int(p), useAlternateFile, 1, "", selection) // p is just the precision that was used for the big floats, which will be reported by printResults...
	// Richard's mods end: -------------------------------------------------------------

}
func workers2(p uint) func(id int, result chan *big.Float) {
	B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B2 := new(big.Float).SetPrec(p).SetInt64(2)
	B4 := new(big.Float).SetPrec(p).SetInt64(4)
	B5 := new(big.Float).SetPrec(p).SetInt64(5)
	B6 := new(big.Float).SetPrec(p).SetInt64(6)
	B8 := new(big.Float).SetPrec(p).SetInt64(8)
	B16 := new(big.Float).SetPrec(p).SetInt64(16)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(1)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B16)
		}

		C2 := new(big.Float).SetPrec(p)
		C2.Mul(B8, Bn)

		T1 := new(big.Float).SetPrec(p)
		T1.Add(C2, B1)
		T1.Quo(B4, T1)

		T2 := new(big.Float).SetPrec(p)
		T2.Add(C2, B4)
		T2.Quo(B2, T2)

		T3 := new(big.Float).SetPrec(p)
		T3.Add(C2, B5)
		T3.Quo(B1, T3)

		T4 := new(big.Float).SetPrec(p)
		T4.Add(C2, B6)
		T4.Quo(B1, T4)

		R := new(big.Float).SetPrec(p)
		R.Sub(T1, T2)
		R.Sub(R, T3)
		R.Sub(R, T4)
		R.Quo(R, C1)

		result <- R
	}
	// Richard significantly modified a version he copied from ... somewhere
} // end of BBPF() set // -- AMFBBPFB
