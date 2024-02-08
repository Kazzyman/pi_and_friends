package main

import "fmt"

func DisplayPythagorean(selection int) { // case 11:
	fmt.Print("\n\n\n -- You entered '", selection, "' to review the derivation of the Pythagorean, which was needed in method #4. We will\n")
	fmt.Println("be geometrically deriving the Pythagorean theorem according to the 12th century Indian, Bhaskara.")
	fmt.Println("    We begin with a square of area c\u00b2. We then partially fill that square with four congruent")
	fmt.Println("right triangles each with its right angle opposite to one of the sides 'c'. Each of the four congruent")
	fmt.Println("triangles now have sides c, a, and b. 'c' being the hypotenuse of each; 'a' being the shortest side and 'b'")
	fmt.Println("being the longer of sides 'a' and 'b'. Thus leaving a small square in the center, which we can ignore.")

	fmt.Println("     Next, we detach and slide two adjacent right triangles, each to its opposite side of the parent square.")
	fmt.Println("We now have two attached squares a x a and b x b. We can then clearly see that", colorCyan, "c\u00b2 = a\u00b2 + b\u00b2", colorReset)
	fmt.Println(" --- The Pythagorean theorem is thereby derived, per Bhaskara --- ")

	fmt.Println("\n     Apparently it was not until 1876 that James Garfield (yea, the former U.S. president) discovered")
	fmt.Println("the following geometric proof of the Pythagorean theorem. He began with a right triangle similar ")
	fmt.Println("to those we encountered from Bhaskara. But any shape of right triangle will suffice for his proof. ")
	fmt.Println("Next, he cloned a second triangle placing it such that its most-acutely-angled point touches the other")
	fmt.Println("triangle's least-acutely-angled point such that side 'a' of the first is colinear with side 'c' of the")
	fmt.Println("second triangle. Thus creating the potential of a new triangle: an isosceles right triangle of dimensions:")
	fmt.Println("c, c, x. ")
	fmt.Println("     After having converted that potential to an actual we are left with a trapezoid. That is, a rectangle")
	fmt.Println("mated to a right triangle; in this case a trapezoid of dimensions a, a+b, b, and x. We know that for any")
	fmt.Println("trapezoid the area is equal to the height [in this case a+b] times the average of the lengths of the ")
	fmt.Println("two adjacent sides [in this case, half of the sum of a+b]. Which is to say that, in the case of OUR trapezoid, ")
	fmt.Println("the Area is  : ", colorCyan, "A = (a+b) * 1/2(a+b)  -- Or: --  A = 1/2(a+b)\u00b2 ")
	fmt.Println(colorReset, "    But the area of this particular trapezoid that we have constructed is, obviously, also a*b")
	fmt.Println("(the combined area of the initial two right triangles), plus 1/2*c\u00b2.   Or:  -- ", colorCyan)
	fmt.Println("A = ab + 1/2*c\u00b2 \n", colorReset)
	fmt.Println("           a new equality can then be constructed from the two above equalities: \n", colorCyan)
	fmt.Println("    1/2(a+b)\u00b2  =  ab + 1/2(c\u00b2) \n", colorReset)
	fmt.Println("                          and reducing it thusly : ", colorCyan)
	fmt.Println("        (a+b)\u00b2 = 2ab + c\u00b2 ")
	fmt.Println(" a\u00b2 + 2ab + b\u00b2 = 2ab + c\u00b2 ")
	fmt.Println("       a\u00b2 + b\u00b2 = c\u00b2 \n", colorReset)
	fmt.Println("... proves the Pythagorean per Garfield; though, obviously, many other proofs do exist.\n\n")
	fmt.Println("Select 12 at menu to display prior results")
}
func DisplayPythagoreanCode() { // case 31:
	var DisplayPythagoreanRune = `

This section was conceived of and written entirely by yours-truly. 

func DisplayPythagorean(selection int){  // case 11: 
    fmt.Print("\n\n\n -- You entered '", selection, "' to review the derivation of the Pythagorean, which was needed in method #4. We will\n")
fmt.Println("be geometrically deriving the Pythagorean theorem according to the 12th century Indian, Bhaskara.")
fmt.Println("    We begin with a square of area c\u00b2. We then partially fill that square with four congruent")
fmt.Println("right triangles each with its right angle opposite to one of the sides 'c'. Each of the four congruent")
fmt.Println("triangles now have sides c, a, and b. 'c' being the hypotenuse of each; 'a' being the shortest side and 'b'")
fmt.Println("being the longer of sides 'a' and 'b'. Thus leaving a small square in the center, which we can ignore.")

fmt.Println("     Next, we detach and slide two adjacent right triangles, each to its opposite side of the parent square.")
fmt.Println("We now have two attached squares a x a and b x b. We can then clearly see that", string(colorCyan), "c\u00b2 = a\u00b2 + b\u00b2", string(colorReset))
fmt.Println(" --- The Pythagorean theorem is thereby derived, per Bhaskara --- ")

fmt.Println("\n     Apparently it was not until 1876 that James Garfield (yea, the former U.S. president) discovered")
fmt.Println("the following geometric proof of the Pythagorean theorem. He began with a right triangle similar ")
fmt.Println("to those we encountered from Bhaskara. But any shape of right triangle will suffice for his proof. ")
fmt.Println("Next, he cloned a second triangle placing it such that its most-acutely-angled point touches the other")
fmt.Println("triangle's least-acutely-angled point such that side 'a' of the first is colinear with side 'c' of the")
fmt.Println("second triangle. Thus creating the potential of a new triangle: an isosceles right triangle of dimensions:")
fmt.Println("c, c, x. ")
fmt.Println("     After having converted that potential to an actual we are left with a trapezoid. That is, a rectangle")
fmt.Println("mated to a right triangle; in this case a trapezoid of dimensions a, a+b, b, and x. We know that for any")
fmt.Println("trapezoid the area is equal to the height [in this case a+b] times the average of the lengths of the ")
fmt.Println("two adjacent sides [in this case, half of the sum of a+b]. Which is to say that, in the case of OUR trapezoid, ")
fmt.Println("the Area is  : ", string(colorCyan), "A = (a+b) * 1/2(a+b)  -- Or: --  A = 1/2(a+b)\u00b2 ")
fmt.Println(string(colorReset), "    But the area of this particular trapezoid that we have constructed is, obviously, also a*b")
fmt.Println("(the combined area of the initial two right triangles), plus 1/2*c\u00b2.   Or:  -- ", string(colorCyan))
fmt.Println("A = ab + 1/2*c\u00b2 \n", string(colorReset))
fmt.Println("           a new equality can then be constructed from the two above equalities: \n", string(colorCyan))
fmt.Println("    1/2(a+b)\u00b2  =  ab + 1/2(c\u00b2) \n", string(colorReset))
fmt.Println("                          and reducing it thusly : ", string(colorCyan))
fmt.Println("        (a+b)\u00b2 = 2ab + c\u00b2 ")
fmt.Println(" a\u00b2 + 2ab + b\u00b2 = 2ab + c\u00b2 ")
fmt.Println("       a\u00b2 + b\u00b2 = c\u00b2 \n", string(colorReset))
fmt.Println("... proves the Pythagorean per Garfield; though, obviously, many other proofs do exist.\n\n")
    fmt.Println("Select 12 at menu to display prior results")
}
// By Richard Woolley
 `
	fmt.Println(DisplayPythagoreanRune)
}
