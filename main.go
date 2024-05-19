package main

// Unix-0713-0524-v2-dev-2023.go

// The initial inspiration for all of this was : Veritassium https://youtu.be/gMlf1ELvRzc?list=LL
/*
   Go.lang is a relatively-new general-purpose freely-available programming language developed by google.

   One can obtain the Go Compiler from : https://go.dev/dl/
*/

// Unix/Linux/Mac variant build instructions : ==========================================================================
/*
   compile with: "go build -o desired-name-of-your-executable name-of-this-source-code-file.go"
   ... thereafter you can run it on a Unix system with "/fullpathname/desired-name-of-your-executable"
       ( e.g.  "./Richards_go_project" )
   ... or, having first obtained the Go compiler, ... just run the source with: "go run name-of-this-source-code-file.go"
*/
// =======================================================================================================================

// Windows variant : ----------------------------------------------------------------------------------------------------
/*
   build with: "go build -o desired-name-of-your-executable.exe name-of-this-source-code-file.go"
   ... or, having first obtained the Go compiler, ... just run the source with: "go run name-of-this-source-code-file.go"
*/
// -----------------------------------------------------------------------------------------------------------------------
// change
// -- AMFmainA-a

import (
	"fmt"    // Used for printing etc.
	"regexp" // used in multiple instance to create regular expression patterns
)

