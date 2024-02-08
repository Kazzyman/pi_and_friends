package main

import (
	"fmt"
	"math"
)

func EulersNumber(selection int) { // case 9: // -- AMFEulersNumberA
	var n float64
	var sum float64
	var Eulers float64
	fmt.Println("\n\n\nEuler's Number \u2107 or \u2147 the natural logarithmic base")
	fmt.Println("\u2147 = (1+1/n)^n")
	fmt.Println("... the limit of an increasing value for n\n\n")
	n = 9
	sum = 1 + 1/n
	Eulers = math.Pow(sum, n)
	fmt.Print(Eulers)
	fmt.Printf(" was calculated with an exponent of %0.f \n", n)
	n = 99
	sum = 1 + 1/n
	Eulers = math.Pow(sum, n)
	fmt.Print(Eulers)
	fmt.Printf("  was calculated with an exponent of %0.f \n", n)
	n = 999
	sum = 1 + 1/n
	Eulers = math.Pow(sum, n)
	fmt.Print(Eulers)
	fmt.Printf(" was calculated with an exponent of %0.f \n", n)
	n = 9999
	sum = 1 + 1/n
	Eulers = math.Pow(sum, n)
	fmt.Print(Eulers)
	fmt.Printf(" was calculated with an exponent of %0.f \n", n)
	n = 99999999999
	sum = 1 + 1/n
	Eulers = math.Pow(sum, n)
	fmt.Print(Eulers)
	fmt.Printf(" was calculated with an exponent of %0.f \n", n)
	fmt.Println("2.71828182845904523536028747135266249775724 is Euler's Number from the web")
	fmt.Println("2.718281828 is the dollar value of $1 compounded continuously for one year.")
	fmt.Println("2.714567 is from daily compound interest which is near-enough to continuous interest.\n")
	// change to cyan
	fmt.Println(colorCyan, "An account starts with $1.00 and pays 100 percent interest per year. If the interest is credited once, ")
	fmt.Println("at the end of the year, the value of the account at year-end will be $2.00. What happens if the interest")
	fmt.Println("is computed and credited more frequently during the year?\n")
	fmt.Println("If the interest is credited twice in the year, the interest rate for each 6 months will be 50%, so the ")
	fmt.Println("initial $1 is multiplied by 1.5 twice, yielding $2.25 at the end of the year. Compounding quarterly")
	fmt.Println("yields $2.44140625, and compounding monthly yields $2.613035 = $1.00 × (1 + 1/12)^12 Generally, if there")
	fmt.Println("are n compounding intervals, the interest for each interval will be 100%/n and the value at the end of")
	fmt.Println("the year will be $1.00 × (1 + 1/n)^n.")
	// And now, here comes a fun rune to print a multi-line "string"
	// ... define a rune with the ` :: back-quote character located on the ~ tilda key
	Ricks_rune_Paragraph := `  
Bernoulli noticed that this sequence approaches a limit (the force of interest) with larger n and, thus, smaller 
compounding intervals. Compounding weekly (n = 52) yields $2.692596..., while compounding daily (n = 365) yields
$2.714567... (approximately two cents more). The limit as n grows large is the number that came to be known as e.
That is, with continuous compounding, the account value will reach $2.718281828 
`
	fmt.Println(Ricks_rune_Paragraph, colorReset)
	// written entirely by Richard Woolley
} // end of EulersNumber() // -- AMFEulersNumberB
