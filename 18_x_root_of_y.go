package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

// case 18: -- AMFxRootOfyA
func xRootOfy(selection int) { // calculates either square or cube root of any integer

	usingBigFloats = false

	var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP

	TimeOfStartFromTop := time.Now()

	radical_index := getInputFromUserAndSetPrecision()

	buildTableOfPerfectProducts(radical_index) // 800,000 entries, 400,000 pairs

	// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------------------------------------------

	startBeforeCall := time.Now()

	for index < 400000 { // the table has 825,000 entries, > 410,000 pairs; so index increments by 2 at the bottom of this loop (200,000 iters)

		readTheTableOfPP(index, startBeforeCall, radical_index) // pass-in an index to the table: 400,000 indexs corresponding to the number of pairs of entries

		handlePerfectSquaresAndCubes(TimeOfStartFromTop, radical_index, selection) // handle the rare case of a perfect square or cube

		if diffOfLarger == 0 || diffOfSmaller == 0 { // Then, it was a perfect square or cube; so, we need to ...
			break // ... out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube
		}

		if index == 80000 {
			fmt.Println("\n80,000 ... still working ...")
		}
		if index == 160000 {
			fmt.Println("\n160,000 ... still working ...")
		}
		if index == 240000 {
			fmt.Println("\n240,000 ... still working ...")
		}
		if index == 320000 {
			fmt.Println("\n320,000 ... still working, almost there ...\n")
		}

		index = index + 2 // increment the index and read the table again
	} // end of for loop // the above break statement is NOT the only way to exit this for loop, it also terminates after 380,000 iterations of index

	// All of the remaining sections are conditional for workpiece NOT being a perfect square or cube
	if perfectResult2 == 0 && perfectResult3 == 0 { // Then, it was NOT a perfect square or cube, so handle that case
		// the remaining sections are only reached after having exited the primary for loop above via a break statement or an exaustive reading of the table ------------
		// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
		// calculate elapsed time
		t_s2 := time.Now()
		elapsed_s2 := t_s2.Sub(TimeOfStartFromTop)

		// the following sections log the final results to a text file (and also does one conditional Printf) -------------------------------------------------
		// -----------------------------------------------------------------------------------------------------------------------------------------------------
		fileHandle, err31 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err31)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandle.Close()                                                                                                 // It’s idiomatic to defer a Close immediately after opening a file.

		Hostname, _ := os.Hostname()
		_, err30 := fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of perfect Products -- selection #%d on %s \n", radical_index, workPiece, selection, Hostname)
		check(err30)

		current_time := time.Now()
		_, err36 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
		check(err36)

		// index = index
		_, err35 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", index)
		check(err35)

		// Sort the slice sortedResults by its pdiff field :
		// -----------------------------------------------------------------------------------------------------------
		sort.Slice(sortedResults, func(i, j int) bool { return sortedResults[i].pdiff < sortedResults[j].pdiff })

		/*
		   // print the sorted slice twice; once for each field
		       fmt.Println("Here are the results:")
		       resultCount := 1
		       for _, result := range sortedResults {
		           fmt.Printf("%d, %0.16f \n", resultCount, result.result)
		           resultCount++
		       }
		       fmt.Println("And here are the p-diffs:")
		       pdiffCount := 1
		       for _, result := range sortedResults {
		           fmt.Printf("%d, %0.16f \n", pdiffCount, result.pdiff)
		           pdiffCount++
		       }
		*/

		// display and print the best-fitting result based solely on the lowest pdiff :
		// -----------------------------------------------------------------------------

		// display the best fitting result :
		if radical_index == 2 {
			fmt.Printf("%0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
		}
		if radical_index == 3 {
			fmt.Printf("%0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
		}

		// Fprint/log the best fitting result :
		if radical_index == 2 {
			_, err48 := fmt.Fprintf(fileHandle, "%0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
			check(err48)
		}
		if radical_index == 3 {
			_, err49 := fmt.Fprintf(fileHandle, "%0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
			check(err49)
		}
		// ---- we are still in case 18: -------------------------------------------------------------------------

		TotalRun := elapsed_s2.String() // cast time durations to a String type for Fprintf "formatted print"
		_, err57 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
		check(err57)

		fileHandle.Close()

		/*
		   // all this crap with these 3 arrays was cute, but it does not sort as a record with fields, so it is not what I want
		               array_len := len(List_of_2_results_case18)
		               _ , err8 := fmt.Fprintf(fileHandle, "%d was len of array \n", array_len)
		                   check(err8)
		               if array_len > 0 {
		                   index := 0
		                   for array_len > 0 {
		                       result_from_array := List_of_2_results_case18[index]
		                       array_len--
		                        _ , err9 := fmt.Fprintf(fileHandle, "%0.16f with a diff of %d, percent diff of %0.4f percent\n",
		                           result_from_array, corresponding_diffs[index], diffs_as_percent[index]*100000)
		                               check(err9)
		                       index++
		                   }
		               }
		               List_of_2_results_case18 = nil
		               corresponding_diffs = nil
		*/
	}
} // --- we are in case 18:

func handlePerfectSquaresAndCubes(TimeOfStartFromTop time.Time, radical_index int, selection int) {
	// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
	// ... it also is responsible for causing the algorithm to terminate via a break if workpiece was a perfect square or cube
	// -------------------------------------------------------------------------------------------------------------------------------
	if diffOfLarger == 0 || diffOfSmaller == 0 { // Then, it was a perfect square or cube

		t_s1 := time.Now()
		elapsed_s1 := t_s1.Sub(TimeOfStartFromTop) // need to pass this to the func we are planning to build ?? NO, "two" "perfect".

		fileNameToWriteTo := "dataLog-From_calculate-pi-and-friends.txt" // would have been used/needed if we emplement a func for this.

		// fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		fileHandle, err1 := os.OpenFile(fileNameToWriteTo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err1)              // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandle.Close() // It’s idiomatic to defer a Close immediately after opening a file.

		Hostname, _ := os.Hostname()
		_, err0 := fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n",
			radical_index, workPiece, selection, Hostname)
		check(err0)

		current_time := time.Now()
		_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
		check(err6)

		TotalRun := elapsed_s1.String() // cast time durations to a String type for Fprintf "formatted print"
		_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
		check(err7)

		if radical_index == 2 {
			_, err8 := fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult2)
			check(err8)
		}
		if radical_index == 3 {
			_, err38 := fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult3)
			check(err38)
		}

		fileHandle.Close()

		// break // break out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube

	} // end of if :: if it was a perfect square or cube
	// -------------------------------------------------------------  CASE 18: ------------------------------------------------------------------

}

