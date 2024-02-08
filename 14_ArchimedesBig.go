package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

func ArchimedesBig(selection int) { // case 14: // --AMFArchimedesBigA
	usingBigFloats = true
	fmt.Println("\nYou selected #", selection, "  --  An improved version of Archimedes' method")
	start := time.Now()
	r := big.NewFloat(1)             // radius is a constant 1
	s1 := big.NewFloat(1)            // starts out as 1
	numberOfSides := big.NewFloat(6) // the number of sides of the polygon

	a := new(big.Float)    // height of bisected triangle
	b := new(big.Float)    // new short side
	p := new(big.Float)    // perimeter
	s2 := new(big.Float)   // new hypotenuse / new side
	p_d := new(big.Float)  // pi
	s1_2 := new(big.Float) // s1/2

	// Set the precision to a higher value
	precision := 55000
	p_d.SetPrec(uint(precision))
	a.SetPrec(uint(precision))
	s1_2.SetPrec(uint(precision))
	s2.SetPrec(uint(precision))
	b.SetPrec(uint(precision))
	p.SetPrec(uint(precision))
	r.SetPrec(uint(precision))
	s1.SetPrec(uint(precision))
	numberOfSides.SetPrec(uint(precision))

	// Calculate initial value of p and set p_d to the same value
	numberOfSides.Mul(numberOfSides, big.NewFloat(2)) // doubleing the number of sides was not orriginally done here ????????????????

	s1_2.Quo(s1, big.NewFloat(2))                                 // s1_2 = s1/2
	a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2))) // height of bisected triangle
	//     create.new.obj     1 - create.new.obj = same as squaring below
	// a = math.Sqrt(r-(math.Pow(s1_2,2)))                         // height of bisected triangle
	b.Sub(r, a) // new short side, r is initially 1
	// b = 1-a                              // new short side
	s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // new hypotenuse / NewSide
	// s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2)))                                 // new hypotenuse / NewSide

	s1.Set(s2) // where is the corresponding s1 = s2 ???????????????????????????????????????????????????????????????????????????????

	p.Mul(numberOfSides, s1)
	// p = numberOfSides*s1
	p_d.Set(p) // p_d should instead be set to p/2 (half p) ????????????????????????????????????????????????????????????????????????
	// p_d = p/2

	var iterInt64 int64 = 0
	for iterInt64 < 5000 {
		iterInt64++
		numberOfSides.Mul(numberOfSides, big.NewFloat(2))                                     // double the number of sides
		s1_2.Quo(s1, big.NewFloat(2))                                                         // halve s1
		a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))                         // height of bisected triangle
		b.Sub(r, a)                                                                           // new short side
		s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // new hypotenuse / NewSide
		s1.Set(s2)
		p.Mul(numberOfSides, s1)
		p_d.Set(p)
		p_d.Quo(p_d, big.NewFloat(2)) // Halve the value of p_d

		if iterInt64 == 24 { // last inter of prior version was 25 ??
			fmt.Println("\n   #1 234567890123456# :: counting the first 16 actual digits of π")
			fmt.Printf("    %.20f is the big.Float of what we have calculated per Archimedes' at 24 iters, 20f\n", p_d)
			fmt.Println("    3.141592653589793238 is the value of π from the web")
			fmt.Println("   #1 234567890123456# :: counting the first 16 actual digits of π")

			fmt.Print(" the above was estimated from a ")
			fmt.Printf("%.0f sided polygon\n", numberOfSides)
			fmt.Printf("%.0f as parsed against ...\n", numberOfSides)
			fmt.Println("100000000 which is one-hundred-million, for comparison to the above line")
			fmt.Println("... Which is to say a 201,326,592 sided polygon, Mr. A. would have wept!\n")
		}

		if iterInt64 == 50 {
			fmt.Println("\n   #1 234567890123456789012345678901# :: counting the actual digits of π")
			fmt.Printf("    %.33f is the big.Float of what we have calculated per Archimedes' at 50 iters, 33f\n", p_d)
			fmt.Println("    3.141592653589793238462643383279502 is the value of π from the web")
			fmt.Println("   #1 234567890123456789012345678901# :: counting the actual digits of π")
			fmt.Println(iterInt64, " iterations were completed yielding 31 correct digits of π")
			fmt.Printf("the above was estimated from a %.0f sided polygon (formatted as a .0f) \n\n", numberOfSides)
		}

		if iterInt64 == 150 {
			fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123#")
			fmt.Printf("  %.95f per Archimedes'\n", p_d)
			fmt.Println("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211 is from web")
			fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123#")
			fmt.Println("           10        20         30       40        50         60        70        80        90 ")
			fmt.Println(iterInt64, " iterations were completed yielding 92 correct digits of π")
			fmt.Printf("Calculated from a %.0f sided polygon\n\n", numberOfSides)
		}

		if iterInt64 == 200 {
			fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890#")
			fmt.Printf("  %.122f per Archimedes'\n", p_d)
			fmt.Println("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709 is from web")
			fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890#")
			fmt.Println("           10        20         30       40        50         60        70        80        90      100        110      120")
			fmt.Println(iterInt64, " iterations were completed yielding 121 correct digits of π")
			fmt.Printf("Calculated from a %.0f sided polygon\n\n\n", numberOfSides)
			fmt.Println(" ... working ...\n\n")
		}

		if iterInt64 == 1200 {
			fmt.Println("... still working, 1200 iterations completed ...\n")
		}
		if iterInt64 == 2200 {
			fmt.Println("... still working, 2200 iterations completed ...\n")
		}
		if iterInt64 == 3200 {
			fmt.Println("... still working, 3200 iterations completed ...\n")
		}
		if iterInt64 == 4200 {
			fmt.Println("... still working, 4200 iterations completed ...\n")
		}

		if iterInt64 == 5000 { // was 5500
			fmt.Printf(" A peek at the result formatted 1500f is: %.1500f \nper Archimedes'\n", p_d) // show the first 1,500 digits of calculated pi
			fmt.Println(iterInt64, " iterations were completed, \n ... which generated a ", numberOfSides, "sided polygon!!\n")
			// fmt.Println(iterInt64, " iterations were completed yielding 2,712 correct digits of π!!!\n")
			fmt.Println("Go's math/big objects were set to a precision value of:", precision)

			fmt.Println("\n\n\n")
			printResultStatsLong(p_d, precision, "AM", 1, "", selection) // sumBig *big.Float, precision int, useAlternateFile string

			fmt.Println("\n\nThese digits are also in a text file named: dataLog-From_AM_Method_lengthy_prints.txt")
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	// store reults in a log file which can be displayed from within the program by selecting option #12
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
	defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
	Hostname, _ := os.Hostname()
	_, err0 := fmt.Fprintf(fileHandle, "\n  -- Archimedes of Syracuse -- selection #%d on %s \n", selection, Hostname)
	check(err0)
	current_time := time.Now()
	_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err6)

	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
	check(err7)

	fmt.Println("\n   -- After hitting Return for menu redisplay, enter '11' for the derivation and proof of the Pythagorean\n\n")
	fmt.Println("Select 12 at menu to display prior results")

}

// written entirely by Richard Woolley
// end of :: ArchimedesBig (selection int) { // case 14:  // --AMFArchimedesBigB
