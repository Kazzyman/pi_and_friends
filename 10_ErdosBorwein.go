package main

import (
	"fmt"
	"math"
)

func ErdosBorwein(selection int) { // case 10: // -- AMFErdosBorweinA
	rune := `The Erdős–Borwein constant is the sum of the reciprocals of the Mersenne numbers. 
It is named after Paul Erdős and Peter Borwein. 

Paul Erdős was a Hungarian mathematician. He was one of the most prolific mathematicians 
and producers of mathematical conjectures of the 20th century. Erdős pursued and proposed 
problems in discrete mathematics, graph theory, number theory, mathematical analysis, 
approximation theory, set theory, and probability theory. Much of his work centered 
around discrete mathematics, cracking many previously unsolved problems in the field. 
He championed and contributed to Ramsey theory, which studies the conditions in which 
order necessarily appears. Overall, his work leaned towards solving previously open 
problems, rather than developing or exploring new areas of mathematics.

Erdős published around 1,500 mathematical papers during his lifetime, a figure that
remains unsurpassed. He firmly believed mathematics to be a social activity, 
living an itinerant lifestyle with the sole purpose of writing mathematical papers 
with other mathematicians. He was known both for his social practice of mathematics, 
working with more than 500 collaborators, and for his eccentric lifestyle; Time magazine 
called him "The Oddball's Oddball". He devoted his waking hours to mathematics, even 
into his later years—indeed, his death came only hours after he solved a geometry 
problem at a conference in Warsaw. Erdős's prolific output with co-authors prompted 
the creation of the Erdős number, the number of steps in the shortest path between a 
mathematician and Erdős in terms of co-authorships. 


Peter Benjamin Borwein (born St. Andrews, Scotland, May 10, 1953 – 23 August 2020) 
was a Canadian mathematician and a professor at Simon Fraser University. He is known 
as a co-author of the paper which presented the Bailey–Borwein–Plouffe algorithm 
(discovered by Simon Plouffe) for computing π. 

Borwein was born into a Jewish family. He became interested in number theory and classical 
analysis during his second year of university. He had not previously been interested in 
math, although his father was the head of the University of Western Ontario's mathematics 
department and his mother is associate dean of medicine there. Borwein and his two siblings 
majored in mathematics. 

After completing a Bachelor of Science in Honours Math at the University of Western 
Ontario in 1974, he went on to complete an MSc and Ph.D. at the University of British 
Columbia. He joined the Department of Mathematics at Dalhousie University. While he 
was there, he, his brother Jonathan Borwein and David H. Bailey of NASA wrote the 1989 
paper that outlined and popularized a proof for computing one billion digits of π. 
The authors won the 1993 Chauvenet Prize and Merten M. Hasse Prize for this paper.

In 1993, he moved to Simon Fraser University, joining his brother Jonathan in establishing 
the Centre for Experimental and Constructive Mathematics (CECM) where he developed the 
Inverse Symbolic Calculator. `
	fmt.Println("")

	fmt.Println(colorCyan, rune, "\n", colorReset)
	fmt.Println("We calculate E as E = the sum of 1/((2^n)-1) as n grows from 1 to 'infinity'")

	var Erdos_Borwein float64
	Erdos_Borwein = 1
	var iter float64
	iter = 1
	for iter < 100 {
		iter++ // iter will therefore begin at 2
		Erdos_Borwein = Erdos_Borwein + (1 / ((math.Pow(2, iter)) - 1))
		if iter == 10 || iter == 20 || iter == 30 || iter == 40 || iter == 50 || iter == 60 || iter == 70 || iter == 100 || iter == 101 {
			fmt.Println(Erdos_Borwein)
		}
		// 100 and 101 prove that we ended on 100 as the final exponent
		// ... we only get 8 results, not nine
	}
	fmt.Println("for 10, 20, 30, 40, 50, 60, 70, and 100 iterations respectively\n")
	fmt.Println("Our calculated Erdos-Borwein constant is ")
	fmt.Println(Erdos_Borwein, "after", iter, "iterations, i.e., with a final exponent of", iter)
	fmt.Println("1.606695152415291763 is what we get from the web\n")
	// written entirely by Richard Woolley
} // // -- AMFErdosBorweinB
