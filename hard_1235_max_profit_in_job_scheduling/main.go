package main

import (
	"fmt"
	"sort"
)

func main() {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}

	out := jobScheduling(startTime, endTime, profit)

	fmt.Printf("%v\n", out)
}

// Dynamic Programming
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]job, 0, len(startTime))

	for i := range startTime {
		jobs = append(jobs, job{start: startTime[i], end: endTime[i], profit: profit[i]})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	profits := make(map[int]int, len(jobs))
	profits[0] = jobs[0].profit

	for i := 1; i < len(jobs); i++ {
		profit := jobs[i].profit
		for j := i - 1; j >= 0; j-- {
			if jobs[j].end <= jobs[i].start {
				profit += profits[j]
				break
			}
		}

		if profit > profits[i-1] {
			profits[i] = profit
		} else {
			profits[i] = profits[i-1]
		}
	}

	return profits[len(jobs)-1]
}

type job struct {
	start  int
	end    int
	profit int
}

// Bruteforce
func jobSchedulingBrute(startTime []int, endTime []int, profit []int) int {
	jobs := make([]jobBrute, 0, len(startTime))

	for i := range startTime {
		jobs = append(jobs, jobBrute{start: startTime[i], end: endTime[i], profit: profit[i]})
	}

	fmt.Printf("%v\n", jobs)

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	fmt.Printf("%v\n", jobs)

	for i := 0; i < len(jobs); i++ {
		for j := i + 1; j < len(jobs); j++ {
			if jobs[i].end <= jobs[j].start {
				jobs[i].nexts = append(jobs[i].nexts, j)
			}
		}
	}

	max := 0
	for i := 0; i < len(jobs); i++ {
		m := getMax(jobs, i)
		if max < m {
			max = m
		}
	}

	fmt.Printf("%v\n", jobs)

	return max
}

func getMax(jobs []jobBrute, i int) int {
	max := 0
	for _, j := range jobs[i].nexts {
		m := getMax(jobs, j)
		if max < m {
			max = m
		}
	}

	return jobs[i].profit + max
}

type jobBrute struct {
	start  int
	end    int
	profit int
	nexts  []int
}
