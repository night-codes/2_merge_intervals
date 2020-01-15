// The input is a list of intervals (two integer points); write a function/method that merges any
// N intervals in the list that have common points (intervals [1, 3] and [3, 6] have a common point
// of 3; [4, 8] and [6, 10] have common points of 6, 7, 8) and returns the merged list (where no two
// intervals intersect). it is allowed to modify the input. Try to avoid allocating extra memory for
// the output.

package main

import (
	"fmt"
	"sort"
)

type (
	interval  [2]int
	intervals []interval
)

func (a intervals) Len() int           { return len(a) }
func (a intervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intervals) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(ins []interval) []interval {
	if len(ins) == 0 {
		return ins
	}
	sort.Sort(intervals(ins))

	j := 0
	for i := 1; i < len(ins); i++ {
		if ins[j][1] >= ins[i][0] {
			if ins[j][1] < ins[i][1] {
				ins[j][1] = ins[i][1]
			}
		} else {
			j++
			ins[j] = ins[i]
		}
	}
	return append([]interval{}, ins[:j+1]...)
}

func main() {
	input := []interval{{47, 60}, {1, 3}, {17, 18}, {20, 24}, {24, 25}, {7, 11}, {45, 50}, {10, 17}, {30, 32}, {31, 35}, {4, 15}}
	fmt.Println(input)
	fmt.Println(merge(input)) // [[1 18] [20 25] [30 35] [45 50]]
}
