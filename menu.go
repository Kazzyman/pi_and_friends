package main

import (
	"fmt"
	"time"
)

// This next func WAS extremely long, CONTAINED all menus, and also a very long switch statement -- AMFmenusA
func DisplayMenus(totalLines, nonEmptyLines int, filenameOfThisFile, SansVerOfNameOfThisFile string) {

	pages_of_code := totalLines / 49

	date := time.Now() // only used here once as date.Format("(01-02-2006)") at the top of the main menu **** NOT FOR RELEASE **********
	// *********** the above line will need to be commented out for release, and the date below entered manually ************************

	fmt.Print(colorYellow,
		"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n Menu, 1 of 2 ",

		colorGreen, " inception: ", colorReset, SansVerOfNameOfThisFile, colorGreen,
		"    build date: ", colorCyan, date.Format("(01-02-2006)"), colorReset, "\n\n")

	fmt.Println(colorRed, "Enter your selection from below, or 96 for notes on using this app\n", colorReset)
	fmt.Print("2: ", colorCyan, " Pi: ", colorReset, "Bailey–Borwein–Plouffe formula for π, discovered in 1995", colorYellow)
	fmt.Print("\n", colorReset)
	fmt.Println("        ---> uses big.Float types, and also uses channels: GOMAXPROCS(numCPU)\n")

	fmt.Print("14:", colorCyan, " Pi: ", colorReset, "modified Archimedes' method using Go's math/big objects: 3,012 digits\n")
	fmt.Println("        ---> uses big.Float (Richard's favorite: 'cause it's fully understood)")
	fmt.Print("         (employs the Pythagorean, ", colorGreen, "ENTER '11'", colorReset, " for a review of its derivation)\n\n")

	fmt.Print("15:", colorCyan, " Pi: ", colorReset, "Chudnovsky method,", colorYellow, " state of the art! ", colorReset, "But God only knows how it works\n")
	fmt.Println("        ---> uses big.Float types", colorYellow, "Gold Standard! \n", colorReset)

	fmt.Print("5: ", colorCyan, " Pi: ", colorReset, "An infinite series by Indian astronomer Nilakantha Somayaji circa 1500\n")
	fmt.Println("        ---> uses big.Floats types         AKA: Keļallur Comatiri: 1444–1545")
	fmt.Println("     π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) ...")
	fmt.Println("         ... the equation can be found in an ancient Sanskrit verse\n")

	fmt.Println("6:", colorCyan, "Pi:", colorReset, "Gottfried Wilhelm Leibniz, late 17th century ( and by Isaac Newton )")
	fmt.Println("     π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 - 1/11 + 1/13 ...")
	fmt.Println("       Four-Billion iterations to be executed ------ using float64 types <---")
	fmt.Println("       10 digits of π : in about ten seconds \n")

	fmt.Println("7:", colorCyan, "Pi:", colorReset, "The Gregory-Leibniz series circa 1676")
	fmt.Println("     π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) ...")
	fmt.Println("       Three-Hundred-Million iterations are executed ----- float64 types <---")
	fmt.Println("       9 digits of π : under 4 seconds\n")

	fmt.Println(colorYellow, "12: Display the Stats and Utilities Menu \n")

	fmt.Println(" 13: Menu, 2 of 2 \n", colorReset)

	fmt.Println("47: to End/Exit", colorCyan, "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, colorPurple,
		"\u00a9 2023, by Richard Hart Woolley", colorReset)
	fmt.Println(colorCyan, "                That is", pages_of_code, "pages of code at 49 lines per page\n", colorReset)
	fmt.Print("Enter your selection, 1 -> x", colorRed, " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", colorReset)

	fmt.Scan(&selection) // Pause and obtain Selection

	if selection == 13 { // ==========================================================================================================================

		fmt.Print(colorYellow)
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

		// --AMFmenu01a
		fmt.Println(colorGreen, "    build date :", colorCyan, date.Format("(01-02-2006)"), colorReset, "  inception :", SansVerOfNameOfThisFile)
		// --AMFmenu01b

		fmt.Print(colorYellow)
		fmt.Print("\nSECOND MENU:")
		fmt.Print(colorRed,
			" Enter a selection from below, or 96 for notes on using the app", colorCyan, "\n")

		fmt.Println("\n18:  Calculate", colorGreen, "the second or third Root of y (x\u221Ay)", colorReset, "from first-principles")
		fmt.Println("      ... via the ratio of y:1 of perfect products \n", colorCyan)
		fmt.Println("9:   Calculate", colorGreen, "Euler's Number: \u2107 or \u2147", colorReset, "the natural logarithmic base")
		fmt.Println("        Explore the limit of interest\n", colorCyan)
		fmt.Println("10:  Calculate the", colorGreen, "Erdős-Borwein constant", colorReset, "from a breif infinite series\n")
		fmt.Println("8:", colorGreen, " Pi:", colorReset, "An infinite series by John Wallis circa 1655")
		fmt.Println("      π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ... ")
		fmt.Println("        One-Billion iterations will be executed; option for 40 billion")
		fmt.Println("        9 digits of π -- a billion loops, in seconds -- option for 10 digits\n")
		fmt.Println("19:", colorGreen, "Pi:", colorReset, "Open the 'Spigot' algorithm, instantly bakes way too much pie :)")
		fmt.Println("      ... Can easily spit out 10,000 digits of π !!!!!\n", colorCyan)

		fmt.Println("37:", colorGreen, "Pi:", colorReset, "Gauss–Legendre algorithm ")

		fmt.Println("40:", colorGreen, "Pi:", colorReset, "Nifty 'ScoreBoard' using Nilakantha's formula", colorYellow,
			"(Ctrl-C to exit it)", colorCyan)
		fmt.Println("41:", colorGreen, "Pi:", colorReset, "Bailey–Borwein–Plouffe formula", colorCyan, "[concurent]",
			colorReset)
		fmt.Println("42:", colorGreen, "Pi:", colorReset, "BBP formula to 46 digits")
		fmt.Println("44:", colorGreen, "Pi:", colorReset, "via Leibniz method in one billion iterations [runs a while]")
		fmt.Println("50:", colorGreen, "Pi:", colorReset, "comparison of float64 and big.Float types \n\n")

		fmt.Println(colorYellow, "12: Display Stats and Utilities Menu  ")

		fmt.Println(colorYellow, "\n  0:  To return to main menu", colorRed, "\n")

		// ("47:  to End/Exit vvvvvvvvvvvvvvvv
		fmt.Println("47: to End/Exit", colorCyan, "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, colorPurple,
			"\u00a9 2023, by Richard Hart Woolley\n", colorReset)
		fmt.Print("Enter your selection, 1 -> x", colorRed, " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", colorReset)

		fmt.Scan(&selection) // pause and request input from the user

	} // ===============================================================================================================================================

	if selection == 12 { // ==============================================================================================================================

		// case 12: // display contents of prior results file
		fmt.Print(colorYellow)
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nStats and Utilities Menu:",

			colorRed, "Enter your selection from below", colorReset, "\n")

		fmt.Println("212:   Read log file containing stats from prior runs of ALL long-running cases \n")

		fmt.Println("202:   Review the output from prior runs of BBPF (case 2:) \n")

		fmt.Println("214:   Review the output from prior runs of modified Archimedes' method (case 14:) \n")

		fmt.Println("215:   Read stats from prior runs of the Chudnovsky method (case 15:) which did > 800 loops \n")
		fmt.Println("315:   Look at the really big pie that came out of the Chudnovsky method (case 15:) \n")

		fmt.Println("205:   Review the data from prior runs of Nilakantha Somayaji (case 5:) \n")

		fmt.Println("241:   Review stats and data from prior runs of Bailey–Borwein–Plouffe [concurent] (case 41:) \n")

		fmt.Println("180:  ", colorRed, "Run", colorReset, "the endless test of case 18: (Ctrl-C will usually end it) \n")
		fmt.Println("280:   Read log file from the aformentioned testing case \n\n")

		fmt.Println("999: Delete all log files \n\n")

		fmt.Println(colorYellow, "0:  To return to main menu", colorRed, "\n")

		// ("47:  to End/Exit vvvvvvvvvvvvvvvv
		fmt.Println("47: to End/Exit", colorCyan, "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, colorPurple,
			"\u00a9 2023, by Richard Hart Woolley\n", colorReset)
		fmt.Print("Enter your selection, 1 -> x", colorRed, " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", colorReset)

		fmt.Scan(&selection) // pause and request input from the user

	} // ===============================================================================================================================================
	menu_switch(filenameOfThisFile)
} // end DisplayMenus() -- AMFmenusB
