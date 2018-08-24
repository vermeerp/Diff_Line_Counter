package main

import (
	"fmt"

	dmppkg "github.com/sergi/go-diff/diffmatchpatch"
)

/*
1. one
2. two
3. four
4. five
5. six
6. seven
7. eight
8. nine

Expected output lines: 3, 5, 9 (8?)
*/

func main() {
	text1 := "one\ntwo\nthree\nfour\nfive\neight\nnine\nten"
	text2 := "one\ntwo\nfour\nfive\nsix\nseven\neight\nnine"

	dmp := dmppkg.New()

	diffs := dmp.DiffMain(text1, text2, false)

	for i, diff := range diffs {
		if diff.Type != dmppkg.DiffEqual {
			line := findDiffLine(i, diffs)
			fmt.Printf("Diff around line %v\n%v\n%v\n", line, diff.Type.String(), diff.Text)

		}

	}

	// fmt.Println()
	// fmt.Println(dmp.DiffPrettyText(diffs))

}

func findDiffLine(i int, diffs []dmppkg.Diff) int {
	if i == 0 {
		return 0
	}

	counter := 0
	numDiffs := len(diffs)
	for j := 0; j < i; j++ {
		isLastDiff := (j == numDiffs-1)
		if diffs[j].Type != dmppkg.DiffDelete {
			counter += numLinesInDiff(diffs[j], isLastDiff)
		}

	}

	return counter + 1
}

func numLinesInDiff(diff dmppkg.Diff, bool isLastDiff) int {
	counter := 0
	var c rune
	for _, c := range diff.Text {
		if c == '\n' {
			counter++
		}
	}

	// Last line of the text may not end with a
	if isLastDiff && c != '\n' {
		counter++
	}
	return counter
}
