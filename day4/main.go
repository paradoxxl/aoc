package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"time"
)

var sampleInput = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:45] falls asleep
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:55] wakes up`

var re = regexp.MustCompile(`(?m)\[(.*)\] (Guard #(\d*) begins shift)?(falls asleep)?(wakes up)?`)

func main() {
	challenge2()
}

func challenge1Sample() {
	list := re.FindAllStringSubmatch(sampleInput, -1)

	sort.Slice(list, func(i, j int) bool {
		//Mon Jan 2 15:04:05 -0700 MST 2006
		iDate, err := time.Parse("2006-01-02 15:04", list[i][1])
		if err != nil {
			panic(err)
		}
		jDate, err := time.Parse("2006-01-02 15:04", list[j][1])
		if err != nil {
			panic(err)
		}
		return iDate.Before(jDate)
	})

	guards := make(map[int][]int)
	guardsTotalAsleep := make(map[int]int)

	var guard int
	var err error

	for i, v := range list {
		if v[3] != "" && v[3] != " " {
			guard, err = strconv.Atoi(v[3])
			if err != nil {
				panic(err)
			}
			continue
		}

		if v[4] != "" && v[4] != " " {
			gotosleep, err := time.Parse("2006-01-02 15:04", list[i][1])
			if err != nil {
				panic(err)
			}
			wakeup, err := time.Parse("2006-01-02 15:04", list[i+1][1])
			if err != nil {
				panic(err)
			}

			guardsTotalAsleep[guard] += int(wakeup.Sub(gotosleep))

			if _, found := guards[guard]; !found {
				guards[guard] = make([]int, 60)

			}

			for x := gotosleep.Minute(); x < wakeup.Minute(); x++ {

				guards[guard][x]++
			}
		}
	}

	var largestMin = 0
	var largestGuard = guard

	for k, v := range guardsTotalAsleep {
		if v > guardsTotalAsleep[largestGuard] {
			largestGuard = k
		}
	}

	for t, num := range guards[largestGuard] {
		if num > guards[largestGuard][largestMin] {
			largestMin = t
		}
	}

	fmt.Println("Guard:", largestGuard, "nummin", largestMin, "val", largestGuard*largestMin)

	for gid, g := range guards {
		for min, num := range g {
			if num > guards[largestGuard][largestMin] {
				largestGuard = gid
				largestMin = min
			}
		}
	}

	fmt.Println("Guard:", largestGuard, "nummin", largestMin, "val", largestGuard*largestMin)

}

type Person struct {
	Name string
	Age  int
}

func challenge1() {

	input, err := ioutil.ReadFile("C:\\Users\\dominik\\Documents\\go\\src\\github.com\\paradoxxl\\aoc\\day4\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	list := re.FindAllStringSubmatch(string(input), -1)

	sort.Slice(list, func(i, j int) bool {
		//Mon Jan 2 15:04:05 -0700 MST 2006
		iDate, err := time.Parse("2006-01-02 15:04", list[i][1])
		if err != nil {
			panic(err)
		}
		jDate, err := time.Parse("2006-01-02 15:04", list[j][1])
		if err != nil {
			panic(err)
		}
		return iDate.Before(jDate)
	})

	guards := make(map[int][]int)
	guardsTotalAsleep := make(map[int]int)

	var guard int

	for i, v := range list {
		if v[3] != "" && v[3] != " " {
			guard, err = strconv.Atoi(v[3])
			if err != nil {
				panic(err)
			}
			continue
		}

		if v[4] != "" && v[4] != " " {
			gotosleep, err := time.Parse("2006-01-02 15:04", list[i][1])
			if err != nil {
				panic(err)
			}
			wakeup, err := time.Parse("2006-01-02 15:04", list[i+1][1])
			if err != nil {
				panic(err)
			}

			guardsTotalAsleep[guard] += int(wakeup.Sub(gotosleep))

			if _, found := guards[guard]; !found {
				guards[guard] = make([]int, 60)

			}

			for x := gotosleep.Minute(); x < wakeup.Minute(); x++ {

				guards[guard][x]++
			}
		}
	}

	var largestMin = 0
	var largestGuard = guard

	for k, v := range guardsTotalAsleep {
		if v > guardsTotalAsleep[largestGuard] {
			largestGuard = k
		}
	}

	for t, num := range guards[largestGuard] {
		if num > guards[largestGuard][largestMin] {
			largestMin = t
		}
	}

	fmt.Println("Guard:", largestGuard, "nummin", largestMin, "val", largestGuard*largestMin)

}

func challenge2() {

	input, err := ioutil.ReadFile("C:\\Users\\dominik\\Documents\\go\\src\\github.com\\paradoxxl\\aoc\\day4\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	list := re.FindAllStringSubmatch(string(input), -1)

	sort.Slice(list, func(i, j int) bool {
		//Mon Jan 2 15:04:05 -0700 MST 2006
		iDate, err := time.Parse("2006-01-02 15:04", list[i][1])
		if err != nil {
			panic(err)
		}
		jDate, err := time.Parse("2006-01-02 15:04", list[j][1])
		if err != nil {
			panic(err)
		}
		return iDate.Before(jDate)
	})

	guards := make(map[int][]int)

	var guard int

	for i, v := range list {
		if v[3] != "" && v[3] != " " {
			guard, err = strconv.Atoi(v[3])
			if err != nil {
				panic(err)
			}
			continue
		}

		if v[4] != "" && v[4] != " " {
			gotosleep, err := time.Parse("2006-01-02 15:04", list[i][1])
			if err != nil {
				panic(err)
			}
			wakeup, err := time.Parse("2006-01-02 15:04", list[i+1][1])
			if err != nil {
				panic(err)
			}

			if _, found := guards[guard]; !found {
				guards[guard] = make([]int, 60)

			}

			for x := gotosleep.Minute(); x < wakeup.Minute(); x++ {

				guards[guard][x]++
			}
		}
	}

	var largestMin = 0
	var largestGuard = guard

	for gid, g := range guards {
		for min, num := range g {
			if num > guards[largestGuard][largestMin] {
				largestGuard = gid
				largestMin = min
			}
		}
	}

	fmt.Println("Guard:", largestGuard, "nummin", largestMin, "val", largestGuard*largestMin)

}
