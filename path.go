package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	ID             string
	Predecessors   []string
	Duration       int
	EarlyStart     int
	LateStart     int
	EarlyFinish    int
	LateFinish    int
	Slack          int
	Critical       bool
}

func main() {
	// Example activities and their details
	activities := map[string]*Activity{
		"A": {ID: "A", Predecessors: []string{}, Duration: 3},
		"B": {ID: "B", Predecessors: []string{"A"}, Duration: 2},
		"C": {ID: "C", Predecessors: []string{"A"}, Duration: 6},
		"D": {ID: "D", Predecessors: []string{"B", "C"}, Duration: 4},
	}

	// Sort activities by ID for easier processing
	sortedActivities := make([]*Activity, 0, len(activities))
	for _, activity := range activities {
		sortedActivities = append(sortedActivities, activity)
	}
	sort.Slice(sortedActivities, func(i, j int) bool {
		return sortedActivities[i].ID < sortedActivities[j].ID
	})

	// Calculate forward pass
	for _, activity := range sortedActivities {
		for _, predecessorID := range activity.Predecessors {
			predecessor := activities[predecessorID]
			if predecessor.EarlyFinish > activity.EarlyStart {
				activity.EarlyStart = predecessor.EarlyFinish
			}
		}
		activity.EarlyFinish = activity.EarlyStart + activity.Duration
	}

	// Calculate backward pass
	for i := len(sortedActivities) - 1; i >= 0; i-- {
		activity := sortedActivities[i]
		for _, successorID := range findSuccessors(activities, activity.ID) {
			successor := activities[successorID]
			if successor.LateStart < activity.LateFinish {
				activity.LateFinish = successor.LateStart
			}
		}
		activity.LateStart = activity.LateFinish - activity.Duration
		activity.Slack = activity.LateFinish - activity.EarlyFinish
	}

	// Determine critical activities
	for _, activity := range sortedActivities {
		activity.Critical = activity.Slack == 0
	}

	// Print results
	fmt.Println("Activity | Early Start | Early Finish | Late Start | Late Finish | Slack | Critical")
	for _, activity := range sortedActivities {
		fmt.Printf("%s\t | %d\t\t\t | %d\t\t\t | %d\t\t\t | %d\t\t\t | %d\t | %v\n",
			activity.ID, activity.EarlyStart, activity.EarlyFinish,
			activity.LateStart, activity.LateFinish, activity.Slack, activity.Critical)
	}

	// Print critical path
	fmt.Println("Critical Path:")
	printCriticalPath(activities)
}

func findSuccessors(activities map[string]*Activity, activityID string) []string {
	successors := []string{}
	for _, activity := range activities {
		for _, predecessorID := range activity.Predecessors {
			if predecessorID == activityID {
				successors = append(successors, activity.ID)
			}
		}
	}
	return successors
}

func printCriticalPath(activities map[string]*Activity) {
	criticalActivity := activities["D"]
	path := []string{criticalActivity.ID}
	for {
		predecessors := findPredecessors(activities, criticalActivity.ID)
		if len(predecessors) == 0 {
			break
		}
		criticalActivity = activities[predecessors[0]]
		path = append(path, criticalActivity.ID)
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i] + " -> ")
	}
	fmt.Println()
}
