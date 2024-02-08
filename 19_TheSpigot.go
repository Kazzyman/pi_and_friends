package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// and now for the spigot func itself ... (it's just so tight!)
func TheSpigot(selection int) { // case 19: // -- AMFTheSpigotA
	usingBigFloats = false
	var numberOfDigitsToCalc int
	fmt.Println("How much pi can you handle?")
	fmt.Println("How many digits of pi do you really want? Enter that number now:") // prompt the user
	fmt.Scan(&numberOfDigitsToCalc)
	fmt.Println("")

	Spigot(numberOfDigitsToCalc, selection)

	fmt.Print("\n\nThis trick was made possible by a bit of code I mooched off of GitHub ...\n\n")
	fmt.Println("... to view the code with attribution Enter '39' at either menu\n")
	fmt.Println("To view an explaination of how it works enter 99 at either menu\n")
}

// func Spigot(n int) string { // called by the previous func
func Spigot(n int, selection int) { // Rick's version does not return a string // called by the previous func
	start := time.Now()
	pi := "" // allocate a string var "pi"
	boxes := n * 10 / 3
	remainders := make([]int, boxes)
	for i := 0; i < boxes; i++ {
		remainders[i] = 2
	}
	digitsHeld := 0
	for i := 0; i < n; i++ {
		carriedOver := 0
		sum := 0
		for j := boxes - 1; j >= 0; j-- {
			remainders[j] *= 10
			sum = remainders[j] + carriedOver
			quotient := sum / (j*2 + 1)
			remainders[j] = sum % (j*2 + 1)
			carriedOver = quotient * j
		}
		remainders[0] = sum % 10
		q := sum / 10
		switch q {
		case 9:
			digitsHeld++
		case 10:
			q = 0
			for k := 1; k <= digitsHeld; k++ {
				replaced, _ := strconv.Atoi(pi[i-k : i-k+1])
				if replaced == 9 {
					replaced = 0
				} else {
					replaced++
				}
				pi = delChar(pi, i-k) // delChar func follows
				pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
			}
			digitsHeld = 1
		default:
			digitsHeld = 1
		}
		// Rick's code :
		if i == 1 {
			fmt.Print(".")
		} // insert the decimal point between the 3 and the 1
		fmt.Print(strconv.Itoa(q)) // Rick's new method of displaying pi, one digit at a time
		if i == 5000 || i == 10000 || i == 20000 || i == 40000 || i == 80000 || i == 140000 || i == 200000 {
			t := time.Now()
			elapsed := t.Sub(start)
			// log stats to a log file
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- partial Spigot in process -- selection #%d on %s \n", selection, Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "... running at: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Runtime this far is %s \n", TotalRun)
			check(err7)
			_, err8 := fmt.Fprintf(fileHandle, "... while calculating Pi to %d digits, having completed %d digits\n", n, i)
			check(err8)
		}
		// end Rick's code
		pi += strconv.Itoa(q) // in orriginal method, entire string accumulator was being returned
	}
	// Rick's code :
	t := time.Now()
	elapsed := t.Sub(start)
	// log stats to a log file
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
	defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
	Hostname, _ := os.Hostname()
	_, err0 := fmt.Fprintf(fileHandle, "\n  -- Spigot -- selection #%d on %s \n", selection, Hostname)
	check(err0)
	current_time := time.Now()
	_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err6)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n", TotalRun)
	check(err7)
	_, err8 := fmt.Fprintf(fileHandle, "To calculate Pi to %d digits\n", n)
	check(err8)
	// end Rick's code
	// return pi // prior code
}

func delChar(s string, index int) string {
	tmp := []rune(s)
	return string(append(tmp[0:index], tmp[index+1:]...))
	// written largely by Richard Woolley
} // end TheSpigot() set // -- AMFTheSpigotB
