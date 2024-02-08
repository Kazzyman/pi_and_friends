package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

// case 180:  /// this procedure is really only meant to be used for testing or illustration ...
//
//	... it calculates the root of all integers; and builds a file of all its results
//
// -- AMFxroyCCA
func xRootOfY_continuousCaller(selection int) {
	usingBigFloats = false
	precisionOfRoot = 1700 // default value
	workPiece = 2
	radical_index_number := 2
	fmt.Println("\nThis proc endlessly calculates the root of all integers beginging with 2; and builds a txt file of its results.\n")
	fmt.Println("\nEnter 2 for SquareRoot, or 3 for CubeRoot\n")
	fmt.Scan(&radical_index_number) // pause and request input from the user
	radical_index := radical_index_number
	for 1 == 1 {
		sortedResults = nil // reset the global var *** maybe "fix" this ??
		xRootOfyC(selection, workPiece, radical_index)
		workPiece++
	}
}

// case 180: /// note the C suffix on this and on all func-s used in this func /// should i, can i, just use the non-suffixed func-s ???
func xRootOfyC(selection int, workPiece int, radical_index int) { // calculates either square or cube root of any integer

	var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP // cannot be a global

	TimeOfStartFromTop := time.Now()

	if radical_index == 3 { // if doing a cube root special tolerances are set here for certain problem values, i.e., 2, 11, 17, 3, 4, or 14
		if workPiece > 4 {
			precisionOfRoot = 1655
			fmt.Println("\n Default precision is 1700 \n")
		}
		if workPiece == 2 || workPiece == 11 || workPiece == 17 {
			precisionOfRoot = 580
			fmt.Println("\n resetting precision to 600 \n")
		}
		if workPiece == 3 || workPiece == 4 || workPiece == 14 {
			precisionOfRoot = 875
			fmt.Println("\n resetting precision to 900 \n")
		}
	}
	if radical_index == 2 { // if doing a square root we just use a tolerance of 4 for all workpieces.
		precisionOfRoot = 4
	}

	// radical_index := getInputFromUserAndSetPrecision() // this step is not needed in this continuous automated version

	buildTableOfPerfectProductsC(radical_index) // 825,000 entries

	// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------------------------------------------
	// case 180:

	startBeforeCall := time.Now()

	// for index < 380000 { // the table has 800,000 entries, 400,000 pairs; so index increments by 2 at the bottom of this loop
	for index < 400000 { // the table has 825,000 entries, > 410,000 pairs; so index increments by 2 at the bottom of this loop

		readTheTableOfPPC(index, startBeforeCall, radical_index) // pass-in an index to the table: 380,000 indexs corresponding to the number of pairs of entries

		handlePerfectSquaresAndCubesC(TimeOfStartFromTop, radical_index, selection) // handle the rare case of a perfect square or cube

		if diffOfLarger == 0 || diffOfSmaller == 0 { // Then, it was a perfect square or cube; so, we need to ...
			break // ... out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube
		}

		fileHandle_180, err31 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err31)                                                                                                      // ... gets a file handle to logfile_from_selection_180.txt
		defer fileHandle_180.Close()                                                                                      // It’s idiomatic to defer a Close immediately after opening a file.

		if index == 80000 {
			fmt.Printf("\n80,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			_, err01 := fmt.Fprintf(fileHandle_180, "\n80,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			check(err01)
		}
		if index == 160000 {
			fmt.Printf("\n160,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			_, err02 := fmt.Fprintf(fileHandle_180, "\n160,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			check(err02)
		}
		if index == 240000 {
			fmt.Printf("\n240,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			_, err03 := fmt.Fprintf(fileHandle_180, "\n240,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			check(err03)
		}
		if index == 320000 {
			fmt.Printf("\n320,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			_, err04 := fmt.Fprintf(fileHandle_180, "\n320,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
			check(err04)
		}
		fileHandle_180.Close()

		index = index + 2 // increment the index and read the table again
	} // end of for loop // the above break statement is NOT the only way to exit this for loop, it also terminates after 380,000 iterations of index

	// All of the remaining sections are conditional for workpiece NOT being a perfect square or cube
	// if perfectResult2 == 0 && perfectResult3 == 0 {  // Then, it was NOT a perfect square or cube, so handle that case
	// the remaining sections are only reached after having exited the primary for loop above via a break statement or an exaustive reading of the table ------------
	// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
	// calculate elapsed time
	t_s2 := time.Now()
	elapsed_s2 := t_s2.Sub(TimeOfStartFromTop)

	// the following sections log the final results to a text file (and also does one conditional Printf) -------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------------------------------------------
	fileHandle_180, err311 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err311)                                                                                                      // ... gets a file handle to logfile_from_selection_180.txt
	defer fileHandle_180.Close()                                                                                       // It’s idiomatic to defer a Close immediately after opening a file.

	Hostname, _ := os.Hostname()
	_, err30 := fmt.Fprintf(fileHandle_180, "\n  -- %d root of %d by a ratio of perfect Products -- selection #%d on %s \n", radical_index, workPiece, selection, Hostname)
	check(err30)

	current_time := time.Now()
	_, err36 := fmt.Fprint(fileHandle_180, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err36)

	_, err35 := fmt.Fprintf(fileHandle_180, "%d was total Iterations, precisionOfRoot was %d \n", index, precisionOfRoot)
	check(err35)

	fileHandle_180.Close()

	// Sort the slice sortedResults by its pdiff field :
	// -----------------------------------------------------------------------------------------------------------
	sort.Slice(sortedResults, func(i, j int) bool { return sortedResults[i].pdiff < sortedResults[j].pdiff })

	// display and print the best-fitting result based solely on the lowest pdiff :
	// -----------------------------------------------------------------------------
	// case 180:

	// display the best fitting result :
	if radical_index == 2 {
		if len(sortedResults) > 0 {
			fmt.Printf("\n %0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
		}
	}
	if radical_index == 3 {
		if len(sortedResults) > 0 {
			fmt.Printf("\n %0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
		}
	}

	// Fprint/log the best fitting result :
	fileHandle_180, err310 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err310)                                                                                                      // ... gets a file handle to logfile_from_selection_180.txt
	defer fileHandle_180.Close()                                                                                       // It’s idiomatic to defer a Close immediately after opening a file.

	if radical_index == 2 {
		fmt.Println("\n len of sorted was ", len(sortedResults))
		_, err11 := fmt.Fprintf(fileHandle_180, "len of sorted was %d", len(sortedResults))
		check(err11)
		if len(sortedResults) > 0 {
			_, err48 := fmt.Fprintf(fileHandle_180, "\n %0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
			check(err48)
		}
		fmt.Println("\nThat's all folks!\n\nNext ...")
		_, err22 := fmt.Fprint(fileHandle_180, "\nThat's all folks!\n\nNext ...")
		check(err22)
	}
	if radical_index == 3 {
		fmt.Println("\n len of sorted was ", len(sortedResults))
		_, err11 := fmt.Fprintf(fileHandle_180, "len of sorted was %d", len(sortedResults))
		check(err11)
		if len(sortedResults) > 0 {
			_, err49 := fmt.Fprintf(fileHandle_180, "\n %0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
			check(err49)
		}
		fmt.Println("\nThat's all folks!\n\nNext ...")
		_, err22 := fmt.Fprint(fileHandle_180, "\nThat's all folks!\n\nNext ...")
		check(err22)
	}
	// ------------------ case 180: ---------------------------------------------------------

	TotalRun := elapsed_s2.String() // cast time durations to a String type for Fprintf "formatted print"
	_, err57 := fmt.Fprintf(fileHandle_180, "Total run was %s \n ", TotalRun)
	check(err57)

	fileHandle_180.Close()
}

// in case 180: ----------
func handlePerfectSquaresAndCubesC(TimeOfStartFromTop time.Time, radical_index int, selection int) {
	// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
	// ... it also is responsible for causing the algorithm to terminate via a break if workpiece was a perfect square or cube
	// -------------------------------------------------------------------------------------------------------------------------------
	if diffOfLarger == 0 || diffOfSmaller == 0 { // Then, it was a perfect square or cube

		t_s1 := time.Now()
		elapsed_s1 := t_s1.Sub(TimeOfStartFromTop)

		fileHandle_180, err1 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err1)                  // ... gets a file handle to logfile_from_selection_180.txt
		defer fileHandle_180.Close() // It’s idiomatic to defer a Close immediately after opening a file.

		Hostname, _ := os.Hostname()
		_, err0 := fmt.Fprintf(fileHandle_180, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n",
			radical_index, workPiece, selection, Hostname)
		check(err0)

		current_time := time.Now()
		_, err6 := fmt.Fprint(fileHandle_180, "was run on: ", current_time.Format(time.ANSIC), "\n")
		check(err6)

		TotalRun := elapsed_s1.String() // cast time durations to a String type for Fprintf "formatted print"
		_, err7 := fmt.Fprintf(fileHandle_180, "Total run was %s \n ", TotalRun)
		check(err7)

		if radical_index == 2 {
			_, err8 := fmt.Fprintf(fileHandle_180, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult2)
			check(err8)
		}
		if radical_index == 3 {
			_, err38 := fmt.Fprintf(fileHandle_180, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult3)
			check(err38)
		}
		fileHandle_180.Close()

		// break // break out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube

	} // end of if :: if it was a perfect square or cube
	// -------------------------------------------------------------------------------------------------------------------------------

}

// in case 180:
func readTheTableOfPPC(index int, startBeforeCall time.Time, radical_index int) { // this gets called 380,000 times.

	// The first time it is called index is 0

	// read it ...
	smallerPerfectProductOnce := Table_of_perfect_Products[index]
	// ... and save it locally, do this just-once per func call. // ... index may be 0 up to 380,000
	RootOfsmallerPerfectProductOnce := Table_of_perfect_Products[index+1]
	// ^^^ also read the root wich corresponds

	iter := 0
	for iter < 410000 { // 410,000 loops. Why do we need so many?, Because index may be 0 to 400,000 ?? and we need to read through 825,000 table entries
		iter++

		index = index + 2
		largerPerfectProduct := Table_of_perfect_Products[index]
		// to approximate the root of an imperfect square x we will need a ratio of two perfect squares wich is about equal to x
		// ...we need to find two perfect squares such that one is about x times larger than the other
		// get next perfect square from table for testing to see if it is more than x * bigger than smallerPerfectProductOnce

		if largerPerfectProduct > smallerPerfectProductOnce*workPiece {
			// if largerPerfectProduct is a candidate based on it being just-a-bit larger than workPiece* the smaller PP deal with that, else loop to the next potential

			ProspectiveHitOnLargeSide := largerPerfectProduct                     // make a copy under a more suitable name :)
			rootOfProspectiveHitOnLargeSide := Table_of_perfect_Products[index+1] // the current value of index plus one holds the root of largerPerfectSquare hence the root of ProspectiveHitOnLargeSide

			ProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-2]
			// save that smaller one too //                               ^^ 2 now instead of 1 because we have added roots to the slice
			rootOfProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-1]

			diffOfLarger = ProspectiveHitOnLargeSide - workPiece*smallerPerfectProductOnce
			// diffOfSmaller = -(ProspectiveHitOnSmallerSide - workPiece*smallerPerfectProductOnce) // this was dumb ??
			diffOfSmaller = workPiece*smallerPerfectProductOnce - ProspectiveHitOnSmallerSide

			// detect perfect squares and set global vars to their roots -----------------------------------------------
			if diffOfLarger == 0 {
				fmt.Println(colorCyan, "\n The", radical_index, "root of", workPiece, "is", colorGreen,
					float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce), colorReset, "\n")

				perfectResult2 = (math.Sqrt(float64(workPiece))) // these global values are used later to Fprint to a log file
				perfectResult3 = math.Cbrt(float64(workPiece))
				break // out of the for loop because the workPiece is itself a perfect square
			}
			if diffOfSmaller == 0 {
				fmt.Println(colorCyan, "\n The", radical_index, "root of", workPiece, "is", colorGreen,
					float64(rootOfProspectiveHitOnSmallerSide)/float64(RootOfsmallerPerfectProductOnce), colorReset, "\n")

				perfectResult2 = (math.Sqrt(float64(workPiece))) // these global values are used later to Fprint to a log file
				perfectResult3 = math.Cbrt(float64(workPiece))
				break // out of the for loop because the workPiece is itself a perfect square
			}
			// ---------------------------------------------------------------------------------------------------------
			// case 180:

			// larger side section: of case 180: ---------------------------------------------------------------------------------------------------------------------------
			// --------------------------------------------------------------------------------------------------------------------------------------------------------------

			if diffOfLarger < precisionOfRoot { // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt
				fmt.Println("small PP is", colorCyan, smallerPerfectProductOnce, colorReset, "and, slightly on the higher side of", workPiece,
					"* that we found a PP of", colorCyan, ProspectiveHitOnLargeSide, colorReset, "a difference of", diffOfLarger)

				fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", colorGreen,
					float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce), colorReset)

				fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000)

				fileHandle_180, err1 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check(err1)                  // ... gets a file handle to logfile_from_selection_180.txt
				defer fileHandle_180.Close() // It’s idiomatic to defer a Close immediately after opening a file.

				_, err2 := fmt.Fprint(fileHandle_180, "\nsmall PP is ", smallerPerfectProductOnce, " and, slightly on the higher side of ",
					workPiece, " * that we found a PP of ", ProspectiveHitOnLargeSide, " a difference of ", diffOfLarger)
				check(err2)
				_, err3 := fmt.Fprint(fileHandle_180, "\nthe ", radical_index, " root of ", workPiece, " is calculated as ",
					float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce))
				check(err3)
				_, err6 := fmt.Fprintf(fileHandle_180, "\nwith pdiff of %0.4f, precisionOfRoot is %d\n",
					(float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000, precisionOfRoot)
				check(err6)

				fileHandle_180.Close()

				// save the result to an accumulator array so we can Fprint all such hits at the very end
				// List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce) )
				// corresponding_diffs = append(corresponding_diffs, diffOfLarger)
				// diffs_as_percent = append(diffs_as_percent, float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))

				// in the next five lines we load (append) a record into/to the file (array) of Results
				Result1 := Results{
					result: float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce),
					pdiff:  float64(diffOfLarger) / float64(ProspectiveHitOnLargeSide),
				}
				sortedResults = append(sortedResults, Result1)

				t2 := time.Now()
				elapsed2 := t2.Sub(startBeforeCall)
				// if needed, notify the user that we are still working
				Tim_win = 0.22
				if radical_index == 3 {
					if workPiece > 13 {
						Tim_win = 0.002
					} else {
						Tim_win = 0.0045
					}
				}
				if elapsed2.Seconds() > Tim_win {
					fmt.Println(elapsed2.Seconds(), "Seconds have elapsed ... working ..., time window is: ", Tim_win, "\n")
				}
			}
			// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
			// in case 180:

			// smaller side section: of case 180: ---------------------------------------------------------------------------------------------------------------------------
			// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
			if diffOfSmaller < precisionOfRoot { // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt
				fmt.Println("small PP is", colorCyan, smallerPerfectProductOnce, colorReset, "and, slightly on the lesser side of", workPiece,
					"* that we found a PP of", colorCyan, ProspectiveHitOnSmallerSide, colorReset, "a difference of", diffOfSmaller)

				fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", colorGreen,
					float64(rootOfProspectiveHitOnSmallerSide)/float64(RootOfsmallerPerfectProductOnce), colorReset)

				fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))*100000)

				fileHandle_180, err1 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check(err1)                  // ... gets a file handle to logfile_from_selection_180.txt
				defer fileHandle_180.Close() // It’s idiomatic to defer a Close immediately after opening a file.

				_, err2 := fmt.Fprint(fileHandle_180, "\nsmall PP is ", smallerPerfectProductOnce, " and, slightly on the higher side of ",
					workPiece, " * that we found a PP of ", ProspectiveHitOnLargeSide, " a difference of ", diffOfLarger)
				check(err2)
				_, err3 := fmt.Fprint(fileHandle_180, "\nthe ", radical_index, " root of ", workPiece, " is calculated as ",
					float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce))
				check(err3)
				_, err6 := fmt.Fprintf(fileHandle_180, "\nwith pdiff of %0.4f, precisionOfRoot is %d\n",
					(float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000, precisionOfRoot)
				check(err6)

				fileHandle_180.Close()

				// save the result to three accumulator arrays so we can Fprint all such hits, diffs, and p-diffs, at the very end of run
				// List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce) )
				// corresponding_diffs = append(corresponding_diffs, diffOfSmaller)
				// diffs_as_percent = append(diffs_as_percent, float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))

				// in the next five lines we load (append) a record into/to the file (array) of Results
				Result1 := Results{
					result: float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce),
					pdiff:  float64(diffOfSmaller) / float64(ProspectiveHitOnSmallerSide),
				}
				sortedResults = append(sortedResults, Result1)

				t2 := time.Now()
				elapsed2 := t2.Sub(startBeforeCall)
				// if needed, notify the user that we are still working
				Tim_win = 0.22
				if radical_index == 3 {
					if workPiece > 13 {
						Tim_win = 0.002
					} else {
						Tim_win = 0.0045
					}
				}
				if elapsed2.Seconds() > Tim_win {
					fmt.Println(elapsed2.Seconds(), "Seconds have elapsed ... working ..., time window is: ", Tim_win, "\n")
				}
			}
			// ---------------------------------------------------------------------------------------------------------------------------------------------------------------

			break // each time we find a prospect we break out of the for loop --- if we found any prospects using the current index value we break

		} // end of if :: if largerPerfectProduct > smallerPerfectProductOnce*workPiece  //  we only handle reads that were big enough to be prospects

	} // this is the end of the afforementioned for loop that we break out of each time we have found a prospect and handled it

} // the end of the readTheTableOfPP func that gets called 380,000 times

// Build a table of 825,000 pairs of PPs with their roots, does either squares or cubes:
// --------------------------------------------------------------------------------------

func buildTableOfPerfectProductsC(radical_index int) { // C suffix denotes case 180:

	var PerfectProduct int
	Table_of_perfect_Products = nil // this fixed my bug
	root := 10
	iter := 0
	for iter < 825000 { // a table of 825,000 pairs: PPs with their roots. That ought to do it !!
		iter++
		root++
		if radical_index == 3 { // build an array of perfect cubes
			PerfectProduct = root * root * root
		}
		if radical_index == 2 { // build an array of perfect squares
			PerfectProduct = root * root
		}
		Table_of_perfect_Products = append(Table_of_perfect_Products, PerfectProduct)
		Table_of_perfect_Products = append(Table_of_perfect_Products, root) // the root of the prior PP
	}
}

// written entirely by Richard Woolley
// end of case 180: -- AMFxroyCCB