func main() { // top-level program logic flow -- explore several ways to calculate Pi, plus THREE other constants --AMFmain-A-b

	filenameOfThisFile := "/Users/quasar/pi-main/main.go"

	filenameOfThisFile1 := "/Users/quasar/pi-main/main.go"
	totalLines1, nonEmptyLines1 := reportSLOCstats(filenameOfThisFile1) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile2 := "/Users/quasar/pi-main/2_BBPF.go"
	totalLines2, nonEmptyLines2 := reportSLOCstats(filenameOfThisFile2) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile5 := "/Users/quasar/pi-main/5_Nilakantha_Somayaji_with_big_Float_types.go"
	totalLines5, nonEmptyLines5 := reportSLOCstats(filenameOfThisFile5) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile6 := "/Users/quasar/pi-main/6_GottfriedWilhelmLeibniz.go"
	totalLines6, nonEmptyLines6 := reportSLOCstats(filenameOfThisFile6) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile7 := "/Users/quasar/pi-main/7_GregoryLeibniz.go"
	totalLines7, nonEmptyLines7 := reportSLOCstats(filenameOfThisFile7) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile8 := "/Users/quasar/pi-main/8_JohnWallis.go"
	totalLines8, nonEmptyLines8 := reportSLOCstats(filenameOfThisFile8) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile9 := "/Users/quasar/pi-main/9_EulersNumber.go"
	totalLines9, nonEmptyLines9 := reportSLOCstats(filenameOfThisFile9) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile10 := "/Users/quasar/pi-main/10_ErdosBorwein.go"
	totalLines10, nonEmptyLines10 := reportSLOCstats(filenameOfThisFile10) // another locally-defined func; returns, and creates, local values of predetermined type

	/*
		// todo: need to add tags in 11_display_pythagorean.go
		filenameOfThisFile11 := "/Users/quasar/pi-main/11_display_pythagorean.go"
		totalLines11, nonEmptyLines11 := reportSLOCstats(filenameOfThisFile11) // another locally-defined func; returns, and creates, local values of predetermined type

	*/

	filenameOfThisFile14 := "/Users/quasar/pi-main/14_ArchimedesBig.go"
	totalLines14, nonEmptyLines14 := reportSLOCstats(filenameOfThisFile14) // another locally-defined func; returns, and creates, local values of predetermined type
	
	filenameOfThisFile15 := "/Users/quasar/pi-main/15_chud.go"
	totalLines15, nonEmptyLines15 := reportSLOCstats(filenameOfThisFile15) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile16 := "/Users/quasar/pi-main/16_MonteCarloMethod.go"
	totalLines16, nonEmptyLines16 := reportSLOCstats(filenameOfThisFile16) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile18 := "/Users/quasar/pi-main/18_x_root_of_y.go"
	totalLines18, nonEmptyLines18 := reportSLOCstats(filenameOfThisFile18) // another locally-defined func; returns, and creates, local values of predetermined type

	/*
		// Check the file descriptor limits
		var rLimit18 syscall.Rlimit
		err18 := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit18)
		if err18 != nil {
			log.Fatalf("Error getting file descriptor limit after 18: %v\n", err18)
		}
		fmt.Printf("Current file descriptor limit: %d\n", rLimit18.Cur)

	*/

	filenameOfThisFile19 := "/Users/quasar/pi-main/19_TheSpigot.go"
	totalLines19, nonEmptyLines19 := reportSLOCstats(filenameOfThisFile19) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile37 := "/Users/quasar/pi-main/37_Gauss_Legendre.go"
	totalLines37, nonEmptyLines37 := reportSLOCstats(filenameOfThisFile37) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile40 := "/Users/quasar/pi-main/40_nifty_scoreBoard.go"
	totalLines40, nonEmptyLines40 := reportSLOCstats(filenameOfThisFile40) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile41 := "/Users/quasar/pi-main/41_BBPfConcurent.go"
	totalLines41, nonEmptyLines41 := reportSLOCstats(filenameOfThisFile41) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile42 := "/Users/quasar/pi-main/42_bbp_formula.go"
	totalLines42, nonEmptyLines42 := reportSLOCstats(filenameOfThisFile42) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile44 := "/Users/quasar/pi-main/44_Leibniz_method_one_billion_iters.go"
	totalLines44, nonEmptyLines44 := reportSLOCstats(filenameOfThisFile44) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile50 := "/Users/quasar/pi-main/50_compareFloat64withBigFloats.go"
	totalLines50, nonEmptyLines50 := reportSLOCstats(filenameOfThisFile50) // another locally-defined func; returns, and creates, local values of predetermined type

	/*
		// todo: need to add tags in 96_using_this_app.go
		filenameOfThisFile96 := "/Users/quasar/pi-main/96_using_this_app.go"
		totalLines96, nonEmptyLines96 := reportSLOCstats(filenameOfThisFile96) // another locally-defined func; returns, and creates, local values of predetermined type

	*/

	filenameOfThisFile180 := "/Users/quasar/pi-main/180_x_root_of_y_Continuous.go"
	totalLines180, nonEmptyLines180 := reportSLOCstats(filenameOfThisFile180) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFileDAF := "/Users/quasar/pi-main/deleteAllLogFiles.go"
	totalLinesDAF, nonEmptyLinesDAF := reportSLOCstats(filenameOfThisFileDAF) // another locally-defined func; returns, and creates, local values of predetermined type

	/*
		// todo: need to add tags in globals_const.go
		filenameOfThisFileGC := "/Users/quasar/pi-main/globals_const.go"
		totalLinesGC, nonEmptyLinesGC := reportSLOCstats(filenameOfThisFileGC) // another locally-defined func; returns, and creates, local values of predetermined type

	*/

	filenameOfThisFileMenu := "/Users/quasar/pi-main/menu.go"
	totalLinesMenu, nonEmptyLinesMenu := reportSLOCstats(filenameOfThisFileMenu) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFileMenuElements := "/Users/quasar/pi-main/menu_elements.go"
	totalLinesMenuElements, nonEmptyLinesMenuElements := reportSLOCstats(filenameOfThisFileMenuElements) // another locally-defined func; returns, and creates, local values of predetermined type

	/*
		// todo: need to add tags in menu_switch.go
		filenameOfThisFileMenuSwitch := "/Users/quasar/pi-main/menu_switch.go"
		totalLinesMenuSwitch, nonEmptyLinesMenuSwitch := reportSLOCstats(filenameOfThisFileMenuSwitch) // another locally-defined func; returns, and creates, local values of predetermined type

	*/

	filenameOfThisFilePR := "/Users/quasar/pi-main/printResultStatsLong.go"
	totalLinesPR, nonEmptyLinesPR := reportSLOCstats(filenameOfThisFilePR) // another locally-defined func; returns, and creates, local values of predetermined type

	/*
		// todo: need to add tags in show_code.go
		filenameOfThisFileSC := "/Users/pi-main/quasar/show_code.go"
		totalLinesSC, nonEmptyLinesSC := reportSLOCstats(filenameOfThisFileSC) // another locally-defined func; returns, and creates, local values of predetermined type

	*/

	totalLines := totalLines1 + totalLines2 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 +
		totalLines10 + totalLines14 + totalLines15 + totalLines16 +
		+totalLines18 + totalLines19 +
		totalLines37 + totalLines40 + totalLines41 +
		totalLines42 + totalLines44 + totalLines50 +
		totalLines180 + totalLinesDAF + totalLinesMenu + totalLinesMenuElements +
		totalLinesPR

	nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 +
		nonEmptyLines10 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 +
		nonEmptyLines18 + nonEmptyLines19 +
		nonEmptyLines37 + nonEmptyLines40 + nonEmptyLines41 +
		nonEmptyLines42 + nonEmptyLines44 + nonEmptyLines50 +
		nonEmptyLines180 + nonEmptyLinesDAF + nonEmptyLinesMenu + nonEmptyLinesMenuElements +
		nonEmptyLinesPR

	//
	// The following is for menu header data generation (inception: stripped file name) //  <<------------ below ---------<<
	// ... there should be exactly nine of each comment tag in this dev Windows ver
	// /* Unix variant
	re2 := regexp.MustCompile(`Unix-(.+)\.go`)
	// Unix variant */
	//
	/* Windows variant
	    re2 := regexp.MustCompile(`win-(.+)\.go`)
	Windows variant */

	match2 := re2.FindStringSubmatch(filenameOfThisFile) // grab the stuff between win- and .go (or Unix- and .go) in the name of this file
	SansVerOfNameOfThisFile := "pi-main"                 // this var, having been initialize to "string", is of type string
	if len(match2) >= 2 {

		// /* Unix variant
		SansVerOfNameOfThisFile = match2[1] // Unix SansVerOfNameOfThisFile is loaded with our base file name
		// Unix variant */
		//
		/* Windows variant
		    SansVerOfNameOfThisFile = match2[1]                   // Windows SansVerOfNameOfThisFile is loaded with our base file name
		Windows variant */

	} else {
		fmt.Println("SansVerOfNameOfThisFile via match2 not found in main")
	}

	for {
		DisplayMenus(totalLines, nonEmptyLines, filenameOfThisFile, SansVerOfNameOfThisFile)
		fmt.Println(colorRed, "Finished. Hit Enter to redisplay the Main Menu", colorReset) // this will be the last line of every case #:
		var JustToPauseTheLoop int

		// /* Unix variant
		fmt.Scanf("%d", &JustToPauseTheLoop) // must pause the endless loop after every switch case to prevent menu redisplay over results
		// fmt.Scan(&JustToPauseTheLoop) // Scan() would work as a pause, but it requires data input to continue, so we use a Scanf() instead
		// Unix variant */
		//
		/* Windows variant
		    fmt.Scanf("%d", &JustToPauseTheLoop) // must pause the endless loop after every switch case to prevent menu redisplay over results
		    fmt.Scanf("%d", &JustToPauseTheLoop)
		        //fmt.Scan(&JustToPauseTheLoop) // Scan does not work quite as well as the Scanf combo above (but a single Scanf worked well in Unix environment)
		Windows variant */

	}
} // end of main --AMFmain-A-c

func check(e error) { // create a func named check which takes one parameter "e" of type error
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
