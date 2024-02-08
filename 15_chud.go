package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"time"
)

// case 15: // -- AMFchudA
// Chudnovsky method, based on https://arxiv.org/pdf/1809.00533.pdf
/*
The Chudnovsky algorithm is an incredibly fast algorithm for calculating the digits of pi. It was developed by Gregory Chudnovsky and his
brother David Chudnovsky in the 1980s. It is more efficient than other algorithms and is based on the theory of modular equations. It has
been used to calculate pi to over 62 trillion digits.
*/
//  Using this procedure, calculating 1,000,000 digits requires 70516 loops, per the run on:
//  Sun May  7 08:50:23 2023
//  Total run was 8h4m39.7847064s
// AND, THAT CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!

func chud(selection int) { // case 15:
	usingBigFloats = true
	var digits int
	var loops int
	start := time.Now() // saved start time to be compared with end time t
	fmt.Println("\nRichard invites you to enter the number of digits of pi to calculate per the Chudnovsky method: \n")
	fmt.Println("The sky is the limit with this method, so don't be shy.  \n")

	// /* Unix variant
	fmt.Scanf("%d", &digits)
	// Unix variant */
	//
	/* Windows variant
	    fmt.Scanf("%d", &digits)
	    fmt.Scanf("%d", &digits)
	Windows variant */

	pi := new(big.Float)
	loops, pi = CalcPi(float64(digits), start, loops, selection) // returns i which gets a new name "loops"
	//          CalcPi is responsible for most of its printing and logging

	if loops < 100 {
		fmt.Printf("\na peek at the prospective value of Pi as a big float, and formatted 0.122f, is : \n%0.122f \n\n", pi)

		fmt.Println("... but : ") // but, also deploy printResultStatsLong()
		printResultStatsLong(pi, digits, "ChudDidLessThanOneHundredLoops", 1, "", selection)
	}

	// always do the following elapsed time and conditional print work at the end of this selection (case 15:)
	// none of the following will be reached if we exit out of CalcPi with an os.Exit(), so we break instead
	// for now we take this out :
	// fmt.Printf("%1.[1]*[2]f \n", digits, pi)
	fmt.Printf("\n loops were %d, and digits requested was %d \n", loops, digits)
	// fmt.Printf("\nprecR is: %d, and precision is, as yet, undefined near the top of this case.\n", precR)
	// os.Exit(1) // was testing
	t := time.Now()
	elapsed := t.Sub(start)
	// only if ===== all this print work is conditional on some time having elapsed ==================================================
	if int(elapsed.Seconds()) != 0 {
		fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.

		Hostname, _ := os.Hostname()
		_, err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
		check(err0)
		current_time := time.Now()
		_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
		check(err6)
		// the whole pi would be printed to the datalog file on the line below
		// _ , err8 := fmt.Fprintf(fileHandle, "pi was %1.[1]*[2]f \n", digits, pi)
		//    check(err8)
		// ... after printing the whole pi, some nice stats are appended to the file's log entry
		TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
		_, err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %d \n ", TotalRun, digits)
		check(err7)
	}
}

