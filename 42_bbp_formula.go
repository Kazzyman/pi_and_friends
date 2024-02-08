package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"time"
)

func bbp_formula(selection int) { // case 42: // -- AMFbbp_formulaA
	usingBigFloats = true
	iters_bbp := 1
	start := time.Now()
	// numCPU := runtime.NumCPU()
	// runtime.GOMAXPROCS(numCPU)

	n := 46
	p := uint((int(math.Log2(10)))*n + 3)

	result := make(chan *big.Float, n)
	worker := workers(p)

	pi := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < n; i++ {
		go worker(i, result)
		iters_bbp = i
	}

	for i := 0; i < n; i++ {
		pi.Add(pi, <-result)
		iters_bbp = i
	}

	dur := time.Since(start)
	fmt.Printf("took %v to calculate %d digits of pi \n", dur, n)
	fmt.Printf("%[1]*.[2]*[3]f \n", 1, n, pi)

	// log run stats to a log file
	t := time.Now()
	elapsed := t.Sub(start)
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
	defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
	Hostname, _ := os.Hostname()
	_, err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using the bbp formula -- selection #%d on %s \n", selection, Hostname)
	check(err0)
	current_time := time.Now()
	_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err6)
	_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters_bbp)/elapsed.Seconds())
	check(err4)
	_, err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters_bbp)
	check(err5)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
	check(err7)
}

func workers(p uint) func(id int, result chan *big.Float) {
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
	// adapted by Richard Woolley
} // end of bbp_formula() set // -- AMFbbp_formulaB