func readTheTableOfPP(index int, startBeforeCall time.Time, radical_index int) { // this gets called 400,000 times.

	// The first time it is called index is 0

	// read it ...
	smallerPerfectProductOnce := Table_of_perfect_Products[index]
	// ... and save it locally, do this just-once per func call. // ... index may be 0 up to 380,000
	RootOfsmallerPerfectProductOnce := Table_of_perfect_Products[index+1]
	// ^^^ also read the root wich corresponds

	iter := 0
	for iter < 410000 { // 410,000 loops. Why do we need so many?, Because we need to read through 825,000 table entries pairs
		iter++ //  ... iters are therefore half the number of pairs. There are actually 1,600,000 items, but who's counting?
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
			// we are in case 18:

			// larger side section: ----------------------------------------------------------------------------------------------------------------------------------------
			// --------------------------------------------------------------------------------------------------------------------------------------------------------------

			if diffOfLarger < precisionOfRoot { // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt
				fmt.Println("small PP is", colorCyan, smallerPerfectProductOnce, colorReset, "and, slightly on the higher side of", workPiece,
					"* that we found a PP of", colorCyan, ProspectiveHitOnLargeSide, colorReset, "a difference of", diffOfLarger)

				fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", colorGreen,
					float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce), colorReset)

				fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000)

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
				Tim_win = 0.178
				if radical_index == 3 {
					if workPiece > 13 {
						Tim_win = 0.0012
					} else {
						Tim_win = 0.003
					}
				}
				if elapsed2.Seconds() > Tim_win {
					fmt.Println(elapsed2.Seconds(), "Seconds have elapsed ... working ...\n")
				}
			}
			// ---------------------------------------------------------------------------------------------------------------------------------------------------------------

			// we are in case 18:
			// smaller side section: ----------------------------------------------------------------------------------------------------------------------------------------
			// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
			if diffOfSmaller < precisionOfRoot { // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt
				fmt.Println("small PP is", colorCyan, smallerPerfectProductOnce, colorReset, "and, slightly on the lesser side of", workPiece,
					"* that we found a PP of", colorCyan, ProspectiveHitOnSmallerSide, colorReset, "a difference of", diffOfSmaller)

				fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", colorGreen,
					float64(rootOfProspectiveHitOnSmallerSide)/float64(RootOfsmallerPerfectProductOnce), colorReset)

				fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))*100000)

				// save the result to three accumulator arrays so we can Fprint all such hits, diffs, and p-diffs, at the very end of run
				// List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce) )
				// corresponding_diffs = append(corresponding_diffs, diffOfSmaller)
				// diffs_as_percent = append(diffs_as_percent, float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))
				// ***** ^^^^ ****** the preceeding was replaced with the following five lines *******************************************

				// in the next five lines we load (append) a record into/to the file (array) of Results
				Result1 := Results{
					result: float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce),
					pdiff:  float64(diffOfSmaller) / float64(ProspectiveHitOnSmallerSide),
				}
				sortedResults = append(sortedResults, Result1)

				t2 := time.Now()
				elapsed2 := t2.Sub(startBeforeCall)
				// if needed, notify the user that we are still working
				Tim_win = 0.178
				if radical_index == 3 {
					if workPiece > 13 {
						Tim_win = 0.0012
					} else {
						Tim_win = 0.003
					}
				}
				if elapsed2.Seconds() > Tim_win {
					fmt.Println(elapsed2.Seconds(), "Seconds have elapsed ... working ...\n")
				}
			} // end of if
			// -------------  we are in case 18:   we are in case 18:   we are in case 18:   we are in case 18:   we are in case 18: ----------------

			break // each time we find a prospect we break out of the for loop --- if we found any prospects using the current index value we break

		} // end of if :: if largerPerfectProduct > smallerPerfectProductOnce*workPiece  //  we only handle reads that were big enough to be prospects
	} // this is the end of the afforementioned for loop that we break out of each time we have found a prospect and handled it
} // the end of the readTheTableOfPP func that gets called 380,000 times

