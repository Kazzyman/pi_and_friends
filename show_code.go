package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// --------------- Display code to user functions follow -----------------------------------------------------------------------------------------//
// ------------------------------------------------------------------------------------------------------------------------------------------------//
// These functions used to take the name of the one massive file main.go which was obtained from getTheNameOfThisSourceCodeFile() as filenameOfThisFile
// ... which is now defunct, though still can be found in menu_elements.go
// In addition, these functions used to return filenameOfThisFile obtained from getTheNameOfThisSourceCodeFile(), requiring the pattern matching

func Show_Gauss_Legendre() { // case 57:
	pattern1 := regexp.MustCompile(`(?s)AMFGauss_LegendreA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFGauss_LegendreB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "37_Gauss_Legendre.go")
}
func ShowArchimedesBig() { // case 34:
	pattern1 := regexp.MustCompile(`(?s)AMFArchimedesBigA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFArchimedesBigB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "14_ArchimedesBig.go")
}
func ShowMain() { // case 3:
	pattern1 := regexp.MustCompile(`(?s)AMFmainA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFmainB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "main.go")
}
func ShowMenus() { // case 4:
	pattern1 := regexp.MustCompile(`(?s)AMFmenusA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFmenusB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "menu.go")
}
func ShowxRootOfY_continuousCaller() { // case 380:
	pattern1 := regexp.MustCompile(`(?s)AMFxroyCCA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFxroyCCB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "180_x_root_of_y_Continuous.go")
}
func ShowxRootOfy() { // case 38:
	pattern1 := regexp.MustCompile(`(?s)AMFxRootOfyA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFxRootOfyB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "18_x_root_of_y.go")
}
func ShowGottfriedWilhelmLeibniz() { // case 26:
	pattern1 := regexp.MustCompile(`(?s)AMFGottfriedWilhelmLeibnizA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFGottfriedWilhelmLeibnizB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "6_GottfriedWilhelmLeibniz.go")
}
func ShowGregoryLeibniz() { // case 27:
	pattern1 := regexp.MustCompile(`(?s)AMFGregoryLeibnizA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFGregoryLeibnizB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "7_GregoryLeibniz.go")
}
func ShowJohnWallis() { // case 28:
	pattern1 := regexp.MustCompile(`(?s)AMFJohnWallisA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFJohnWallisB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "8_JohnWallis.go")
}
func ShowEulersNumber() { // case 29:
	pattern1 := regexp.MustCompile(`(?s)AMFEulersNumberA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFEulersNumberB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "9_EulersNumber.go")
}
func ShowErdosBorwein() { // case 30:
	pattern1 := regexp.MustCompile(`(?s)AMFErdosBorweinA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFErdosBorweinB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "10_ErdosBorwein.go")
}
func ShowcompareFloat64withBigFloats() { // case 70:
	pattern1 := regexp.MustCompile(`(?s)AMFcompareFloat64withBigFloatsA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFcompareFloat64withBigFloatsB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "50_compareFloat64withBigFloats.go")
}
func ShowNilakantha_Somayaji_with_big_Float_types() { // case 25:
	pattern1 := regexp.MustCompile(`(?s)AMFNilakantha_Somayaji_with_big_Float_typesA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFNilakantha_Somayaji_with_big_Float_typesB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "5_Nilakantha_Somayaji_with_big_Float_types.go")
}
func Showchud() { // case 35:
	pattern1 := regexp.MustCompile(`(?s)AMFchudA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFchudB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "15_chud.go")
}
func ShowBBPF() { // case 22:
	pattern1 := regexp.MustCompile(`(?s)AMFBBPFA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFBBPFB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "2_BBPF.go")
}
func ShowLeibniz_method_one_billion_iters() { // case 64:
	pattern1 := regexp.MustCompile(`(?s)AMFLeibniz_method_one_billion_itersA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFLeibniz_method_one_billion_itersB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "44_Leibniz_method_one_billion_iters.go")
}
func ShowBBPfConcurent() { // case 61:
	pattern1 := regexp.MustCompile(`(?s)AMFBBPfConcurentA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFBBPfConcurentB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "41_BBPfConcurent.go")
}
func Showbbp_formula() { // case 62:
	pattern1 := regexp.MustCompile(`(?s)AMFbbp_formulaA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFbbp_formulaB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "42_bbp_formula.go")
}
func Shownifty_scoreBoard() { // case 60:
	pattern1 := regexp.MustCompile(`(?s)AMFnifty_scoreBoardA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFnifty_scoreBoardB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "40_nifty_scoreBoard.go")
}
func ShowTheSpigot() { // case 39:
	pattern1 := regexp.MustCompile(`(?s)AMFTheSpigotA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFTheSpigotB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "19_TheSpigot.go")
}
func ShowdeleteAllLogFiles() { // case 20999:
	pattern1 := regexp.MustCompile(`(?s)AMFdeleteAllLogFilesA.*?`) // Create a regular expression that matches the first flag anywhere in the line
	pattern2 := regexp.MustCompile(`(?s)AMFdeleteAllLogFilesB.*?`) // Create a regular expression that matches the second flag anywhere in the line
	doTheRest(pattern1, pattern2, "deleteAllLogFiles.go")
}

// As explained above, the var filenameOfThisFile used to be the pathname of the one single massive file main.go, requiring all this pattern matching.
// func doTheRest(pattern1 *regexp.Regexp, pattern2 *regexp.Regexp, file *os.File) {  // decided that it would not be prudent to pass the file handle
func doTheRest(pattern1 *regexp.Regexp, pattern2 *regexp.Regexp, filenameOfThisFile string) { // better to pass the name of the file instead
	file, err := os.Open(filenameOfThisFile) // create a handle to our file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()                // It’s idiomatic to defer a Close immediately after opening a file.
	alreadyFoundFirstTag := false     // Create a boolean variable to track if we are within the desired range of lines we want to print
	scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line
	countTheLines := 1
	for scanner.Scan() { // Iterate over the lines in our file
		countTheLines++
		line := scanner.Text() // read a line and store the line in var line
		// deal with the case of having read the line containing the first flag
		if pattern1.MatchString(line) { // if line contains pattern1 (the AMF string in the comment at the begining of the func)
			alreadyFoundFirstTag = !alreadyFoundFirstTag // flips this bool to true
			fmt.Println(colorCyan, line)                 // print line containing first flag
			continue                                     // we read another line
		}
		// deal with the case of having encountered the line containing the second flag
		if pattern2.MatchString(line) { // if line contains pattern2 (the AMF string in the comment at the end of the func)
			fmt.Println(colorCyan, line)
			break // we break to stop reading or printing additional lines
		}
		// deal with the case of having already found and printed the line containing the first flag
		// print the line following the line containing the first flag, and all interviening lines
		if alreadyFoundFirstTag {
			fmt.Println(colorCyan, line)
			continue // read another line
		}
	}
	fmt.Println(colorYellow, "The total number of lines in this set is: ", countTheLines)
}
