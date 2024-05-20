package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// --AMFreportSLOCstats01a

func reportSLOCstats(filepath string) (int, int) {
	// Patterns to identify comments, blank lines, and strings
	singleLineCommentPattern := `^\s*//`
	multiLineCommentPattern := `(?s)/\*.*?\*/`
	blankLinePattern := `^\s*$`
	stringLiteralPattern := `(?s)"(?:\\.|[^\\"])*?"|` + "`(?:\\.|[^`])*?`"

	// Compile regular expressions
	singleLineCommentRe := regexp.MustCompile(singleLineCommentPattern)
	multiLineCommentRe := regexp.MustCompile(multiLineCommentPattern)
	blankLineRe := regexp.MustCompile(blankLinePattern)
	stringLiteralRe := regexp.MustCompile(stringLiteralPattern)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		// return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLines := 0
	sloc := 0
	inMultiLineComment := false
	inMultiLineString := false

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		// Check for blank lines
		if blankLineRe.MatchString(line) {
			continue
		}

		// Check for single-line comments
		if singleLineCommentRe.MatchString(line) {
			continue
		}

		// Check for multi-line comments
		if inMultiLineComment {
			if strings.Contains(line, "*/") {
				inMultiLineComment = false
				line = multiLineCommentRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					continue
				}
			} else {
				continue
			}
		}

		if strings.Contains(line, "/*") {
			inMultiLineComment = true
			line = multiLineCommentRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				continue
			}
		}

		// Check for multi-line strings
		if inMultiLineString {
			if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
				inMultiLineString = false
				line = stringLiteralRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					continue
				}
			} else {
				continue
			}
		}

		if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
			inMultiLineString = true
			line = stringLiteralRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				continue
			}
		}

		sloc++
	}

	if err := scanner.Err(); err != nil {
		// return 0, 0, err
	}

	return totalLines, sloc
} // --AMFreportSLOCstats01b -- AMFmainB
