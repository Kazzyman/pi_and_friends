package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func getTheNameOfThisSourceCodeFile() (filenameOfThisFile string) { // filenameOfThisFile var is created local to this func by its having been listed as a returned val
	// get the name of this source code file to be used in the display-of-source-code funcs
	pathNameOfMeAsExe := os.Args[0] // the objective is to find this source code file by name from the command line Args
	fmt.Println(pathNameOfMeAsExe)
	// /* Unix variant
	re := regexp.MustCompile(`Unix(.*?)2023`)
	// Unix variant */
	//
	/* Windows variant
	    re := regexp.MustCompile(`\\exe\\(.*?)\.exe`)
	Windows variant */

	match := re.FindStringSubmatch(pathNameOfMeAsExe) // grab the stuff between /exe/ and .exe in the full pathname of the generated executable
	var substring string                              // ... both Windows and Unix  ^^^
	if len(match) >= 2 {

		// /* Unix variant
		substring = match[0] // Unix substring is loaded with our base file name
		// Unix variant */
		//
		/* Windows variant
		    substring = match[1]                   // Windows substring is loaded with our base file name
		Windows variant */

	} else {
		fmt.Printf("\n\n rick was here: %s\n\n", pathNameOfMeAsExe)

		fmt.Println("substring not found in getTheNameOfThisSourceCodeFile()")
	}
	filenameOfThisFile = substring + ".go" // append the .go on the end of the semi-stripped file name
	return                                 // a naked return of the defined/specified value
}

// --AMFreportSLOCstats01a
func reportSLOCstats(filenameOfThisFile string) (totalLines, nonEmptyLines int) { // returns two values of type int
	file, err := os.Open(filenameOfThisFile) // open this source code file and get a handle to it
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close() // It’s idiomatic to defer a Close immediately after opening a file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		totalLines++ // this var was created locally by having declared it as a return value in this funcs def
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			nonEmptyLines++ // this var was created locally by having declared it as a return value in this funcs def
		}
	}
	return totalLines, nonEmptyLines
} // --AMFreportSLOCstats01b -- AMFmainB
