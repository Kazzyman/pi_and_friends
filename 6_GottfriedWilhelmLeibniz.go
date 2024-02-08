package main

import (
	"fmt"
	"os"
	"time"
)

func GottfriedWilhelmLeibniz(selection int) { // case 6: -- AMFGottfriedWilhelmLeibnizA
	usingBigFloats = false
	fmt.Println("\n\nYou selected #", selection, " Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
	fmt.Println("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton")
	fmt.Println("    ... and Gottfried Wilhelm Leibniz")
	fmt.Println("   4 Billion iterations will be executed ... ")
	fmt.Println("")
	fmt.Println(" ... working ...\n")
	start := time.Now()
	iterFloat64 = 0
	var denom float64
	denom = 3
	var sum float64
	sum = 1 - (1 / denom)
	iterInt64 = 1
	for iterInt64 < 4000000000 {
		iterFloat64++
		iterInt64++
		denom = denom + 2
		if iterInt64%2 == 0 {
			sum = sum + 1/denom
		} else {
			sum = sum - 1/denom
		}
		π = 4 * sum
		if iterInt64 == 100000000 {
			fmt.Println("... 100,000,000 completed iterations ...")
			fmt.Println("\n   #1 2345678#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.1415926,53589793 is from the web")
			fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n\n")
		}
		if iterInt64 == 200000000 {
			fmt.Println("... 200,000,000 gets another digit ...")
			fmt.Println("\n   #1 23456789#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.14159265,3589793 is from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
		}
		if iterInt64 == 400000000 {
			fmt.Println("... 400,000,000 iterations completed, still at nine ...")
			fmt.Println("\n   #1 23456789#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.14159265,3589793 is from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  400,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
		}
		if iterInt64 == 600000000 {
			fmt.Println("... 600,000,000 iterations, still at nine ...")
			fmt.Println("\n   #1 23456789#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.14159265,3589793 is from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  600,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
		}
		if iterInt64 == 1000000000 {
			fmt.Println("... 1 Billion iterations completed, still nine ...")
			fmt.Println("\n   #1 23456789#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.14159265,3589793 is from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  1,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
		}
		if iterInt64 == 2000000000 {
			fmt.Println("... 2 Billion, and still just nine ...")
			fmt.Println("\n   #1 23456789#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.14159265,3589793 is from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  2,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
		}
		if iterInt64 == 4000000000 { // last one
			fmt.Println("\n... 4 Billion, gets us ten digits  ...")
			fmt.Println("\n   #1 234567890#")
			fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
			fmt.Println("    3.141592653,589793 is from the web")
			fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  4,000,000,000 iterations in ", elapsed, " yields 10 digits of π\n\n")
			fmt.Println(" per option #", selection, "  --  the Gottfried Wilhelm Leibniz formula\n")

			LinesPerIter = 14
			fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)
			fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

			// store reults in a log file which can be displayed from within the program by selecting option #12
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz -- selection #%d on %s \n", selection, Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
			check(err2)
			_, err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
			check(err4)
			_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
			check(err5)
			TotalRun := elapsed.String()                                            // cast time duration to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun) // add total runtime of this calculation
			check(err7)
			fmt.Println("Select 12 at menu to display prior results")
		}
	} // end of first for loop

	fmt.Println(colorGreen, "Enter any positive digit to continue with an additional 9 billion iterations, 0 to exit", colorReset)
	option9 := 0
	fmt.Scan(&option9)

	if option9 > 0 {

		fmt.Println(colorCyan, "\n\nYou elected to continue the Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
		fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...", colorReset)

		fmt.Println("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton")
		fmt.Println("    ... and Gottfried Wilhelm Leibniz")
		fmt.Println("    9 billion iterations will be executed \n\n", colorCyan, "   ... working ...\n", colorReset)

		start := time.Now()

		/*        iterFloat64 = 0
		          var denom float64
		              denom = 3
		          var sum float64
		          iterInt64 = 1
		*/
		//      sum = 1-(1/denom)

		for iterInt64 < 9000000000 {
			iterFloat64++
			iterInt64++
			denom = denom + 2
			if iterInt64%2 == 0 {
				sum = sum + 1/denom
			} else {
				sum = sum - 1/denom
			}
			π = 4 * sum

			if iterInt64 == 6000000000 {
				fmt.Println("... 6 Billion completed ... \n")
				fmt.Println("   #1 234567890#")
				fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
				fmt.Println("    3.141592653,589793 is from the web")
				fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println("  6,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
				fmt.Println(colorCyan, "  ... working ...\n", colorReset)
			}
			if iterInt64 == 8000000000 {
				fmt.Println("... 8 Billion completed. still ten ...\n")
				fmt.Println("   #1 234567890#")
				fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
				fmt.Println("    3.141592653,589793 is from the web")
				fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Println("  8,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
				fmt.Println(colorCyan, "  ... working ...\n", colorReset)
			}
			if iterInt64 == 9000000000 {
				fmt.Println("   #1 234567890#")
				fmt.Println("   ", π, "was calculated by the Gottfried Wilhelm Leibniz formula")
				fmt.Println("    3.141592653,589793 is from the web")
				fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
				// fmt.Print("   ", iter)
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Print("\n... 9B iterations in ", elapsed, " , but to get 10 digits we only needed 4B iterations\n\n")
				fmt.Println(" per option #", selection, "  --  the Gottfried Wilhelm Leibniz formula\n")

				t = time.Now()
				elapsed = t.Sub(start)

				LinesPerIter = 14
				fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
				LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
				fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)
				fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

				// store reults in a log file which can be displayed from within the program by selecting option #12
				fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
				defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
				Hostname, _ := os.Hostname()
				_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz (cont.) -- selection #%d on %s \n", selection, Hostname)
				check(err0)
				current_time := time.Now()
				_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
				check(err6)
				_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
				check(err2)
				_, err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
				check(err4)
				_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
				check(err5)
				TotalRun := elapsed.String()                                            // cast time duration to a String type for Fprintf "formatted print"
				_, err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun) // add total runtime of this calculation
				check(err7)
				fmt.Println("Select 12 at menu to display prior results")
			}
		} // end of second for
	} // end of option9 if
	// written entirely by Richard Woolley
} // end GottfriedWilhelmLeibniz() // case 6: // -- AMFGottfriedWilhelmLeibnizB
