package main

import "fmt"

func Using_this_app() { // case 96:
	fmt.Print(colorCyan)
	var rune_Using_this_app = `
Any selection from any menu can be made at any menu. 
 
Each selection has a corresponding selector (+20) which displays the source code for that particular algorithm. For example, to view the 
code for selection 18 one simply enters 38 at any menu – one reason that you might want to do this is to discover the section’s authorship. 

Additionally, there are a few other operational selections which are not found on any of the menus. For example, can enter 3 to see the 
magic behind main, or enter 4 to see the code for the menus and switches.

===============
About this app:
===============

The majority of the source code which comprises this app was conceived of, designed, and implemented by Richard (Rick) Hart Wolley in late 
2022 and early 2023.   

Why does this app exist? Well, it was a rather rainy day sometime late in October and I had some time to kill. I had not done any 
software engineering for a few years and there were two languages that I had never really tried before, so that seemed like fun. 
Fortran, and Python were on my bucket list. I had also been hearing a lot of good things about Google’s new language go.lang (simply Go 
within the inner circles at Google). Revisiting Lisp (both the Emacs and the Common variants) had also been on my mind but I’ve yet to
get around to it. I had also always wanted to try constructing an algorithm that would calculate Pi. I was especially curious to see how 
many digits one could easily calculate from first principles using a home computer and a simple algorithm. I coded up an identical 
prototype in the three languages that were new to me and found that Go was so much better in every way that I can now not imagine messing 
around with any other language. Though, admittedly, I have found that Go is a bit “buggy” on Windows11, Go being intended mainly for use 
on Unix variants. Which is Ok because Linux and Mac are my preferred programming environments. 

I then got a bit carried away, and a few thousand lines of code later, here we are. 

    `
	fmt.Println(rune_Using_this_app, colorReset)
}
