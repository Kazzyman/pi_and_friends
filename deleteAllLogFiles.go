package main

import (
	"fmt"
	"os"
)

func deleteAllLogFiles() { // case 999: // -- AMFdeleteAllLogFilesA
	// list the files to be deleted:
	filename111 := "dataLog-From_calculate-pi-and-friends.txt"
	filename112 := "dataLog-From_BBPfConcurent.txt"
	filename113 := "dataLog-From_Nilakantha_Method_lengthy_prints.txt"
	filename114 := "dataLog-From_AM_Method_lengthy_prints.txt"
	filename115 := "dataLog-From_BBPF_Method_lengthy_prints.txt"
	filename116 := "dataLog-From_Chudnovsky_Method_lengthy_prints.txt"
	filename117 := "big_pie_is_in_here.txt"
	filename118 := "logfile_from_selection_180.txt"
	// Check if those files exist:
	if _, err := os.Stat(filename111); err != nil {
		fmt.Println("File dataLog-From_calculate-pi-and-friends.txt does not exist")
	}
	if _, err := os.Stat(filename112); err != nil {
		fmt.Println("File dataLog-From_BBPfConcurent.txt does not exist")
	}
	if _, err := os.Stat(filename113); err != nil {
		fmt.Println("File dataLog-From_Nilakantha_Method_lengthy_prints.txt does not exist")
	}
	if _, err := os.Stat(filename114); err != nil {
		fmt.Println("File dataLog-From_AM_Method_lengthy_prints.txt does not exist")
	}
	if _, err := os.Stat(filename115); err != nil {
		fmt.Println("File From_BBPF_Method_lengthy_prints.txt does not exist")
	}
	if _, err := os.Stat(filename116); err != nil {
		fmt.Println("File dataLog-From_Chudnovsky_Method_lengthy_prints.txt does not exist")
	}
	if _, err := os.Stat(filename117); err != nil {
		fmt.Println("File big_pie_is_in_here.txt does not exist")
	}
	if _, err := os.Stat(filename118); err != nil {
		fmt.Println("File logfile_from_selection_180.txt does not exist")
	}
	// Delete those files:
	err := os.Remove(filename118)
	if err != nil {
		fmt.Println(err)
	}
	err1 := os.Remove(filename111)
	if err1 != nil {
		fmt.Println(err)
	}
	err2 := os.Remove(filename112)
	if err2 != nil {
		fmt.Println(err)
	}
	err3 := os.Remove(filename113)
	if err3 != nil {
		fmt.Println(err)
	}
	err4 := os.Remove(filename114)
	if err4 != nil {
		fmt.Println(err)
	}
	err5 := os.Remove(filename115)
	if err5 != nil {
		fmt.Println(err)
	}
	err6 := os.Remove(filename116)
	if err6 != nil {
		fmt.Println(err)
	}
	err7 := os.Remove(filename117)
	if err7 != nil {
		fmt.Println(err)
	}
} // end of deleteAllLogFiles() // -- AMFdeleteAllLogFilesB