func getInputFromUserAndSetPrecision() int {
	radical_index := 2 // we set it to 2 becuase if the user just hits enter we want radical_index set as 2 by default
	skip_redoing_loop = 0
	for radical_index > 1 && skip_redoing_loop == 0 {
		// radical_index will initially be 2, and the skip flag will be initially 0, therefore we will initially enter this for loop
		fmt.Println("\n\nEnter 2 for SquareRoot or 3 for CubeRoot")
		fmt.Scan(&radical_index) // just hitting enter will leave radical_index set as 2  // same in Unix or Windows variants

		for radical_index > 0 { // but, if we entered any value for radical_index
			if radical_index < 2 || radical_index > 3 { // only then do we check to see if it is a 2 or a 3
				fmt.Println(colorPurple,
					"\n No, for the radical, enter either 2 or 3, just a 2 or a 3, \n ... no other number for the radical please \n", colorReset)
				radical_index = 2 // an entered 2 or 3 would have bypassed this, but this value of 2 is needed to re-enter/continue the outer for loop
				break             // out of inner loop only, and with skip flag still set to 0
			} else { // if a 2 or a 3 was entered go ahead and prompt for a workPiece
				for 1 == 1 {
					fmt.Println("\n\nEnter any integer that you wish to find the", radical_index, "Root of")
					fmt.Scan(&workPiece)
					// if a float was entered earlier for radical_index, the fractional part will be assigned to workPiece, without prompting (which is buggy on the part of Go)
					// the above is the same in Unix or Windows variants
					if workPiece == 0 || workPiece == 1 {
						fmt.Println(colorRed, "\n You must enter an integer greater than one", colorReset)
					} else {
						break // prompt for a valid workPiece and loop back to re scan, or break out of said loop herein
					}
				}
				skip_redoing_loop = 1
				break // thus sending us back to the outer loop with the directive to skip redoing it per "skip_redoing_loop = 1" on previous line
			}
		}
		if radical_index < 1 { //
			fmt.Println(colorPurple,
				"\n Zero is dopey, for the radical, enter either 2 or 3, just a 2 or a 3, \n ... no other number for the radical please \n", colorReset)
			radical_index = 2 // go fish
		}
	}

	if radical_index == 3 { // if doing a cube root special tolerances are set here for certain problem values, i.e., 2, 11, 17, 3, 4, or 14
		if workPiece > 4 {
			precisionOfRoot = 1700
			fmt.Println("\n Default precision is 1700 \n")
		}
		if workPiece == 2 || workPiece == 11 || workPiece == 17 {
			precisionOfRoot = 600
			fmt.Println("\n resetting precision to 600 \n")
		}
		if workPiece == 3 || workPiece == 4 || workPiece == 14 {
			precisionOfRoot = 900
			fmt.Println("\n resetting precision to 900 \n")
		}
	}
	if radical_index == 2 { // if doing a square root we just use a tolerance of 4 for all workpieces.
		precisionOfRoot = 4
	}
	return radical_index
}

// Build a table of 825,000 pairs of PPs with their roots, does either squares or cubes:
// ----- still in case 18: ---------------------------------------------------------------------------------

func buildTableOfPerfectProducts(radical_index int) {

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
	// written entirely by Richard Woolley
} // end of xRootOfy() // case 18: set, -- AMFxRootOfyB
