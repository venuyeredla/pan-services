package greedy

import (
	"fmt"
	"testing"
)

func TestJobSched(t *testing.T) {
	jobs := make([]Job, 5)
	jobs = append(jobs, Job{Name: "a", Deatline: 2, Profit: 100})
	jobs = append(jobs, Job{Name: "b", Deatline: 1, Profit: 19})
	jobs = append(jobs, Job{Name: "c", Deatline: 2, Profit: 27})
	jobs = append(jobs, Job{Name: "d", Deatline: 1, Profit: 25})
	jobs = append(jobs, Job{Name: "e", Deatline: 3, Profit: 15})

	scheduled := ScheduleJobs(jobs)

	fmt.Printf("jobs =%v", scheduled)

}
