package main

import (
	"fmt"
	"os"
	"time"
)

// case 44: // -- AMFLeibniz_method_one_billion_itersA
func Leibniz_method_one_billion_iters(selection int) {
	usingBigFloats = false
	start := time.Now() // mine, option 44 from second menu
	fmt.Println("\n you selected option 44, which is going to run for 'a minute' \n")
	/**
	 *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com)
	 *
	 *   Approximation of PI using Leibniz method.
	 *   1-1/3+1/5-1/7
	 *
	 *   Compile and Run -> go run leibniz_pi.go
	 *
	 *   License: MIT
	 */

	// ******* RUNS TEN BIL ITERS SO IT RUNS A LONG TIME ***************************************

	// var t1, t2, r, sign, i float64
	var t1, t2, sign, i float64

	t1 = 1.0

	t2 = 1.0 - float64(1.0/3.0)

	// r = t1 - t2;

	i = 5

	// eps := 0.00001

	sign = 1

	iters := 0

	// for float64(4 * r) >= eps {
	for iters < 10000000000 {

		t1 = t2

		t2 += 1.0 / i * sign

		if t1 > t2 {

			// r = float64(t1 - t2)

		} else {

			// r = float64(t2 - t1)
		}

		sign = (-1) * sign

		i += 2

		iters++
	}
	// fmt.Println("%x", 4 * t2)

	// Rick's code follows:
	pi := 4 * t2
	fmt.Println("pi is:\n", pi)
	fmt.Println(" 3.141592653589793 <-- actual")
	fmt.Println(" 1 234567890123 so, 10 digits were Calculated correctly after 10,000,000,000 iterations\n")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Ten Billion iterations were completed in ", elapsed, " yielding 10 digits of π\n")

	// log stats to a log file
	var LinesPerSecondInt int
	LinesPerIter := 6                                                                                                       // an estimate
	LinesPerSecondInt = (LinesPerIter * iters) / int(elapsed.Seconds())                                                     // .Seconds() returns a float64
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
	defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
	Hostname, _ := os.Hostname()
	_, err0 := fmt.Fprintf(fileHandle, "\n  -- Approximation of PI using Leibniz method -- selection #%d on %s \n", selection, Hostname)
	check(err0)
	current_time := time.Now()
	_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err6)
	_, err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt)
	check(err2)
	_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters)/elapsed.Seconds())
	check(err4)
	_, err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters)
	check(err5)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
	check(err7)
	// adapted by Richard Woolley
} // end of Leibniz_method_one_billion_iters() // -- AMFLeibniz_method_one_billion_itersB
