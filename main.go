package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/gorhill/cronexpr"
)

// todo: point this at an actual cron file

var startDate time.Time = time.Date(
	2021, 6, 4, 00, 00, 00, 0, time.UTC)
var endDate time.Time = time.Date(
	2021, 6, 5, 00, 00, 00, 0, time.UTC)

func main() {

	data := []string{"5 2 1,16 * *"}
	startDate = startDate.Add(time.Second * -1)
	times := map[int64]int{}
	timesString := map[string]int{}

	for _, line := range data {
		expr := cronexpr.MustParse(line)
		runningTime := startDate

		for i := 1; i < 2000; i++ {
			nextTime := expr.Next(runningTime)
			if nextTime.After(endDate) {
				break
			}
			timeUNIX := nextTime.Unix()
			// fmt.Println(timeUNIX)
			timeString := nextTime.Truncate(time.Minute * 5).Format(time.RFC822Z)
			timesString[timeString] = timesString[timeString] + 1
			times[timeUNIX] = times[timeUNIX] + 1
			runningTime = nextTime.Add(time.Second * 4)
		}

	}

	// for line, count := range times {
	// 	fmt.Printf("%d\t%d\n", line, count)
	// }

	keys := make([]string, 0, len(timesString))
	for k := range timesString {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("DATE,COUNT")
	for _, count := range keys {
		fmt.Printf("\"%s\",%d\n", count, timesString[count])
	}

}
