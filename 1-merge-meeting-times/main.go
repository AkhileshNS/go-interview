package main

import (
	"fmt"
	"sort"
)

type entry struct {
	startTime int
	endTime   int
}

func myFunction(entries []entry) []entry {
	// Write the body of your function here

	// Sort by startTime
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].startTime < entries[j].startTime
	})

	res := []entry{}
	holder := entry{}

	for i, item := range entries {
		if i == 0 {
			holder.startTime = item.startTime
			holder.endTime = item.endTime
		} else {
			if item.startTime <= holder.endTime && item.endTime > holder.endTime {
				holder.endTime = item.endTime
			}
			if item.startTime > holder.endTime {
				res = append(res, holder)
				holder.startTime = item.startTime
				holder.endTime = item.endTime
			}
		}
	}

	if len(res) == 0 {
		res = append(res, holder)
	}

	if res[len(res)-1].startTime != holder.startTime && res[len(res)-1].endTime != holder.endTime {
		res = append(res, holder)
	}

	return res
}

func main() {

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(myFunction([]entry{
		{startTime: 1, endTime: 2},
		{startTime: 2, endTime: 3},
	}))
}
