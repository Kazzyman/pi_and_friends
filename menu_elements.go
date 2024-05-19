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
	// re := regexp.MustCompile(`Unix(.*?)2023`)
	re := regexp.MustCompile(`/Users(.*?)pi`)
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
	/*
					cwd, err := os.Getwd()
					if err != nil {
						fmt.Printf("Error getting current working directory: %v\n", err)
					} else {
						fmt.Printf("Current working directory: %s\n", cwd)
					}
			// this was not an issue

		data, err := ioutil.ReadFile(filenameOfThisFile)
		if err != nil {
			fmt.Printf("File content: %s\n", string(data))
			log.Fatalf("Error reading file via ioutil.ReadFile: %v\n", err)
			// this too produced an error (prior to my locating the actual mistake I had made)
		}

	*/

	if _, err := os.Stat(filenameOfThisFile); os.IsNotExist(err) {
		fmt.Printf("File does not exist per os.Stat: %s\n", filenameOfThisFile)
	} else if os.IsPermission(err) {
		fmt.Printf("Permission denied per os.Stat: %s\n", filenameOfThisFile)
	} else if err != nil {
		fmt.Printf("Error stating file per os.Stat: %v\n", err)
	} else {
		fmt.Printf("%s File exists and is accessible per os.Stat\n", filenameOfThisFile)
	}

	file, err := os.Open(filenameOfThisFile) // open this source code file and get a handle to it

	if err != nil {
		fmt.Println(err)
		fmt.Println("rick fucked up!")
		log.Fatal(err)
	}

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File does not exist: %s\n", filenameOfThisFile)
			fmt.Println(err)
			fmt.Println("rick fucked something up! Because File is reported to not exist")
			log.Fatal(err)
		} else if os.IsPermission(err) {
			fmt.Printf("Permission denied: %s\n", filenameOfThisFile)
			fmt.Println(err)
			fmt.Println("rick fucked up! Because Permission denied")
			log.Fatal(err)
		} else {
			fmt.Printf("Some other Error opening file: %v\n", err)
			fmt.Println(err)
			fmt.Println("rick just fucked something up!")
			log.Fatal(err)
		}
		// return
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

func reportSLOCstatsOld(filenameOfThisFile string) (totalLines, nonEmptyLines int) {
	// Check file existence and permissions
	fileInfo, err := os.Stat(filenameOfThisFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("File does not exist: %s\n", filenameOfThisFile)
		} else if os.IsPermission(err) {
			log.Fatalf("Permission denied: %s\n", filenameOfThisFile)
		} else {
			log.Fatalf("Error stating file: %v\n", err)
		}
	} else {
		fmt.Printf("File exists and is accessible: %s\n", filenameOfThisFile)
		fmt.Printf("File size: %d bytes\n", fileInfo.Size())
		fmt.Printf("File mode: %s\n", fileInfo.Mode())
		fmt.Printf("File modified: %s\n", fileInfo.ModTime())
	}

	// Open the file
	file, err := os.Open(filenameOfThisFile)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close() // Ensure the file is closed

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalLines++
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			nonEmptyLines++
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	// Return the results
	return totalLines, nonEmptyLines
}