// calculate Pi for n number of digits
func CalcPi(digits float64, start time.Time, loops int, selection int) (int, *big.Float) {
	usingBigFloats = true
	var i int
	/**
	 *   This is an implementation for https://en.wikipedia.org/wiki/Chudnovsky_algorithm
	 *   "It can be improved using binary splitting http://numbers.computation.free.fr/Constants/Algorithms/splitting.html
	 *   if we split it into two independent parts and simplify the formula for more details:
	 * https://www.craig-wood.com/nick/articles/pi-chudnovsky/""
	 */

	// n, apparently, is the expected number of loops we may need to produce digits number of digits
	n := int64(2 + int(float64(digits)/14.181647462))
	// n := int64(2 + int(float64(digits)/12))  // I tried this, and may try something like it again someday?? like /14.0 ?
	// prec := uint(int(math.Ceil(math.Log2(10)*digits)) + int(math.Ceil(math.Log10(digits))) + 2) // the original
	// prec := uint(digits) // not good, not large enough, so ...
	digitsPlus := digits + digits*0.10 // because we needed a little more than the orriginal programmer had figured on :)
	prec := uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)

	c := new(big.Float).Mul(
		big.NewFloat(float64(426880)),
		new(big.Float).SetPrec(prec).Sqrt(big.NewFloat(float64(10005))),
	)

	k := big.NewInt(int64(6))
	k12 := big.NewInt(int64(12))
	l := big.NewFloat(float64(13591409))
	lc := big.NewFloat(float64(545140134))
	x := big.NewFloat(float64(1))
	xc := big.NewFloat(float64(-262537412640768000))
	m := big.NewFloat(float64(1))
	sum := big.NewFloat(float64(13591409))

	pi := big.NewFloat(0)

	x.SetPrec(prec)
	m.SetPrec(prec)
	sum.SetPrec(prec)
	pi.SetPrec(prec)

	bigI := big.NewInt(0)
	bigOne := big.NewInt(1)

	// this is a flag; if it is set to zero we exit
	queryIfTimeToDie := 1
	i = 1 // a secondary dedicated loop counter

	if n > 8998 {
		fmt.Println("\n Well, this is going to take a while, because you asked for too much pie (> 8990)\n")
	}

	for ; n > 0; n-- {
		i++

		// L calculation
		l.Add(l, lc)

		// X calculation
		x.Mul(x, xc)

		// M calculation
		kpower3 := big.NewInt(0)
		kpower3.Exp(k, big.NewInt(3), nil)
		ktimes16 := new(big.Int).Mul(k, big.NewInt(16))
		mtop := big.NewFloat(0).SetPrec(prec)
		mtop.SetInt(new(big.Int).Sub(kpower3, ktimes16))
		mbot := big.NewFloat(0).SetPrec(prec)
		mbot.SetInt(new(big.Int).Exp(new(big.Int).Add(bigI, bigOne), big.NewInt(3), nil))
		mtmp := big.NewFloat(0).SetPrec(prec)
		mtmp.Quo(mtop, mbot)
		m.Mul(m, mtmp)

		// Sum calculation
		t := big.NewFloat(0).SetPrec(prec)
		t.Mul(m, l)
		t.Quo(t, x)
		sum.Add(sum, t)

		// Pi calculation
		pi.Quo(c, sum)
		k.Add(k, k12)
		bigI.Add(bigI, bigOne)

		useAlternateFile := "no" // no means to use the standard log file rather than some special one
		// the compiler is not happy unless it sees this created outside of an if
		// But, wait. Why is the compiler allowing me to violate the no new var left of the := assignment ??? This IS in a loop !!!!
		if i == 100 {
			// useAlternateFile := "no" // the compiler is not happy unless it sees this created outside of an if
			fmt.Printf("\n we are at %d loops, here comes a 800f of pi as a big float: \n", i)
			fmt.Printf("%0.[1]*[2]f \n", 800, pi)
			// fmt.Printf("%0.[1]*[2]f \n", int(digits), pi) // if digits was known to be the verified string of digits, then this is what we would want
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 200 {
			// useAlternateFile = "no" // still no
			fmt.Printf("\n we are at %d loops, here comes a 999f of pi as a big float: \n", i)
			fmt.Printf("%.999f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 400 {
			// useAlternateFile = "no" // still no
			fmt.Printf("\n we are at %d loops, here comes a 1599f of pi as a big float: \n", i)
			fmt.Printf("%.1599f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		//
		//
		// note below the: useAlternateFile = "chudDid800orMoreLoops"
		if i == 800 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 1999f of pi as a big float: \n", i)
			fmt.Printf("%.1999f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 1600 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 2200f of pi as a big float: \n", i)
			fmt.Printf("%.2200f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 2000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 2300F of pi as a big float: \n", i)
			fmt.Printf("%.2300F", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 2400 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 2600F of pi as a big float: \n", i)
			fmt.Printf("%.2600f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 2800 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 2900F of pi as a big float: \n", i)
			fmt.Printf("%.2900f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 3200 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 3100F of pi as a big float: \n", i)
			fmt.Printf("%.3100f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 4000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 3300F of pi as a big float: \n", i)
			fmt.Printf("%.3300f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 6000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 3400F of pi as a big float: \n", i)
			fmt.Printf("%.3400f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if i == 8000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fmt.Printf("\n we are at %d loops, here comes a 3500F of pi as a big float: \n", i)
			fmt.Printf("%.3500f", pi)
			finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
		}
		if queryIfTimeToDie == 0 {
			fmt.Println("if queryIfTimeToDie is 0, time to die")
			fmt.Printf("\nprecision was: %d \n", prec)
			break
		}
		// 1,000,000 digits requires 70516 loops, per the run on May 7 2023 at 10:30
		//  was run on: Sun May  7 08:50:23 2023
		//  Total run was 8h4m39.7847064s
		// AND THE CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!
	} // end of for loop way up thar :: it prompts periodically to continue or die

	// we are out of the loop, so we do the following just once:

	fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1prslc2c)                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
	defer fileHandleBig.Close()                                                                                    // It’s idiomatic to defer a Close immediately after opening a file.

	_, err9bigpie := fmt.Fprint(fileHandleBig, pi) // dump this big-assed pie to a special log file
	check(err9bigpie)

	_, errGoesHere := fmt.Fprint(fileHandleBig, "\n\n")
	check(errGoesHere)

	fileHandleBig.Close()

	return i, pi // asigning i to loops
}

// a helper func
func finishChudIfs(pi *big.Float, digits float64, selection int, i int, useAlternateFile string, queryIfTimeToDie int) {
	printResultStatsLong(pi, int(digits), useAlternateFile, 1, "", selection)
	fmt.Printf("\n the above was from %d loops \n", i)
	fmt.Println("enter any int to continue, 0 to end")
	// /* Unix variant
	fmt.Scanf("%d", &queryIfTimeToDie)
	// Unix variant */
	//
	/* Windows variant
	    fmt.Scanf("%d", &queryIfTimeToDie)
	    fmt.Scanf("%d", &queryIfTimeToDie) // Windows environments seem to just fly past a single Scanf, so I have used two :)
	Windows variant */
	// most of the above was written by Richard Woolley
} // end of chud() set // -- AMFchudB
