package main

import (
	"fmt"
	"os"
	"time"
)

func JohnWallis(selection int) { // case 8: // -- AMFJohnWallisA
	usingBigFloats = false
	fmt.Println("\n   You selected #", selection, " A Go language exercize which can be used to test the speed of your hardware.")
	fmt.Println("   We will calculate π to a maximum of ten digits of accuracy using an infinite series by John Wallis circa 1655")
	fmt.Println("   Up to 40 Billion iterations of the following formula will be executed ")
	fmt.Println("   π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...")
	start := time.Now()
	iterFloat64 = 0
	var numerators float64
	numerators = 2
	var firstDenom float64
	firstDenom = 1
	var secondDenom float64
	secondDenom = 3
	var cumulativeProduct float64
	cumulativeProduct = (numerators / firstDenom) * (numerators / secondDenom)
	iterInt64 = 0
	for iterInt64 < 1000000000 {
		iterInt64++
		iterFloat64++
		numerators = numerators + 2
		firstDenom = firstDenom + 2
		secondDenom = secondDenom + 2
		cumulativeProduct = cumulativeProduct * (numerators / firstDenom) * (numerators / secondDenom)
		π = cumulativeProduct * 2
		/*            if iterInt64 == 100 {
		                  fmt.Println("   #1 2# ")
		                  fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
		                  fmt.Println("    3.1,41592653589793 is, again, the value of π from the web")
		                  fmt.Println("   #1 2# 34567890123456 :: counting the first 16 actual digits of π")
		                  t := time.Now()
		                  elapsed := t.Sub(start)
		                  fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n")
		              }
		              if iterInt64 == 500 {
		                  fmt.Println("   #1 23# ")
		                  fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
		                  fmt.Println("    3.14,1592653589793 is, again, the value of π from the web")
		                  fmt.Println("   #1 23# 4567890123456 :: counting the first 16 actual digits of π")
		                  t := time.Now()
		                  elapsed := t.Sub(start)
		                  fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 3 digits of π\n")
		              }
		*/
		if iterInt64 == 2000 {
			fmt.Println("\n   #1 234# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.14,1592653589793 is, again, the value of π from the web")
			fmt.Println("   #1 23 4567890123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 4 digits of π\n")
		}
		if iterInt64 == 10000 {
			fmt.Println("   #1 2345# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.1415,92653589793 is, again, the value of π from the web")
			fmt.Println("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("10,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n")
		}
		if iterInt64 == 50000 { // 50,000
			fmt.Println("   #1 2345# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.1415,92653589793 is, again, the value of π from the web")
			fmt.Println("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("50,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n")
		}
		if iterInt64 == 500000 { // 500,000 done
			fmt.Println("   #1 23456# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.14159,2653589793 is, again, the value of π from the web")
			fmt.Println("   #1 23456 7890123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("500,000 iterations were completed in ", elapsed, " yielding 6 digits of π\n")
		}
		if iterInt64 == 2000000 { // 2M done
			fmt.Println("   #1 234567# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.141592,653589793 is, again, the value of π from the web")
			fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("2,000,000 iterations were completed in ", elapsed, " yielding 7 digits of π\n")
		}
		if iterInt64 == 40000000 { // 40M done
			fmt.Println("   #1 2345678# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.1415926,53589793 is, again, the value of π from the web")
			fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("40,000,000 iterations were completed in ", elapsed, " yielding 8 digits of π\n\n")
			fmt.Println("  .. working .. on another factor-of-ten iterations\n")
		}
		if iterInt64 == 400000000 { // 400M done
			fmt.Println("   #1 23456789# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.14159265,3589793 is, again, the value of π from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("400,000,000 iterations were completed in ", elapsed, " yielding 9 digits of π\n\n")

			LinesPerIter = 36 // an estimate
			fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)
			fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
			// a brief Red notification follows :
			fmt.Println(colorRed, " ... will be working on doing Billions more iterations ...\n\n", colorReset)
		}
		//
		if iterInt64 == 600000000 { // 600M done
			fmt.Println("  600M done, still working on another Two-Hundred-Thousand iterations ... working ...\n")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println(elapsed, "\n")
			fmt.Println("Calculating 10th digit (40B iters) which takes a few min \n")
			fmt.Println("- Ctrl-C to End/Exit without saving results\n")
			LinesPerIter = 36 // an estimate
			fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)
			fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
			fmt.Println(" ... still working ...")
		}
		if iterInt64 == 800000000 { // 800M done
			fmt.Println("  800M done, still working on yet another Two Hundred Thousand iterations ... working ...\n")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println(elapsed, " \n")
		}
		if iterInt64 == 1000000000 { // 1B done
			fmt.Println("   #1 23456789# ")
			fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
			fmt.Println("    3.14159265,3589793 is the value of π from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("\nOne Billion iterations were completed in ", elapsed, " still only yielding π to 9 digits\n")
			fmt.Println(" per option #", selection, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------

			LinesPerIter = 36 // an estimate
			fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)
			fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

			// store reults in a log file which can be displayed from within the program by selecting option #12
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis -- selection #%d on %s \n", selection, Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
			check(err2)
			_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
			check(err4)
			_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
			check(err5)
			TotalRun := elapsed.String()                                         // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) // add total runtime of this calculation
			check(err7)
			fmt.Println("Select 12 at menu to display prior results")
		}
	} // end of first for loop

	fmt.Println(colorGreen, "Enter any positive digit to continue with an additional 39 billion iterations, 0 to exit", colorReset)
	option39 := 0
	fmt.Scan(&option39)

	if option39 > 0 {

		fmt.Println(colorCyan, "\n\nYou elected to continue the infinite series by John Wallis", colorReset)
		fmt.Println("\n    an additionl 39 billion iterations will be executed \n\n", colorCyan, "   ... working ...\n", colorReset)

		fmt.Println(colorRed, " ... still working ... on Billions of iterations, 39 to go ...\n", colorReset)

		fmt.Println("\n ... 39 Billion additional loops now ensue, just to get one additional digit, the tenth digit of pi")

		start := time.Now()

		for iterInt64 < 40000000000 {
			iterInt64++
			iterFloat64++
			numerators = numerators + 2
			firstDenom = firstDenom + 2
			secondDenom = secondDenom + 2
			cumulativeProduct = cumulativeProduct * (numerators / firstDenom) * (numerators / secondDenom)
			π = cumulativeProduct * 2

			if iterInt64 == 2000000000 { // 2B completed
				fmt.Println("  2B done, still working ... on another Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 3000000000 { // 3B completed
				fmt.Println("  3B done, still working ... on another Billion iterations ... working ... Ctrl-C to End/Exit without saving stats")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 4000000000 { // 4B completed
				fmt.Println("  4B done, still working ... on another Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 5000000000 { // 5B completed
				fmt.Println("  5B done, still working ... on another Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 6000000000 { // 6B completed
				fmt.Println("  6B done, still working ... on another Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 7000000000 { // 7B completed
				fmt.Println("  7B done, still working ... on another Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 8000000000 { // 8B completed
				fmt.Println("  8B done, still working ... on another Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 9000000000 { // 9B completed
				fmt.Println("  9B done, still working ... on another five Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 14000000000 { // 14B completed
				fmt.Println("  14B done, still working ... on another five Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 19000000000 { // 19B completed
				fmt.Println("  19B done, still working ... on another five Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 24000000000 { // 24B completed
				fmt.Println("  24B done, still working ... on another five Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 29000000000 { // 29B completed
				fmt.Println("  29B done, still working ... on another five Billion iterations ... working ...")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 34000000000 { // 34B completed
				fmt.Println("  34B done, still working ... just another six Billion iterations to go! ... ")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println(elapsed)
			}
			if iterInt64 == 40000000000 { // 40B completed
				fmt.Println("   #1 234567890# ")
				fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655")
				fmt.Println("    3.141592653,589793 is the value of π from the web")
				fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")

				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println("Forty Billion iterations were completed in ", elapsed, " yielding π to 10 digits\n")
				fmt.Println(" per option #", selection, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------
				LinesPerIter = 36                                                                               // an estimate
				fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
				LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
				fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)
				fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

				// store reults in a log file which can be displayed from within the program by selecting option #12
				fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
				defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
				Hostname, _ := os.Hostname()
				_, err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis (cont.) -- selection #%d on %s \n", selection, Hostname)
				check(err0)
				current_time := time.Now()
				_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
				check(err6)
				_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
				check(err2)
				_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
				check(err4)
				_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
				check(err5)
				TotalRun := elapsed.String()                                         // cast time durations to a String type for Fprintf "formatted print"
				_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) // add total runtime of this calculation
				check(err7)
				fmt.Println("Select 12 at menu to display prior results")
			}
		} // end of second for loop
	} // end of 40B continuation if
	// written entirely by Richard Woolley
} // end of JohnWallis() // case 8: // -- AMFJohnWallisB
