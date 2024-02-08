package main

import (
	"fmt"
	"os"
	"time"
)

func GregoryLeibniz(selection int) { // case 7: -- AMFGregoryLeibnizA
	usingBigFloats = false
	fmt.Println("\n\nYou selected #", selection, " the Gregory-Leibniz series ...")
	fmt.Println("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
	fmt.Println("Three hundred million iterations are being executed ... working ...\n")
	start := time.Now()
	iterFloat64 = 0
	var nextOdd float64
	nextOdd = 1
	four = 4
	var tally float64
	tally = (four / nextOdd)
	iterInt64 = 0
	for iterInt64 < 300000000 {
		iterInt64++
		iterFloat64++
		nextOdd = nextOdd + 2
		tally = tally - (tally / nextOdd)
		tally = tally + (tally / nextOdd) // pi (tally) is set equl to the sum of a subtraction and an addition, alternatively

		if iterInt64 == 10000000 {
			fmt.Println("... 10,000,000 of three hundred million completed. still working, but ...")
			fmt.Println("\n   #1 234567#")
			fmt.Println("   ", tally, "was calculated by the Gregory-Leibniz series")
			fmt.Println("    3.141592,653589793 is from the web")
			fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  10,000,000 iterations in ", elapsed, " yields 7 digits of π\n\n")
		}
		// 7 digits of Pi were found per the above code
		// the next two ifs give eight digits of Pi
		if iterInt64 == 50000000 {
			fmt.Println("... 50,000,000 of three hundred million completed. still working, but ...")
			fmt.Println("\n   #1 2345678#")
			fmt.Println("   ", tally, "was calculated by the Gregory-Leibniz series")
			fmt.Println("    3.1415926,53589793 is from the web")
			fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  50,000,000 iterations in ", elapsed, " yields 8 digits of π\n")
			fmt.Println(" ")
		}
		if iterInt64 == 100000000 {
			fmt.Println("... 100,000,000 of three hundred million completed. still working, but ...")
			fmt.Println("\n   #1 2345678#")
			fmt.Println("   ", tally, "was calculated by the Gregory-Leibniz series")
			fmt.Println("    3.1415926,53589793 is from the web")
			fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n\n")
		}
		// 9 digits of Pi are found below
		if iterInt64 == 200000000 {
			fmt.Println("... 200,000,000 of three hundred million completed. still working, but ...")
			fmt.Println("\n   #1 23456789#")
			fmt.Println("   ", tally, "was calculated by the Gregory-Leibniz series")
			fmt.Println("    3.14159265,3589793 is from the web")
			fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
		}
		if iterInt64 == 300000000 { // last one, still 9 digits
			fmt.Println("\n   #1 23456789#")
			fmt.Print("    ", tally, " was calculated by the Gregory-Leibniz series")
			fmt.Println("\n    3.141592653589793 is from the web")
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("  300 million iterations still yields 9 digits, ") // no Println here
			fmt.Print("in ", elapsed, "\n\n")
			fmt.Println(" per option #", selection, "  --  the Gregory-Leibniz series, circa 1676\n")

			LinesPerIter = 11 // an estimate of the number of lines per iteration
			fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fmt.Printf("       %.0f lines of code were executed per second \n", LinesPerSecond)
			fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

			// store reults in a log file which can be displayed from within the program by selecting option #12
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- selection #%d on %s \n", selection, Hostname)
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
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
			check(err7)
			fmt.Println("Select 12 at menu to display prior results")
		}
	}
	// written entirely by Richard Woolley
} // end of GregoryLeibniz() // case 7: // -- AMFGregoryLeibnizB
