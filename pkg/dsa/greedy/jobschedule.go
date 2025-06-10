package greedy

import (
	"fmt"
	"sort"
)

type Job struct {
	Name     string
	Deatline int
	Profit   int
}

func ScheduleJobs(jobs []Job) []string {
	slots := 3
	var scheduledJobs []string = make([]string, slots)
	occupied := make([]bool, 3)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Profit > jobs[j].Profit
	})
	for _, job := range jobs {
		for j := min(slots-1, job.Deatline-1); j >= 0; j-- {
			if occupied[j] == false {
				occupied[j] = true
				scheduledJobs[j] = job.Name
				break
			}
		}
	}
	for _, sjob := range scheduledJobs {
		fmt.Println(sjob)
	}

	return scheduledJobs
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
